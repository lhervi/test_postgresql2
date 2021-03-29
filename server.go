package main

import (
	"github.com/gin-gonic/gin"
)

const (
	serverAddress = "localhost:3020"
)

type TyniServer struct {
	handlers ServerHandlers
}

func NewTyniServer(handlers ServerHandlers) *TyniServer {
	return &TyniServer{
		handlers: handlers,
	}
}

func (s *TyniServer) Init() error {
	router := gin.Default()
	s.handlers.InitRoutes(router)

	err := router.Run(serverAddress)
	return err
}
