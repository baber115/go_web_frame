package framework

import (
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router map[string]map[string]ControllerHandler
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	// 二级路由
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}
	// 一级路由
	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter

	return &Core{
		router: router,
	}
}

// 寻找静态路由
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	url := strings.ToUpper(request.URL.Path)
	method := strings.ToUpper(request.Method)
	if methodHandlers, ok := c.router[method]; ok {
		if handler, ok := methodHandlers[url]; ok {
			return handler
		}
	}
	return nil
}

func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}

func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}

func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}

// ServerHTTP 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request, response)

	// 找路由
	router := c.FindRouteByRequest(request)
	// 没找到报404
	if router == nil {
		ctx.Json(404, "not found")
		return
	}

	// 内部错误直接报错
	if err := router(ctx); err != nil {
		ctx.Json(500, "server error")
		return
	}
}
