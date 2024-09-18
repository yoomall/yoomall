package core

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// handler 在 app 内注册的，注释写文档没有办法推理路由前缀
func (r *RouterGroup) WithDoc(doc *DocItem, fn func(ctx *gin.Context)) *RouterGroup {
	// register handler
	r.Handle(doc.Method, doc.Path, fn)

	// write doc
	basePath := reflect.ValueOf(r).Elem().FieldByName("basePath")
	doc.Path = basePath.String() + doc.Path
	DocInstance.Add(doc)

	return r
}

// GROUP
func (r *RouterGroup) Group(path string) *RouterGroup {
	return &RouterGroup{RouterGroup: r.RouterGroup.Group(path)}
}

// USE
func (r *RouterGroup) Use(middleware ...gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.Use(middleware...)
	return r
}

// GET
func (r *RouterGroup) GET(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.GET(path, handler)
	return r
}

// POST
func (r *RouterGroup) POST(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.POST(path, handler)
	return r
}

// PUT
func (r *RouterGroup) PUT(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.PUT(path, handler)
	return r
}

// DELETE
func (r *RouterGroup) DELETE(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.DELETE(path, handler)
	return r
}

// PATCH
func (r *RouterGroup) PATCH(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.PATCH(path, handler)
	return r
}

// Handle
func (r *RouterGroup) Handle(method, path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.Handle(method, path, handler)
	return r
}
