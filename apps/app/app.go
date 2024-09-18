package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/core"
	coremiddleware "lazyfury.github.com/yoomall-server/core/middleware"
)

type DefaultApp struct {
	*core.App
	Config          *viper.Viper
	AuthMiddlewares []gin.HandlerFunc
}

var _ core.IApp = (*DefaultApp)(nil)

func (d *DefaultApp) Migrate() {
}

func (d *DefaultApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		coremiddleware.RecoverHandlerFunc,
	}
}

func (d *DefaultApp) Register(router *core.RouterGroup) {

}
