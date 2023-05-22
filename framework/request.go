package framework

import "github.com/spf13/cast"

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
