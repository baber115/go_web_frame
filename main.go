package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	go func() {
		server.ListenAndServe()
	}()

	// 当前的 Goroutine 等待信号量
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	server.RegisterOnShutdown(func() {
		fmt.Println("server.RegisterOnShutdown")
	})
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown", err)
	}
}
