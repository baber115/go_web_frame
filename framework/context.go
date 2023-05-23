package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter

	ctx context.Context
	// 当前请求的handler
	handler    []ControllerHandler
	index      int // 当前请求调用到调用链的哪个节点
	hasTimeout bool
	writeMutex *sync.Mutex

	// 路由参数
	params map[string]string
}

func NewContext(req *http.Request, resWriter http.ResponseWriter) *Context {
	return &Context{
		Request:        req,
		ResponseWriter: resWriter,
		ctx:            req.Context(),
		writeMutex:     &sync.Mutex{},
		index:          -1,
	}
}

func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handler) {
		if err := ctx.handler[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) SetHandler(handler []ControllerHandler) {
	ctx.handler = handler
}

func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

// base function

func (ctx *Context) WriteMux() *sync.Mutex {
	return ctx.writeMutex
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.Request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.ResponseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// implement context.context

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

func (ctx *Context) DeadLine() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}
