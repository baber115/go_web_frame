package route

import "web_frame/framework"

type IGroup interface {
	Get(string, ...framework.ControllerHandler)
	Post(string, ...framework.ControllerHandler)
	Put(string, ...framework.ControllerHandler)
	Delete(string, ...framework.ControllerHandler)

	Group(string) IGroup

	Use(middlewares ...framework.ControllerHandler)
}

type Group struct {
	core        *Core
	prefix      string
	parent      *Group
	middlewares []framework.ControllerHandler
}

func (g *Group) Group(prefix string) IGroup {
	cgroup := NewGroup(g.core, prefix)
	cgroup.parent = g
	return cgroup
}

func (g *Group) Use(middlewares ...framework.ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func NewGroup(c *Core, prefix string) *Group {
	return &Group{
		core:   c,
		prefix: prefix,
	}
}

func (g *Group) GetMiddleware() []framework.ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}

	return append(g.parent.GetMiddleware(), g.middlewares...)
}

func (g *Group) Get(uri string, handler ...framework.ControllerHandler) {
	url := g.prefix + uri
	allHandlers := append(g.middlewares, handler...)
	g.core.Get(url, allHandlers...)
}

func (g *Group) Post(uri string, handler ...framework.ControllerHandler) {
	url := g.prefix + uri
	allHandlers := append(g.middlewares, handler...)
	g.core.Post(url, allHandlers...)
}

func (g *Group) Put(uri string, handler ...framework.ControllerHandler) {
	url := g.prefix + uri
	allHandlers := append(g.middlewares, handler...)
	g.core.Put(url, allHandlers...)
}

func (g *Group) Delete(uri string, handler ...framework.ControllerHandler) {
	url := g.prefix + uri
	allHandlers := append(g.middlewares, handler...)
	g.core.Delete(url, allHandlers...)
}
