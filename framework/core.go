package framework

import (
	"log"
	"net/http"
)

// Core 框架核心结构
type Core struct {
	route map[string]ControllerHandler
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	return &Core{
		route: map[string]ControllerHandler{},
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.route[url] = handler
}

// ServerHTTP 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.ServerHttp")
	ctx := NewContext(request, response)

	route := c.route["foo"]
	if route == nil {
		return
	}
	log.Println("core.route")
	route(ctx)
}
