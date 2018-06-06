package handler

import (
	"service/action"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"logger"
	"service/respcode"
	"cache/message"
	"fmt"
	"service/caller"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	actionName := vars["action"]
	var respByte []byte

	actionCode, ok := action.ActionFromName(actionName)
	if !ok {
		tempResp := message.NewResponse()
		logger.Debug("Unsupport function1111111", actionName)
		tempResp.SetRespCode(respcode.RC_GENERAL_SYS_ERR)
		tempResp.SetParam("error", fmt.Sprintf("Action %s is not supported", actionName))
		respByte, _ = json.Marshal(tempResp)
	} else {
		r.ParseForm()
		params := make(map[string]interface{})
		for k, _ := range r.Form {
			params[k] = r.FormValue(k)
		}
		respByte, _ = caller.CallBytes(actionCode, params)
	}

	w.Write(respByte)
}