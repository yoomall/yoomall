package main

import (
	"html/template"
	"os"
	"yoomall/cmd/server"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/pulse/framework/config"
	"github.com/lazyfury/pulse/framework/constants"

	_template "github.com/lazyfury/pulse/plugins/template"
)

func main() {
	conf := config.NewConfig()
	db := server.NewDB(conf)
	server := server.NewApp(conf, db, func(e *gin.Engine) *gin.Engine {
		// 设置模板
		temp := template.New("main").Funcs(_template.Funcs(nil)).Funcs(template.FuncMap{
			"hello": func() string {
				return "hello world by template funcs!"
			},
		})
		html := template.Must(_template.ParseGlob(temp, "templates", "*.html"))
		e.SetHTMLTemplate(html)
		return e
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = conf.GetString(constants.PORT)
	}

	if port == "0" || port == "" {
		port = "8900"
	}

	server.Start(port)
}
