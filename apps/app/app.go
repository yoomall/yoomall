package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/helper/utils"
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

	// health check
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Tag:    "app",
		Path:   "/health",
	}).GET("/health", func(ctx *gin.Context) {
		// TODO:collect other info if need
		ctx.JSON(200, map[string]any{"ok": true})
	})

	router.GET("/proxy", func(ctx *gin.Context) {
		utils.ProxyRequest(ctx)
	})
}
