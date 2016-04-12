package etcd

import (
	"path/filepath"
	"sync"
	"sync/atomic"
	"github.com/coreos/go-etcd/etcd"
	"github.com/TriangleGo/gylib/logger"
	"google.golang.org/grpc"
	"github.com/TriangleGo/gylib/service/proto"
	"time"
	"strings"
	"strconv"
	"github.com/TriangleGo/gylib/util"
)

type client struct {
	key  string
	conn *grpc.ClientConn
}

type service struct {
	clients []client
	idx     uint32
}

type service_range struct {
	Start int
	End   int
	Node  string
}

type service_pool struct {
	services          map[string]*service
	service_names     map[string]bool // store names.txt
	enable_name_check bool
	client_pool       sync.Pool
	sync.RWMutex
}

var (
	_default_pool service_pool
	pool map[int]*service_range
	secPool map[string]int
)

//
func (p *service_pool) initPool() {
	// etcd client
	p.client_pool.New = func() interface{} {
		return etcd.NewClient(etcdHosts)
	}

	p.services = make(map[string]*service)
	pool = make(map[int]*service_range)
	secPool = make(map[string]int)
}

// connect to all services
func (p *service_pool) connect() {
	client := p.client_pool.Get().(*etcd.Client)
	defer func() {
		p.client_pool.Put(client)
	}()

	for _, name := range clientFiles {
		resp, err := client.Get(name, true, true)
		if err != nil {
			logger.Debug(err)
			return
		}
		// validation check
		if !resp.Node.Dir {
			logger.Debug("not a directory")
			return
		}
		logger.Infof("dir %s", resp.Node.Key)
		for _, node := range resp.Node.Nodes {
			logger.Infof("sub node key %s, val %s", node.Key, node.Value)
			if node.Dir {
				var start, end int
				var nodeName, nodeAddress string
				// handlermap directory
				for _, node := range node.Nodes {
					logger.Infof("sec sub node key %s, val %s", node.Key, node.Value)
					switch {
					case strings.HasSuffix(node.Key, "service"):
						nodeName = node.Key
						nodeAddress = node.Value
					case strings.HasSuffix(node.Key, "Start"):
						start, _ = strconv.Atoi(node.Value)
					case strings.HasSuffix(node.Key, "End"):
						end, _ = strconv.Atoi(node.Value)
					}
				}
				logger.Infof("add service %s, %s, start = %d, end = %d", nodeName, nodeAddress, start, end)
				p.add_service(nodeName, nodeAddress, start, end)
			}
		}
		logger.Debug("services add complete")
	}
}

type temp_service struct {
	name    string
	address string
	start   int
	end     int
}

func (s *temp_service) valid() bool {
	if s == nil {
		return false
	} else {
		return s.address != "" && s.start <= s.end
	}
}

// watcher for data change in etcd directory
func (p *service_pool) watcher() {
	logger.Info("starting watcher.")
	client := p.client_pool.Get().(*etcd.Client)
	defer func() {
		p.client_pool.Put(client)
	}()

	for {
		logger.Info("for loop.")
		ch := make(chan *etcd.Response, 10)
		go func() {
			tempServiceMap := make(map[string]*temp_service)
			serviceDirs := make(map[string]string)
			for {
				if resp, ok := <-ch; ok {
					logger.Debugf("grpc incoming resp %s", test.ToJsonString(resp))
					switch resp.Action {
					case "create":
						dir := resp.Node.Key
						temp := &temp_service{}
						tempServiceMap[dir] = temp
						serviceDirs[dir] = dir
					case "set":
						key := resp.Node.Key
						val := resp.Node.Value
						var tempService *temp_service
						for _, dir := range serviceDirs {
							if strings.HasPrefix(key, dir) {
								tempService = tempServiceMap[dir]
								break
							}
						}
						if tempService != nil {
							if strings.HasSuffix(key, "service") {
								tempService.address = val
								tempService.name = key
							} else if strings.HasSuffix(key, "actionStart") {
								tempService.start, _ = strconv.Atoi(val)
							} else if strings.HasSuffix(key, "actionEnd") {
								tempService.end, _ = strconv.Atoi(val)
							}
							if tempService.valid() {
								logger.Debugf("valid service.")
								p.add_service(tempService.name, tempService.address, tempService.start, tempService.end)
								delete(tempServiceMap, tempService.name)
								delete(serviceDirs, tempService.name)
							}
						}
					case "delete":
						p.remove_service(resp.Node.Key)
					}
					//resp.Action
					//if resp.Node.Dir || (!strings.HasSuffix(resp.Node.Key, "service")) {
					//	continue
					//}
					//key, value := resp.Node.Key, resp.Node.Value
					//logger.Debugf("node status changed: key = %s, val = %s", key, value)
					//if value == "" {
					//	logger.Debugf("node delete: %v.", key)
					//	p.remove_service(key)
					//} else {
					//	logger.Debugf("node add: %v %v.", key, value)
					//	p.add_service(key, value)
					//}
				} else {
					return
				}
			}
		}()

		for _, clientFile := range clientFiles {
			logger.Debug("Watching:", clientFile)
			_, err := client.Watch(clientFile, 0, true, ch, nil)
			if err != nil {
				logger.Debug(err)
			}
		}
		<-time.After(retryDelay)
	}
}

