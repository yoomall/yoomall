package api

import (
	"api/cmd/server"
	"embed"
	"html/template"
	"net/http"
	"yoomall/yoo"
	"yoomall/yoo/config"
	"yoomall/yoo/driver"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_template "yoomall/yoo/template"
)

var (
	//go:embed config.yaml
	configBytes []byte

	//go:embed templates/**
	templateFs embed.FS

	//go:embed manifest.json
	viteManifestJSON []byte
)

var conf *viper.Viper

var _server *yoo.HttpServer

func init() {
	conf = config.NewConfigFromBytes(configBytes)

	_server = server.NewApp(conf, driver.NewPostgresDB(conf.GetString("postgres.dsn")), func(e *gin.Engine) *gin.Engine {
		// 设置模板
		temp := template.New("main").Funcs(_template.Funcs(viteManifestJSON))
		html := template.Must(_template.ParseGlobEmbedFS(temp, templateFs, "templates", "*.html"))
		e.SetHTMLTemplate(html)
		return e
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	_server.Engine.ServeHTTP(w, r)
}
