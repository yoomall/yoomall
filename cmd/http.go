package main

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lazyfury.github.com/yoomall-server/app"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/constants"
	"lazyfury.github.com/yoomall-server/core/driver"
	httpserver "lazyfury.github.com/yoomall-server/core/http"
	"lazyfury.github.com/yoomall-server/docs"
	"lazyfury.github.com/yoomall-server/modules/post"
)

func NewHttpServer(
	app *app.DefaultApp,
	postApp *post.DefaultApp,
) httpserver.HttpServer {
	engine := gin.Default()

	setup(engine)

	v1 := engine.Group("/api/v1")

	register(
		&RegisterApp{app: app, router: v1.Group("")},
		&RegisterApp{app: postApp, router: v1.Group("/posts")},
	)

	return httpserver.HttpServer{
		Engine: engine,
	}
}

func NewDB(config *config.Config) *driver.DB {
	return driver.NewDB(config.MysqlDsn())
}

func setup(engine *gin.Engine) {
	log.Info("hello world")
	engine.SetTrustedProxies(nil)

	// config
	constants.CONFIG = config.GetConfig("./config.yaml")
	log.Info("config load success", "config", constants.CONFIG)

	if constants.CONFIG.DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// db
	constants.DB = driver.NewDB(constants.CONFIG.MysqlDsn())
	// db

	docs.SwaggerInfo.BasePath = "/api"
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true)))

}

type RegisterApp struct {
	router *gin.RouterGroup
	app    core.App
}

func (instance *RegisterApp) Register() {
	log.Info(instance.app.GetName() + "====================================")
	log.Info("注册app", "app", instance.app.GetName())
	if constants.CONFIG.DEBUG {
		log.Info("迁移中", "app", instance.app.GetName())
		instance.app.Migrate()
		log.Info("迁移成功 success", "app", instance.app.GetName())
	}
	instance.router.Use(instance.app.Middleware()...)
	instance.app.Register(instance.router)
	log.Info("注册成功", "app", instance.app.GetName())
}

func register(apps ...*RegisterApp) {
	for _, instance := range apps {
		instance.Register()
	}
}
