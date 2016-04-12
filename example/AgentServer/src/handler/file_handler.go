package handler

import (
	"net/http"
	"github.com/TriangleGo/gylib/logger"
	"io/ioutil"
	cFile "github.com/TriangleGo/gylib/cache/file"
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"fmt"
	"github.com/TriangleGo/gylib/service/respcode"
	"github.com/TriangleGo/gylib/cache/message"
	"github.com/TriangleGo/gylib/service/etcd"
	"github.com/TriangleGo/gylib/service/action"
	"github.com/TriangleGo/gylib/service/proto"
)

// 5MB
const MAX_MEMORY = 5 * 1024 * 1024

/*
Receive upload files, save all the file information into cache server.
 */
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	agentResp := message.NewResponse()

	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		agentResp.SetRespCode(respcode.RC_GENERAL_SYS_ERR)
		agentResp.SetParam("error", "The file uploaded exceeded the limitation 5M.")
	} else {
		//tokenStr := r.MultipartForm.Value["token"][0]
		// cache file information
		fileHeaders := r.MultipartForm.File["file"]

		// store cached file ids.
		if len(fileHeaders) > 0 {
			fileHeader := fileHeaders[0]
			filename := fileHeader.Filename
			fileType := fileHeader.Header["Content-Type"][0]
			file, _ := fileHeader.Open()
			content, _ := ioutil.ReadAll(file)
			fileId := cFile.NewCacheFile(filename, fileType, content, 1)
			agentResp.SetParam("fileId", fileId)
			agentResp.SetRespCode(respcode.RC_GENERAL_SUCC)
		} else {
			agentResp.SetRespCode(respcode.RC_GENERAL_APP_ERR)
			agentResp.SetParam("error", "Dummy input file.")
		}

	}
	respByte, _ := json.Marshal(agentResp)
	w.Write(respByte)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId := vars["fileId"]
	name, contentType, content, exists := cFile.GetCacheFile(fileId, false)

	if !exists {
		logger.Debugf("load file %s request to file node", fileId)
		client, _ := etcd.GetClient(action.Action_LoadFile)

		if client == nil {
			respBytes, _ := json.Marshal(respcode.RC_SERVICE_UNAVAILABLE)
			w.Write(respBytes)
			return
		}

		params := map[string]interface{}{"fileId":fileId}
		request := &message.Request{Action:action.Action_LoadFile, Params:params}
		key, err := message.CacheMsg(request)
		cachedReq := &proto.Request{key}

		clientResp, err := client.Serve(context.Background(), cachedReq)
		logger.Debug("clientResp:", clientResp)
		if err == nil {
			respObj := &message.Response{}
			err = message.GetMsg(clientResp.Key, respObj)
			if err == nil {
				fileId, _ := respObj.Params["fileId"].(string)
				name, contentType, content, exists = cFile.GetCacheFile(fileId, false)
			} else {
				http.NotFoundHandler()
			}
		}
	} else {
		logger.Debug("exists in cache")
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	w.Header().Add("Content-type", contentType)
	w.Write(content)
}