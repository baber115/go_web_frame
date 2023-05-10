package framework

type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

type Group struct {
	core   *Core
	prefix string
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func NewGroup(c *Core, prefix string) *Group {
	return &Group{
		core:   c,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string, handler ControllerHandler) {
	url := g.prefix + uri
	g.core.Get(url, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	url := g.prefix + uri
	g.core.Post(url, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	url := g.prefix + uri
	g.core.Put(url, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	url := g.prefix + uri
	g.core.Delete(url, handler)
}
