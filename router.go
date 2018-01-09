package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/maierru/hlcup_go/controllers/locations"
	"github.com/maierru/hlcup_go/controllers/root"
)

// Setup routes to controllers
func Setup(Router *fasthttprouter.Router) {
	Router.GET("/", root.Index)
	Router.GET("/locations/:id", locations.Item)
}
