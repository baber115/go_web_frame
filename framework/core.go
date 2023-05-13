package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router map[string]*Tree
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{
		router: router,
	}
}

// 寻找静态路由
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	url := strings.ToUpper(request.URL.Path)
	method := strings.ToUpper(request.Method)
	if methodHandlers, ok := c.router[method]; ok {
		return methodHandlers.FindHandler(url)
	}
	return nil
}

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add GET router error: ", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add POST router error: ", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add PUT router error: ", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add DELETE router error: ", err)
	}
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
