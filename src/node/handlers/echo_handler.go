package handlers

import (
	"node/message"
	"node/model"
	"time"
	"node/server"
)

type EchoRequest struct {
	Timestamp int64
}

type EchoResponse struct {
	model.BaseResp
	NodeName  string
	Timestamp int64
}

func EchoHandler(key string) (response interface{}, err error) {
	request := &EchoRequest{}
	err = message.GetMsg(key, request)
	response, err = echoProcessor(request)
	return
}

func echoProcessor(request *EchoRequest) (response *EchoResponse, err error) {
	response = &EchoResponse{}
	response.SetRC(model.RC_SUCC)
	response.NodeName = server.NodeName()
	response.Timestamp = time.Now().Unix()
	return
}
