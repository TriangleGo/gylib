package message

import "gyservice/action"

type Request struct {
	Action action.Action
	Params map[string]interface{}
}