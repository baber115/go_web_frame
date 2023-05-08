package web_frame

import (
	"web_frame/framework"
)

func FooControllerHandler(ctx *framework.Context) error {
	return ctx.Json(200, map[string]interface{}{
		"code": 1,
	})
}

//func (c *framework.Context) ff() {
//	durationCtx, cancel := context.WithTimeout(c.Bas, time.Duration(1*time.Second))
//	defer cancel
//}
