package main

import (
	"node/server"
	"node/handlers"
)

func main() {
	server.InitNode()
	handlers.InitHandlers()
	server.StartServerNode()
}
