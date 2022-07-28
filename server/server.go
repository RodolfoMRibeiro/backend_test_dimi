package server

import (
	config "transaction/configs"
)

func CreateServer() *Server {
	server := NewServer()

	server.SetPort(config.Server.PORT)
	server.SetHost(config.Server.HOST)
	return server
}
