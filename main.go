package main

import (
	"net/http"
	"web_frame/framework"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    ":8080",
	}
	server.ListenAndServe()
	//http.ListenAndServe()
}
