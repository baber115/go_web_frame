package main

import (
	"net/http"
	"web_frame/framework/route"
)

func main() {
	core := route.NewCore()
	RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
