package main

import (
	"net/http"
	"web_frame/framework"
)

func main() {
	core := framework.NewCore()
	RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
