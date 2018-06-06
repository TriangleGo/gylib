package service

import (
	"logger"
	"service/action"
	"examples/FileServer/service/handler"
	sServer "service/server"
)

func InitServer() *sServer.ServiceServer {
	logger.Info("Register file handlers.")
	server := sServer.NewServiceServer()
	server.RegHandler(action.Action_LoadFile, handler.LoadFileHandler)
	server.RegHandler(action.Action_SaveFile, handler.SaveFileHandler)
	server.RegHandler(action.Action_DeleteFile, handler.DeleteFileHandler)
	return server
}
