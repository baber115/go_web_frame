package web_frame

import "web_frame/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
