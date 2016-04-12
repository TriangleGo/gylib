package service

import (
	"github.com/TriangleGo/gylib/logger"
	"github.com/TriangleGo/gylib/service/action"
	"service/handler"
	"github.com/TriangleGo/gylib/service/service"
)

func InitServer() *service.ServiceServer {
	logger.Info("Register file handlers.")
	server := service.NewServiceServer()
	server.RegHandler(action.Action_LoadFile, handler.LoadFileHandler)
	server.RegHandler(action.Action_SaveFile, handler.SaveFileHandler)
	server.RegHandler(action.Action_DeleteFile, handler.DeleteFileHandler)
	server.RegHandler(action.Action_CheckAppVersion, handler.CheckAppVersion)
	return server
}
