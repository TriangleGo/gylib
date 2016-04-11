package handler

import (
	"gyservice/action"
	"golang.org/x/net/context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"gyservice/etcd"
	log "github.com/kyugao/go-logger/logger"
	"gyservice/respcode"
	"gycache/message"
	"fmt"
	"gyservice/proto"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	actionName := vars["action"]
	var respByte []byte

	actionCode, ok := action.ActionFromName(actionName)
	if !ok {
		tempResp := message.NewResponse()
		log.Debug("Unsupport function", actionName)
		tempResp.SetRespCode(respcode.RC_GENERAL_SYS_ERR)
		tempResp.SetParam("error", fmt.Sprintf("Action %s is not supported", actionName))
		respByte, _ = json.Marshal(tempResp)
	} else {
		r.ParseForm()
		params := make(map[string]interface{})
		for k, _ := range r.Form {
			params[k] = r.FormValue(k)
		}
		request := &message.Request{Action:actionCode, Params:params}

		key, err := message.CacheMsg(request)
		if err != nil {
			respByte, _ = json.Marshal(respcode.RC_GENERAL_SYS_ERR)
		} else {
			log.Debugf("Cached request key %s request: %v", key, request)
			client, _ := etcd.GetClient(actionCode)

			if client == nil {
				tempResp := message.NewResponse()
				log.Debug("Service node is unavailable.")
				tempResp.SetRespCode(respcode.RC_GENERAL_SYS_ERR)
				tempResp.SetParam("error", fmt.Sprintf("Service node for %s is unavailable.", actionName))
				respByte, _ = json.Marshal(tempResp)
			} else {
				cachedReq := &proto.Request{}
				cachedReq.Key = key
				clientResp, err := client.Serve(context.Background(), cachedReq)
				log.Debug("response from node:", clientResp, err)
				err = message.GetMsg(clientResp.Key, &respByte)
				if err != nil {
					log.Debugf("Get cache resp error: %s.", err.Error())
					respByte, _ = json.Marshal(respcode.RC_GENERAL_SYS_ERR)
				} else {
					log.Debugf("Get cache resp: %s.", string(respByte))
				}
			}
		}
	}

	w.Write(respByte)
}