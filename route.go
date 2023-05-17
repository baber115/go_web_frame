package main

import (
	"time"
	"web_frame/framework/middleware"
	"web_frame/framework/route"
)

func RegisterRouter(core *route.Core) {
	core.Get("foo", FooControllerHandler)
	// 静态路由
	core.Get("/user/login", middleware.TimeoutHandler(UserLoginController, time.Second))

	// 路由前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Get("/:id", SubjectLoginController)
		subjectApi.Put("/:id", SubjectLoginController)
		subjectApi.Delete("/:id", SubjectLoginController)
		subjectApi.Get("/list/all", SubjectLoginController)
	}
}
