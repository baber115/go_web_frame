package main

import (
	"web_frame/framework"
)

func RegisterRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
	// 静态路由
	core.Get("/user/login", UserLoginController)

	// 路由前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Get("/:id", SubjectLoginController)
		subjectApi.Put("/:id", SubjectLoginController)
		subjectApi.Delete("/:id", SubjectLoginController)
		subjectApi.Get("/list/all", SubjectLoginController)
	}
}
