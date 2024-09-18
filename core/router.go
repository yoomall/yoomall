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
