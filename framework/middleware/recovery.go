package middleware

import (
	"fmt"
	"web_frame/framework"
)

func Recovery() framework.ControllerHandler {
	fmt.Println("middleware recovery")
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.SetStatus(500).Json(err)
			}
		}()
		c.Next()

		return nil
	}
}
