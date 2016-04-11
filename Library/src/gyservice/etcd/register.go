package etcd

import (
	"github.com/coreos/go-etcd/etcd"
	"fmt"
	"github.com/kyugao/go-logger/logger"
)

func RegisterService() (err error) {
	err = _default_pool.registerService()
	return
}

func DeleteService() (err error) {
	err = _default_pool.deleteService()
	return
}

func (p *service_pool) registerService() (err error) {
	client := p.client_pool.Get().(*etcd.Client)
	defer func() {
		p.client_pool.Put(client)
	}()
	serviceDir := fmt.Sprintf("%s/%s", ServiceName, ServicePort)
	serviceKey := fmt.Sprintf("%s/%s/service", ServiceName, ServicePort)
	actionStartKey := fmt.Sprintf("%s/%s/actionStart", ServiceName, ServicePort)
	actionEndKey := fmt.Sprintf("%s/%s/actionEnd", ServiceName, ServicePort)
	val := fmt.Sprintf("%s:%s", ServiceIp, ServicePort)
	_, err = client.CreateDir(serviceDir, 0)
	_, err = client.Set(serviceKey, val, 0)
	_, err = client.Set(actionStartKey, fmt.Sprintf("%d",ActionStart), 0)
	_, err = client.Set(actionEndKey, fmt.Sprintf("%d",ActionEnd), 0)
	logger.Debugf("mkdir service %s, with err %v", serviceDir, err)
	logger.Debugf("set service %s, with err %v", serviceKey, err)
	logger.Debugf("set action start %s, with err %v", actionStartKey, err)
	logger.Debugf("set action end %s, with err %v", actionEndKey, err)
	return
}

func (p *service_pool) deleteService() (err error) {
	client := p.client_pool.Get().(*etcd.Client)
	defer func() {
		p.client_pool.Put(client)
	}()
	serviceDir := fmt.Sprintf("%s/%s/", ServiceName, ServicePort)
	_, err = client.Delete(serviceDir, true)
	logger.Debugf("delete %s, with err %v", serviceDir, err)
	return
}

