package server

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/BurntSushi/toml"
	"net"
	"logger"
	"node/proto"
	"util/json"
	"node/message"
)

type Config struct {
	Name    string
	Address string
}

type Node struct {
	Name     string
	Listener net.Listener
	Actions  map[string]string
	Handlers map[string]func(string) (interface{}, error)
}

var node *Node
var config *Config

func init() {
	node = &Node{
		Actions: make(map[string]string),
		Handlers: make(map[string]func(string) (interface{}, error)),
	}
	config = &Config{}
}

/*
 * 初始化 Node
 */
func InitNode() {
	loadConfig()
	node.Name = config.Name
}

/*
 * 加载 Node 配置文件
 */
func loadConfig() {
	_, err := toml.DecodeFile("./conf/node.conf", config)
	if err != nil {
		panic(err)
	} else {
		logger.Debug(json.ToString(config))
	}
}

func StartServerNode() {
	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		panic(err)
	}
	node.Listener = lis
	logger.Infof("Start %s server node at %s.", config.Name, config.Address)
	grpcServer := grpc.NewServer()
	proto.RegisterNodeServer(grpcServer, node)
	grpcServer.Serve(node.Listener)
}

func NodeName() (name string) {
	name = node.Name
	return
}

func ListActions() (actions []interface{}) {
	for action, uri := range node.Actions {
		var item = struct {
			Action string `json:"action"`
			Uri    string `json:"uri"`
		}{
			Action: action,
			Uri:    uri,
		}
		actions = append(actions, item)
	}
	logger.Infof(json.ToString(actions))
	return
}

func RegisterHandler(action string, uri string, handler func(string) (interface{}, error)) {
	logger.Infof("Register handler action:%s.", action)
	node.Actions[action] = uri
	node.Handlers[action] = handler
}

func (s *Node) Call(ctx context.Context, inMsg *proto.Message) (outMsg *proto.Message, err error) {
	outMsg = &proto.Message{}
	logger.Debugf("called %s.", json.ToString(inMsg))

	handler := s.Handlers[inMsg.Action]
	resp, err := handler(inMsg.Key)
	if err != nil {
		logger.Errorf("%s handler return error %v.", inMsg.Action, err)
		return
	}

	outKey, err := message.CacheMsg(resp)
	if err != nil {
		logger.Errorf("cache %s response message error %v.", inMsg.Action, err)
		return
	}

	outMsg.Action = inMsg.Action
	outMsg.Key = outKey
	return
}
