package main

import "web_frame/framework"

func SubjectLoginController(c *framework.Context) error {
	// 打印控制器名字
	c.SetStatus(200).Json("ok, SubjectLoginController")
	return nil
}
