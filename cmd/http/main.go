package main

import (
	"html/template"
	"os"
	"yoomall/cmd/http/api"
	"yoomall/config"
	"yoomall/core/driver"

	_template "yoomall/core/template"
)

func main() {
	conf := config.NewConfig()
	config.Init(false, conf)
	server := api.NewApp(conf, driver.NewPostgresDB(conf.GetString("postgres.dsn")))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8900"
	}

	// 设置模板
	temp := template.New("main").Funcs(_template.Funcs)
	html := template.Must(_template.ParseGlob(temp, "templates", "*.html"))
	server.Engine.SetHTMLTemplate(html)

	server.Start(port)
}
