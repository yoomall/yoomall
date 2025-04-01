package appHandlerV2

import (
	"github.com/gin-gonic/gin"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/helper/response"
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
func (b *BaseHandlerV2) Register(router *framework.RouterGroup) {
	router.GET("/test", func(ctx *gin.Context) {
		response.Success(nil).Done(ctx)
	})
}

var _ framework.Handler = (*BaseHandlerV2)(nil)
