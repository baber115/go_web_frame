package framework

type IResponse interface {
	// json
	Json(obj interface{}) IResponse

	// jsonp
	Jsonp(obj interface{}) IResponse

	// xml

	// html

	// string

	// 重定向

	// header

	// cookie

	// 设置状态码

	// 设置200状态码
}
