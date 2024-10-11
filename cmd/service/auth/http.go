package main

import (
	"yoomall/apps/auth"
	"yoomall/config"
	"yoomall/core"
	"yoomall/core/constants"
	"yoomall/core/driver"
	httpserver "yoomall/core/http"
	coremiddleware "yoomall/core/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	return driver.NewDB(config.GetString(constants.MYSQL_DSN))
}

func setup(engine *gin.Engine) {
	engine.SetTrustedProxies(nil)

	if config.Config.DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
