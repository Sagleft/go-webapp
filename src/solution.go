package main

import (
	"github.com/gin-gonic/gin"
)

type solution struct {
	Connections solConnections
}

func newSolution() *solution {
	return &solution{}
}

func (sol *solution) run() error {
	err := sol.setupRouter()
	if err != nil {
		return err
	}
	return nil
}

func (sol *solution) setupRouter() error {
	router := router{
		Gin:         gin.Default(),
		Connections: &sol.Connections,
	}
	err := router.setupRoutes()
	if err != nil {
		return err
	}
	return router.Gin.Run()
}
