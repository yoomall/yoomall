package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/middleware"
	"github.com/lazyfury/pulse/helper/utils"
	"github.com/spf13/viper"
)

type DefaultApp struct {
	*framework.App
	Config *viper.Viper
}

var _ framework.IApp = (*DefaultApp)(nil)

func (d *DefaultApp) Migrate() {
}

func (d *DefaultApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.RecoverHandlerFunc,
	}
}

func (d *DefaultApp) Register(router *framework.RouterGroup) {

	// health check
	router.GET("/health", func(ctx *gin.Context) {
		// TODO:collect other info if need
		ctx.JSON(http.StatusOK, map[string]any{"ok": true})
	}).Doc(&framework.DocItem{
		Method: http.MethodGet,
		Tag:    "app",
		Path:   "/health",
	})

	// 用例 1，统一图片域，在需要绘制 canvas 以及其他跨域情况时使用
	// 用例 2，服务器本地的 api 调用, 爬虫数据收集任务，无头浏览器，深度学习，想要了解一下，可能使用 python 或者 js 实现，提供局域网 api 调用
	router.GET("/proxy", func(ctx *gin.Context) {
		utils.ProxyRequest(ctx)
	}).Doc(&framework.DocItem{
		Method: http.MethodGet,
		Tag:    "app",
		Path:   "/proxy",
	})
}
