package main

import (
	"net/http"
	"web_frame/framework/middleware"
	"web_frame/framework/route"
)

func main() {
	core := route.NewCore()
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
