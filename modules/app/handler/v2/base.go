package appHandlerV2

import (
	"yoomall/yoo"
	"yoomall/yoo/helper/response"

	"github.com/gin-gonic/gin"
)

type BaseHandlerV2 struct {
}

func NewBaseHandler() *BaseHandlerV2 {
	return &BaseHandlerV2{}
}

// GetRouterGroupName implements yoo.Handler.
func (b *BaseHandlerV2) GetRouterGroupName() string {
	return "basev2"
}

// Register implements yoo.Handler.
func (b *BaseHandlerV2) Register(router *yoo.RouterGroup) {
	router.GET("/test", func(ctx *gin.Context) {
		response.Success(nil).Done(ctx)
	})
}

var _ yoo.Handler = (*BaseHandlerV2)(nil)
