package main

import (
	"time"
	"web_frame/framework"
)

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	// 等待10秒
	time.Sleep(time.Second * 10)
	// 输出结果
	c.SetStatus(200).Json("ok, UserLoginController: " + foo)
	return nil
}
