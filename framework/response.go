package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"net/http"
	"net/url"
)

type IResponse interface {
	// json
	Json(obj interface{}) IResponse

	// Jsonp输出
	Jsonp(obj interface{}) IResponse

	//xml输出
	Xml(obj interface{}) IResponse

	// html输出
	Html(file string, obj interface{}) IResponse

	// string
	Text(format string, values ...interface{}) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse

	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	SetStatus(code int) IResponse

	// 设置200状态
	SetOkStatus() IResponse
}

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.Request == nil {
		return errors.New("ctx.request is empty")
	}
	// 读取文本，这里读取是一次性的，需要在下面重新赋值
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	// 重新填充request.body，为后续二次读做准备
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// 解析到obj里面
	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *Context) Jsonp(obj interface{}) IResponse {
	// 获取请求参数
	callbackfunc, _ := ctx.QueryString("callback", "callback_func")
	ctx.SetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成 XSS 攻击
	callback := template.JSEscapeString(callbackfunc)
	// 输出函数名
	_, err := ctx.ResponseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}
	// 输出左括号
	_, err = ctx.ResponseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}
	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	_, err = ctx.ResponseWriter.Write(ret)
	if err != nil {
		return ctx
	}
	// 输出右括号
	_, err = ctx.ResponseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx
}

func (ctx *Context) Json(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.ResponseWriter.Write(byt)
	return ctx
}

// html 输出
func (ctx *Context) Html(file string, obj interface{}) IResponse {
	// 读取模版文件，创建 template 实例
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	// 执行 Execute 方法将 obj 和模版进行结合
	if err := t.Execute(ctx.ResponseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Content-Type", "application/html")
	return ctx
}

func (ctx *Context) Text(format string, values ...interface{}) IResponse {
	return nil
}

func (ctx *Context) Xml(obj interface{}) IResponse {
	return nil
}

// 重定向
func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.ResponseWriter, ctx.Request, path, http.StatusMovedPermanently)
	return ctx
}

// header
func (ctx *Context) SetHeader(key string, val string) IResponse {
	ctx.ResponseWriter.Header().Add(key, val)
	return ctx
}

// Cookie
func (ctx *Context) SetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.ResponseWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

// 设置状态码
func (ctx *Context) SetStatus(code int) IResponse {
	ctx.ResponseWriter.WriteHeader(code)
	return ctx
}

// 设置200状态
func (ctx *Context) SetOkStatus() IResponse {
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	return ctx
}
