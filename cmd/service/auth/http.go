package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/auth"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	httpserver "lazyfury.github.com/yoomall-server/core/http"
	coremiddleware "lazyfury.github.com/yoomall-server/core/middleware"
)

func NewHttpServer(
	config *viper.Viper,

	auth *auth.AuthApp,
) httpserver.HttpServer {
	engine := gin.Default()

	setup(engine)

	engine.Use(coremiddleware.CORSMiddleware())

	v1 := &core.RouterGroup{
		RouterGroup: engine.Group("/api/v1"),
	}

	var apps = []core.RegisterApp{
		{App: auth, Router: v1.Group("/auth")},
	}

	for _, app := range apps {
		app.Register()
	}

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
}