// add a handlermap
func (p *service_pool) add_service(nodeName, nodeAddress string, start, end int) {
	p.Lock()
	defer p.Unlock()
	// name check
	if p.enable_name_check && !p.service_names[nodeName] {
		logger.Debugf("handlermap not in names: %v, ignored.", nodeName)
		return
	}

	if p.services[nodeName] == nil {
		p.services[nodeName] = &service{}
		logger.Debugf("new handlermap type: %v.", nodeName)
	}
	service := p.services[nodeName]
	logger.Info(test.ToJsonString(service))
	if conn, err := grpc.Dial(nodeAddress, grpc.WithTimeout(dialTimeout), grpc.WithInsecure()); err == nil {
		service.clients = append(service.clients, client{nodeName, conn})
		AddNode(start, end, nodeName)
		logger.Debugf("handlermap added: %s : %s.", nodeName, nodeAddress)
	} else {
		logger.Debugf("handlermap not added: %s:%s. Error: %v.", nodeName, nodeAddress, err)
	}
}

// remove a handlermap
func (p *service_pool) remove_service(key string) {
	p.Lock()
	defer p.Unlock()
	service_name := filepath.Dir(key)
	service := p.services[service_name]
	if service == nil {
		logger.Debugf("no such handlermap %v.", service_name)
		return
	}

	for k := range service.clients {
		if service.clients[k].key == key {
			// deletion
			service.clients = append(service.clients[:k], service.clients[k + 1:]...)
			logger.Debugf("handlermap removed %v.", key)
			return
		}
	}
}

// provide a specific key for a handlermap, eg:
// path:/backends/snowflake, id:s1
//
// handlermap must be stored like /backends/xxx_service/xxx_id
func (p *service_pool) get_service_with_id(path string, id string) *grpc.ClientConn {
	p.RLock()
	defer p.RUnlock()
	service := p.services[path]
	if service == nil {
		return nil
	}

	if len(service.clients) == 0 {
		return nil
	}

	fullpath := string(path) + "/" + id
	for k := range service.clients {
		if service.clients[k].key == fullpath {
			return service.clients[k].conn
		}
	}

	return nil
}

func (p *service_pool) get_service(path string) *grpc.ClientConn {
	p.RLock()
	defer p.RUnlock()
	service := p.services[path]
	if service == nil {
		return nil
	}

	if len(service.clients) == 0 {
		return nil
	}
	idx := int(atomic.AddUint32(&service.idx, 1))
	return service.clients[idx % len(service.clients)].conn
}

// choose a handlermap randomly
func GetService(path string) *grpc.ClientConn {
	return _default_pool.get_service(path)
}

// get a specific handlermap instance with given path and id
func GetServiceWithId(path string, id string) *grpc.ClientConn {
	return _default_pool.get_service_with_id(path, id)
}

func GetServiceClient(path string) proto.ServiceClient {
	logger.Debugf("looking for grpc service node %s", path)
	conn := _default_pool.get_service(path)
	if conn != nil {
		return proto.NewServiceClient(conn)
	}
	return nil
}