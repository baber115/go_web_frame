package main

import "web_frame/framework"

func registerRouter(core *framework.Core) {
	//core.Get("foo", FooControllerHandler)
	// 静态路由
	core.Get("/user/login", UserControllerHandler)

	// 路由前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Get("/:id", SubjectControllerHandler)
		subjectApi.Put("/:id", SubjectControllerHandler)
		subjectApi.Delete("/:id", SubjectControllerHandler)
		subjectApi.Get("/list/all", SubjectControllerHandler)
	}
}
