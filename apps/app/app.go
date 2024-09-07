package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/app/model"
	"lazyfury.github.com/yoomall-server/core"
	coremiddleware "lazyfury.github.com/yoomall-server/core/middleware"
)

type DefaultApp struct {
	*core.AppImpl
	Config          *viper.Viper
	AuthMiddlewares []gin.HandlerFunc
}

var _ core.App = (*DefaultApp)(nil)

func (d *DefaultApp) Migrate() {
	d.GetDB().AutoMigrate(&model.User{}, &model.UserToken{})
}

func (d *DefaultApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		coremiddleware.RecoverHandlerFunc,
	}
}

func (d *DefaultApp) Register(router *gin.RouterGroup) {
	for _, handler := range d.Handlers {
		handler.Register(router.Group(handler.GetRouterGroupName()))
	}
}
