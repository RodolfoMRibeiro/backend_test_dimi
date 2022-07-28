package server

import (
	"transaction/configs"
)

func CreateServer(envVars configs.EnvironmentVariables) *Server {
	server := NewServer()

	server.SetPort(envVars.ServerPort)
	server.SetHost(envVars.ServerHost)
	return server
}
