package api

import (
	"yoomall/apps/app"
	"yoomall/apps/auth"
	"yoomall/apps/common"
	commonservice "yoomall/apps/common/service"
	"yoomall/apps/post"
	"yoomall/apps/views"
	"yoomall/core"
	"yoomall/core/constants"
	"yoomall/core/driver"
	httpserver "yoomall/core/http"
	coremiddleware "yoomall/core/middleware"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewHttpServer(
	config *viper.Viper,

	app *app.DefaultApp,
	auth *auth.AuthApp,
	postApp *post.PostApp,
	commonApp *common.CommonApp,
	viewsApp *views.ViewsApp,

	noufoundRecordService *commonservice.NotFoundRecordService,

	doc *core.Doc,
	setupEngine func(*gin.Engine) *gin.Engine,
) httpserver.HttpServer {
	engine := gin.Default()

	engine.SetTrustedProxies(nil)               //设置允许请求的域名
	engine.Use(coremiddleware.CORSMiddleware()) // 跨域
	engine.Use(gin.Recovery())                  // 错误恢复

	// 设置 debug mode
	if config.GetBool(constants.DEBUG) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = setupEngine(engine)

	engine.Use(static.Serve("/", static.LocalFile("public", false)))

	engine.NoRoute(viewsApp.NotFoundHandler)

	root := &core.RouterGroup{
		RouterGroup: engine.Group("/"),
	}

	v1 := &core.RouterGroup{
		RouterGroup: engine.Group("/api/v1"),
	}
	v1.GET("/docs/api.json", doc.Handler)

	var apps = []*core.RegisterApp{
		{App: viewsApp, Router: root.Group("")},

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

}
