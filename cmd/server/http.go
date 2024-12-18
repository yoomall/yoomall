package server

import (
	"yoomall/modules/app"
	"yoomall/modules/auth"
	"yoomall/modules/common"
	commonservice "yoomall/modules/common/service"
	"yoomall/modules/post"
	"yoomall/modules/views"
	"yoomall/yoo"
	core "yoomall/yoo"
	"yoomall/yoo/constants"
	"yoomall/yoo/driver"
	coremiddleware "yoomall/yoo/middleware"

	"github.com/charmbracelet/log"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewHttpServer(
	config *viper.Viper,

	// apps
	app *app.DefaultApp,
	appV2 *app.DefaultV2App,
	auth *auth.AuthApp,
	postApp *post.PostApp,
	commonApp *common.CommonApp,
	viewsApp *views.ViewsApp,

	// services
	noufoundRecordService *commonservice.NotFoundRecordService,

	//other
	doc *core.Doc,
	setupEngine func(*gin.Engine) *gin.Engine,
) *yoo.HttpServer {
	// logger setup
	setupLogger(config)
	// logger setup

	log.Info("Start http server.", "debug mode: ", config.GetBool(constants.DEBUG))
	// 设置 debug mode
	if config.GetBool(constants.DEBUG) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := yoo.NewHttpServer(config, setupEngine(gin.Default()))

	server.Engine.Use(gin.Logger())
	server.Engine.SetTrustedProxies(nil)               //设置允许请求的域名
	server.Engine.Use(coremiddleware.CORSMiddleware()) // 跨域
	server.Engine.Use(gin.Recovery())                  // 错误恢复

	server.Engine.Use(static.Serve("/", static.LocalFile("public", false)))

	server.Engine.NoRoute(viewsApp.NotFoundHandler)

	root := yoo.Group(server.Engine, "")

	// v1
	v1 := yoo.Group(server.Engine, "/api/v1")
	v2 := yoo.Group(server.Engine, "/api/v2")

	v1.GET("/docs/api.json", doc.Handler)
	yoo.RegisterApps([]*yoo.RegisterApp{
		{App: viewsApp, Router: root.Group("")},

		{App: app, Router: v1.Group("")},
		{App: auth, Router: v1.Group("/auth")},
		{App: postApp, Router: v1.Group("/post")},
		{App: commonApp, Router: v1.Group("/common")},
		{App: appV2, Router: v2.Group("")},
	})

	return server
}

func setupLogger(config *viper.Viper) {
	logLevel := log.InfoLevel
	if config.GetBool(constants.DEBUG) {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)
}

func NewDoc() *core.Doc {
	return core.NewDoc()
}

func NewDB(config *viper.Viper) *driver.DB {
	return driver.NewDB(config.GetString(constants.MYSQL_DSN))
}
