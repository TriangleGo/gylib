package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"handler"
	"github.com/TriangleGo/gylib/logger"
	"fmt"
	"runtime"
	"github.com/TriangleGo/gylib/cache"
	"github.com/TriangleGo/gylib/service/etcd"
)

func init() {

	logger.InitLogger("./conf/logger.conf")
	cache.InitCache("./conf/cache.conf")
	etcd.InitEtcd("./conf/etcd.conf")
	r := mux.NewRouter()
	r.HandleFunc("/f/{action}", handler.FormHandler).Methods("POST")
	r.HandleFunc("/file/upload", handler.UploadHandler).Methods("POST")
	r.HandleFunc("/file/download/{fileId}", handler.DownloadHandler)
	http.Handle("/", r)
}

func main() {

	logger.Infof("====== Start agent service node @ %s ======", etcd.ServicePort)

	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 1 << 20)
			runtime.Stack(buf, true)
			logger.Debug(buf)
		}
	}()
	port := fmt.Sprintf(":%s", etcd.ServicePort)
	http.ListenAndServe(port, nil)
}
