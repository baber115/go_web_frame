package route

import "web_frame/framework"

type IGroup interface {
	Get(string, framework.ControllerHandler)
	Post(string, framework.ControllerHandler)
	Put(string, framework.ControllerHandler)
	Delete(string, framework.ControllerHandler)
}

type Group struct {
	core   *Core
	prefix string
}

func (g *Group) Group(prefix string) IGroup {
	return NewGroup(g.core, prefix)
}

func NewGroup(c *Core, prefix string) *Group {
	return &Group{
		core:   c,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string, handler framework.ControllerHandler) {
	url := g.prefix + uri
	g.core.Get(url, handler)
}

func (g *Group) Post(uri string, handler framework.ControllerHandler) {
	url := g.prefix + uri
	g.core.Post(url, handler)
}

func (g *Group) Put(uri string, handler framework.ControllerHandler) {
	url := g.prefix + uri
	g.core.Put(url, handler)
}

func (g *Group) Delete(uri string, handler framework.ControllerHandler) {
	url := g.prefix + uri
	g.core.Delete(url, handler)
}
