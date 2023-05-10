package main

import (
	"context"
	"fmt"
	"time"
	"web_frame/framework"
)

func FooControllerHandler(ctx *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		time.Sleep(10 * time.Second)
		ctx.Json(200, "ok")
		// 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()
		fmt.Println(p)
		ctx.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()
		ctx.Json(500, "timeout")
		ctx.SetHasTimeout()
	}
	return nil
}
