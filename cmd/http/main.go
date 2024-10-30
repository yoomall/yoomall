package main

import (
	"embed"
	"html/template"
	"os"
	"yoomall/cmd/http/api"
	"yoomall/config"
	"yoomall/core/driver"

	_template "yoomall/core/template"
)

var (
	//go:embed templates/**
	templateFs embed.FS
)

func main() {
	conf := config.NewConfig()
	server := api.NewApp(conf, driver.NewPostgresDB(conf.GetString("postgres.dsn")))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8900"
	}

	// 设置模板
	temp := template.New("main").Funcs(_template.Funcs)
	html := template.Must(_template.ParseGlobEmbedFS(temp, templateFs, "templates", "*.html"))
	server.Engine.SetHTMLTemplate(html)

	server.Start(port)
}
