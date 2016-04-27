package caller

import (
	"service/action"
	"cache/message"
	"service/respcode"
	"fmt"
	"encoding/json"
	"logger"
	"service/etcd"
	"service/proto"
	"golang.org/x/net/context"
)

func CallObject(actionCode action.Action, params map[string]interface{}) (response *message.Response, err error) {

	respByte, err := CallBytes(actionCode, params)
	if err != nil {
		return
	}
	response = &message.Response{}
	err = json.Unmarshal(respByte, &response)
	return
}

func CallBytes(actionCode action.Action, params map[string]interface{}) (respByte []byte, err error) {

	request := &message.Request{Action:actionCode, Params:params}
	key, err := message.CacheMsg(request)
	if err != nil {
		respByte, _ = json.Marshal(respcode.RC_GENERAL_SYS_ERR)
	} else {
		logger.Debugf("Cached request key %s request: %v", key, request)
		client, _ := etcd.GetClient(actionCode)

		if client == nil {
			tempResp := message.NewResponse()
			logger.Debug("Service node is unavailable.")
			tempResp.SetRespCode(respcode.RC_GENERAL_SYS_ERR)
			tempResp.SetParam("error", fmt.Sprintf("Service node for %s is unavailable.", actionCode.Name()))
			respByte, _ = json.Marshal(tempResp)
		} else {
			cachedReq := &proto.Request{}
			cachedReq.Key = key
			var clientResp *proto.Response
			clientResp, err = client.Serve(context.Background(), cachedReq)
			logger.Debug("response from node:", clientResp, err)
			err = message.GetMsg(clientResp.Key, &respByte)
			if err != nil {
				logger.Debugf("Get cache resp error: %s.", err.Error())
				respByte, _ = json.Marshal(respcode.RC_GENERAL_SYS_ERR)
			} else {
				logger.Debugf("Get cache resp: %s.", string(respByte))
			}
		}
	}
	return
}