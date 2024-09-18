package core

type Handler interface {
	Register(router *RouterGroup)
	GetRouterGroupName() string
}
