package server

import (
	config "transaction/configs"

	"github.com/gin-gonic/gin"
)

type Server struct {
	host         string
	port         string
	serverEngine *gin.Engine
}

func CreateServer() *Server {
	server := NewServer()

	server.SetPort(config.Server.PORT)
	server.SetHost(config.Server.HOST)
	return server
}

func NewServer() *Server {
	newServer := &Server{
		host:         "",
		port:         "",
		serverEngine: gin.Default(),
	}
	return newServer
}

func (s Server) GetServerEngine() *gin.Engine {
	return s.serverEngine
}

func (s *Server) SetPort(port string) {
	s.port = port
}

func (s *Server) SetHost(host string) {
	s.host = host
}
