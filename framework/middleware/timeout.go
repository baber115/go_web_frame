package middleware

import "C"
import (
	"context"
	"fmt"
	"log"
	"time"
	"web_frame/framework"
)

func TimeoutHandler(duration time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
		finishChan := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		// 初始化context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), duration)
		defer cancel()
		c.Request.WithContext(durationCtx)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			// 具体的业务逻辑
			c.Next()
			finishChan <- struct{}{}
		}()
		// 执行业务逻辑后的操作
		select {
		case p := <-panicChan:
			log.Println(p)
			c.ResponseWriter.WriteHeader(500)
		case <-finishChan:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			c.ResponseWriter.Write([]byte("timeout"))
		}

		return nil
	}
}
