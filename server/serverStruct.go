package server

import "github.com/gin-gonic/gin"

type Server struct {
	host         string
	port         string
	serverEngine *gin.Engine
}

func NewServer() *Server {
	newServer := &Server{
		host:         "",
		port:         "",
		serverEngine: gin.Default(),
	}
	return newServer
}

func (s *Server) SetPort(port string) {
	s.port = port
}

func (s *Server) SetHost(host string) {
	s.host = host
}
