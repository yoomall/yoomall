package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lazyfury.github.com/yoomall-server/app"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	httpserver "lazyfury.github.com/yoomall-server/core/http"
	"lazyfury.github.com/yoomall-server/docs"
	"lazyfury.github.com/yoomall-server/modules/post"
)

func NewHttpServer(
	config *viper.Viper,

	app *app.DefaultApp,
	postApp *post.DefaultApp,
) httpserver.HttpServer {
	engine := gin.Default()

	setup(engine)

	v1 := engine.Group("/api/v1")

	register(
		&core.RegisterApp{App: app, Router: v1.Group("")},
		&core.RegisterApp{App: postApp, Router: v1.Group("/posts")},
	)

	return httpserver.HttpServer{
		Engine: engine,
		Config: config,
	}
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

	setupSwag(engine)
}

func setupSwag(engine *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true)))
}

func register(apps ...*core.RegisterApp) {
	for _, instance := range apps {
		instance.Register()
	}
}
