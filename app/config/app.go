package defappConfig

import (
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/app/middleware"
	"lazyfury.github.com/yoomall-server/app/model"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	coremiddleware "lazyfury.github.com/yoomall-server/core/middleware"
)

type DefaultApp struct {
	*core.AppImpl
	Config          *config.Config
	AuthMiddlewares []gin.HandlerFunc
}

var _ core.App = (*DefaultApp)(nil)

func (d *DefaultApp) Migrate() {
	d.GetDB().AutoMigrate(&model.User{})
}

func (d *DefaultApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.CORSMiddleware(),
		coremiddleware.RecoverHandlerFunc,
	}
}

func (d *DefaultApp) Register(router *gin.RouterGroup) {
	for _, handler := range d.Handlers {
		handler.Register(router.Group(handler.GetRouterGroupName()))
	}
}
