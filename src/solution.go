package main

import "github.com/gin-gonic/gin"

type solution struct {
	Connections solConnections
}

func newSolution() (*solution, error) {
	sol := solution{}
	err := sol.setupRouter()
	if err != nil {
		return nil, err
	}
	return &sol, nil
}

func (sol *solution) run() {
	//
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
	return nil
}
