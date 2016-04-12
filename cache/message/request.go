package message

import "github.com/TriangleGo/gylib/service/action"

type Request struct {
	Action action.Action
	Params map[string]interface{}
}