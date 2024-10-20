package main

import (
	"html/template"
	"net/http"

	"yoomall/apps/app"
	"yoomall/apps/auth"
	"yoomall/apps/common"
	commonservice "yoomall/apps/common/service"
	"yoomall/apps/post"
	"yoomall/config"
	"yoomall/core"
	"yoomall/core/constants"
	"yoomall/core/driver"
	"yoomall/core/helper/response"
	httpserver "yoomall/core/http"
	coremiddleware "yoomall/core/middleware"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_template "yoomall/core/template"
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
		response.Html(http.StatusOK, "ok", nil, "index.html", http.StatusOK).Done(ctx)
	})

	engine.NoRoute(func(ctx *gin.Context) {
		noufoundRecordService.Add(ctx.Request.URL.Path, ctx.Request)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "不存在的路由"})
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

	// 设置模板
	temp := template.New("main").Funcs(_template.Funcs)
	html := template.Must(_template.ParseGlob(temp, "templates", "*.html"))
	engine.SetHTMLTemplate(html)

	engine.Use(static.Serve("/", static.LocalFile("public", false)))

	// 设置 debug mode
	if config.Config.DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

}
