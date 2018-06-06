package handlers

import (
	"node/model"
	"node/server"
)

type ListActionRequest struct {
	Timestamp int64
}

type ListActionResponse struct {
	model.BaseResp
	Timestamp int64
}

func ListActionHandler(key string) (response interface{}, err error) {
	response = &ListActionResponse{}
	listActionProcessor()
	return
}

func listActionProcessor() {
	server.ListActions()
}
