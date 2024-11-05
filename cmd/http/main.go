package main

import (
	"html/template"
	"os"
	"yoomall/cmd/server"
	"yoomall/yoo/config"
	"yoomall/yoo/constants"

	"github.com/gin-gonic/gin"

	_template "yoomall/yoo/template"
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
