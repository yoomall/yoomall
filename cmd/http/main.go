package main

import (
	"api/cmd/server"
	"api/yoo/global"
	"html/template"
	"os"
	"yoomall/yoo/config"
	"yoomall/yoo/constants"

	"github.com/gin-gonic/gin"

	_template "yoomall/yoo/template"
)

func main() {
	conf := config.NewConfig()
	server := server.NewApp(conf, server.NewDB(conf), func(e *gin.Engine) *gin.Engine {
		// 设置模板
		temp := template.New("main").Funcs(_template.Funcs(nil))
		html := template.Must(_template.ParseGlob(temp, "templates", "*.html"))
		e.SetHTMLTemplate(html)
		return e
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = global.Config.GetString(constants.PORT)
	}

	if port == "0" || port == "" {
		port = "8900"
	}

	server.Start(port)
}
