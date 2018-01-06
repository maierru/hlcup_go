package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// Server setup
func Server() {
	Router := fasthttprouter.New()
	Setup(Router)
	fasthttp.ListenAndServe(":8080", Router.Handler)
}
