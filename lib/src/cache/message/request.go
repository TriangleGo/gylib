package message

import "service/action"

type Request struct {
	Action action.Action
	Params map[string]interface{}
}