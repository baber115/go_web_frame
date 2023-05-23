package framework

import (
	"github.com/spf13/cast"
	"strconv"
)

type IRequest interface {
	// URL中带的参数
	// 例: foo.com?a=1&b=bar&c[]=bar
	QueryInt(key string, def int) (int, string)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}

	// 路由匹配中的参数
	// 例：/book/:id

	// form表单中的参数
	FormInt(key string, def int) int
	FormString(key string, def string) string
	FormArr(key string, def []string) []string
	FormAll() map[string][]string

	// json body

	// xml body

	// 其他格式

	// 基础信息

	// header

	// cookie
}

// query url

func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		vlen := len(vals)
		if vlen > 0 {
			return cast.ToInt(vals[vlen-1]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		vlen := len(vals)
		if vlen > 0 {
			return vals[vlen-1], true
		}
	}
	return def, false
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

// 匹配路由参数
func (ctx *Context) Param(key string) interface{} {
	if ctx.params != nil {
		if val, ok := ctx.params[key]; ok {
			return val
		}
	}

	return nil
}

// 匹配路由中的变量
// subject/:id
func (ctx *Context) ParamInt(key string, def int) (int, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToInt(val), true
	}

	return def, false
}
