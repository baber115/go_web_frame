package main

import (
	"web_frame/framework/middleware"
	"web_frame/framework/route"
)

func RegisterRouter(core *route.Core) {
	core.Get("foo", FooControllerHandler)
	// 静态路由
	core.Get("/user/login", middleware.Test1(), UserLoginController)

	// 路由前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test2())
		subjectApi.Get("/:id", SubjectLoginController)
		subjectApi.Put("/:id", SubjectLoginController)
		subjectApi.Delete("/:id", SubjectLoginController)
		subjectApi.Get("/list/all", SubjectLoginController)
	}
}
