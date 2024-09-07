package main

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	defappConfig "lazyfury.github.com/yoomall-server/app/config"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	httpserver "lazyfury.github.com/yoomall-server/core/http"
	"lazyfury.github.com/yoomall-server/docs"
	"lazyfury.github.com/yoomall-server/modules/post"
)

func NewHttpServer(
	config *config.Config,

	app *defappConfig.DefaultApp,
	postApp *post.DefaultApp,
) httpserver.HttpServer {
	engine := gin.Default()

	setup(engine, config)

	v1 := engine.Group("/api/v1")

	register(
		&RegisterApp{app: app, router: v1.Group(""), config: config},
		&RegisterApp{app: postApp, router: v1.Group("/posts"), config: config},
	)

	return httpserver.HttpServer{
		Engine: engine,
		Config: config,
	}
}

func NewDB(config *config.Config) *driver.DB {
	return driver.NewDB(config.MysqlDsn())
}

func setup(engine *gin.Engine, config *config.Config) {
	engine.SetTrustedProxies(nil)

	if config.DEBUG {
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

type RegisterApp struct {
	router *gin.RouterGroup
	app    core.App
	config *config.Config
}

func (instance *RegisterApp) Register() {
	log.Info(instance.app.GetName() + "====================================")
	log.Info("注册app", "app", instance.app.GetName())
	if instance.config.DEBUG {
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
