//go:build wireinject
// +build wireinject

package api

import (
	"yoomall/core/driver"
	httpserver "yoomall/core/http"
	"yoomall/modules/app"
	"yoomall/modules/auth"
	"yoomall/modules/common"
	"yoomall/modules/post"
	"yoomall/modules/views"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewApp(conf *viper.Viper, db *driver.DB, setupEngine func(*gin.Engine) *gin.Engine) httpserver.HttpServer {
	wire.Build(NewHttpServer, NewDoc, app.WireSet, post.WireSet, auth.WireSet, common.WireSet, views.WireSet)
	return httpserver.HttpServer{}
}
