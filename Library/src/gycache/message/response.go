package message

import "gyservice/respcode"

type Response struct {
	respcode.RespCode
	Params map[string]interface{}
}

func NewResponse() (resp *Response) {
	resp = &Response{}
	resp.Params = make(map[string]interface{})
	return
}

func (resp *Response)SetRespCode(code *respcode.RespCode) {
	resp.Code = code.Code
	resp.Info = code.Info
}

func (resp *Response)SetParam(key string, val interface{}) {
	resp.Params[key] = val
}