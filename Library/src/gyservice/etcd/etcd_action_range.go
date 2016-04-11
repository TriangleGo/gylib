package etcd

import (
	"gylogger"
	"gyservice/proto"
	"gyservice/action"
)

func AddNode(start int, end int, name string) {
	serviceRange := &service_range{
		Node:name,
		Start:start,
		End:end,
	}
	pool[start] = serviceRange
	secPool[name] = start
}

func DeleteNode(name string) {
	start, ok := secPool[name]
	if ok {
		delete(secPool, name)
		delete(pool, start)
	}

}

func GetClient(actionCode action.Action) (client proto.ServiceClient, serviceNodeName string) {
	code := int(actionCode)
	for key, val := range pool {
		logger.Debugf("check action code:name %d:%s, get pool element key:%d, start:end:name %d:%d:%s.", code, actionCode, key, val.Start, val.End, val.Node)
		if code >= key && code <= val.End {
			client = GetServiceClient(val.Node)
			serviceNodeName = val.Node
			logger.Debugf("got client %v", client)
			break;
		}
	}
	return
}