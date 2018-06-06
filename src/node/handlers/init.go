package handlers

import "node/server"

func InitHandlers() {
	server.RegisterHandler("echo", "/echo", EchoHandler)
	server.RegisterHandler("listAction", "/action/list", ListActionHandler)
}
