package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/app"
	"lazyfury.github.com/yoomall-server/apps/auth"
	"lazyfury.github.com/yoomall-server/apps/common"
	commonservice "lazyfury.github.com/yoomall-server/apps/common/service"
	"lazyfury.github.com/yoomall-server/apps/post"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/constants"
	"lazyfury.github.com/yoomall-server/core/driver"
	httpserver "lazyfury.github.com/yoomall-server/core/http"
	coremiddleware "lazyfury.github.com/yoomall-server/core/middleware"
)

func NewHttpServer(
	config *viper.Viper,

	app *app.DefaultApp,
	auth *auth.AuthApp,
	postApp *post.PostApp,
	commonApp *common.CommonApp,

	noufoundRecordService *commonservice.NotFoundRecordService,

	doc *core.Doc,
) httpserver.HttpServer {
	engine := gin.Default()

	setup(engine)

	engine.GET("", func(ctx *gin.Context) {
		ctx.String(200, ":) yoomall server is running.")
	})

	engine.NoRoute(func(ctx *gin.Context) {
		noufoundRecordService.Add(ctx.Request.URL.Path, ctx.Request)
		ctx.JSON(http.StatusOK, gin.H{"message": "welcome."})
	})

	v1 := &core.RouterGroup{
		RouterGroup: engine.Group("/api/v1"),
	}
	v1.GET("/docs/api.json", doc.Handler)

	var apps = []*core.RegisterApp{
		{App: app, Router: v1.Group("")},
		{App: auth, Router: v1.Group("/auth")},
		{App: postApp, Router: v1.Group("/post")},
		{App: commonApp, Router: v1.Group("/common")},
	}

	for _, app := range apps {
		app.Register()
	}

	return httpserver.HttpServer{
		Engine: engine,
		Config: config,
	}
}

func NewDoc() *core.Doc {
	return core.NewDoc()
}

func NewDB(config *viper.Viper) *driver.DB {
	return driver.NewDB(config.GetString(constants.MYSQL_DSN))
}

func setup(engine *gin.Engine) {
	engine.SetTrustedProxies(nil)               //设置允许请求的域名
	engine.Use(coremiddleware.CORSMiddleware()) // 跨域
	engine.Use(gin.Recovery())                  // 错误恢复

	// 设置 debug mode
	if config.Config.DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

}
