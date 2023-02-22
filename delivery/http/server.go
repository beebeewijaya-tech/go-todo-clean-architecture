package http

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server struct {
	Config *viper.Viper
	Router *gin.Engine
}

func NewServer(c *viper.Viper) *Server {
	h := &Server{Config: c}

	router := gin.Default()

	h.Router = router

	return h
}

func (s *Server) Start(address string) error {
	return s.Router.Run(address)
}
