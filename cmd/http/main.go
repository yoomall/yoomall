package main

import (
	"html/template"
	"os"
	"yoomall/cmd/http/api"
	"yoomall/config"
	"yoomall/core/driver"

	"github.com/gin-gonic/gin"

	_template "yoomall/core/template"
)

func main() {
	conf := config.NewConfig()
	server := api.NewApp(conf, driver.NewPostgresDB(conf.GetString("postgres.dsn")), func(e *gin.Engine) *gin.Engine {
		// 设置模板
		temp := template.New("main").Funcs(_template.Funcs)
		html := template.Must(_template.ParseGlob(temp, "templates", "*.html"))
		e.SetHTMLTemplate(html)
		return e
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8900"
	}

	server.Start(port)
}
