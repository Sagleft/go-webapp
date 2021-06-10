package main

import "github.com/gin-gonic/gin"

type router struct {
	Gin         *gin.Engine
	Connections *solConnections
}

func (r *router) setupRoutes() error {
	// TODO
	return nil
}
