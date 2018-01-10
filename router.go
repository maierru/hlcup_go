package main

import (
	"github.com/buaazp/fasthttprouter"
)

// Setup routes to controllers
func Setup(Router *fasthttprouter.Router) {
	Router.GET("/", Index)
	Router.GET("/locations/:id", Item)
}
