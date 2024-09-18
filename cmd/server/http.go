package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/app"
	"lazyfury.github.com/yoomall-server/apps/app/handler"
	"lazyfury.github.com/yoomall-server/apps/auth"
	"lazyfury.github.com/yoomall-server/apps/post"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	httpserver "lazyfury.github.com/yoomall-server/core/http"
	coremiddleware "lazyfury.github.com/yoomall-server/core/middleware"
)

func NewHttpServer(
	config *viper.Viper,

	app *app.DefaultApp,
	auth *auth.AuthApp,
	postApp *post.PostApp,

	doc *core.Doc,
	dtkHandler *handler.DtkHandler,
) httpserver.HttpServer {
	engine := gin.Default()

	setup(engine)

	engine.Use(coremiddleware.CORSMiddleware())
	engine.Use(gin.Recovery())

	engine.GET("", func(ctx *gin.Context) {
		ctx.String(200, ":) yoomall server is running.")
	})

	v1 := engine.Group("/api/v1")
	v1.GET("/docs/*any", doc.Handler)

	dtkHandler.Register(&core.RouterGroup{
		RouterGroup: v1,
	})

	var apps = []*core.RegisterApp{
		{App: app, Router: v1.Group("")},
		{App: auth, Router: v1.Group("/auth")},
		{App: postApp, Router: v1.Group("/post")},
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
	return driver.NewDB(config.GetString("mysql.dsn"))
}

func setup(engine *gin.Engine) {
	engine.SetTrustedProxies(nil)

	if config.Config.DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

}
