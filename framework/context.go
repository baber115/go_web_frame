package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
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

// query url

func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		vlen := len(vals)
		if vlen > 0 {
			intval, err := strconv.Atoi(vals[vlen-1])
			if err == nil {
				return intval
			}
		}
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		vlen := len(vals)
		if vlen > 0 {
			return vals[vlen-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.Request != nil {
		return ctx.Request.URL.Query()
	}

	return map[string][]string{}
}

// form post

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		vlen := len(vals)
		if vlen > 0 {
			intval, err := strconv.Atoi(vals[vlen-1])
			if err == nil {
				return intval
			}
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		vlen := len(vals)
		if vlen > 0 {
			return vals[vlen-1]
		}
	}
	return def
}

func (ctx *Context) FormArr(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.Request != nil {
		return ctx.Request.PostForm
	}

	return map[string][]string{}
}

// application/json post

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.Request != nil {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			return err
		}
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}

	return nil
}

// response

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(500)
		return err
	}
	_, err = ctx.ResponseWriter.Write(byt)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(500)
		return err
	}
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}
