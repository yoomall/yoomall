package yoo

type Handler interface {
	Register(router *RouterGroup)
	GetRouterGroupName() string
}
