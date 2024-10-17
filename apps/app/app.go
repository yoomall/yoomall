package app

import (
	"net/http"

	"yoomall/core"
	"yoomall/core/helper/utils"
	coremiddleware "yoomall/core/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type DefaultApp struct {
	*core.App
	Config *viper.Viper
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

	// 用例 1，统一图片域，在需要绘制 canvas 以及其他跨域情况时使用
	// 用例 2，服务器本地的 api 调用, 爬虫数据收集任务，无头浏览器，深度学习，想要了解一下，可能使用 python 或者 js 实现，提供局域网 api 调用
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Tag:    "app",
		Path:   "/proxy",
	}).GET("/proxy", func(ctx *gin.Context) {
		utils.ProxyRequest(ctx)
	})
}
