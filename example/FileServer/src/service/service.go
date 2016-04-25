package service

import (
	"github.com/TriangleGo/gylib/logger"
	"github.com/TriangleGo/gylib/service/action"
	"service/handler"
	sServer "github.com/TriangleGo/gylib/service/server"
)

func InitServer() *sServer.ServiceServer {
	logger.Info("Register file handlers.")
	server := sServer.NewServiceServer()
	server.RegHandler(action.Action_LoadFile, handler.LoadFileHandler)
	server.RegHandler(action.Action_SaveFile, handler.SaveFileHandler)
	server.RegHandler(action.Action_DeleteFile, handler.DeleteFileHandler)
	return server
}
