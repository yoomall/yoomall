//go:build wireinject
// +build wireinject

package api

import (
	"yoomall/src/apps/app"
	"yoomall/src/apps/auth"
	"yoomall/src/apps/common"
	"yoomall/src/apps/post"
	"yoomall/src/apps/views"
	"yoomall/src/core/driver"
	httpserver "yoomall/src/core/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewApp(conf *viper.Viper, db *driver.DB, setupEngine func(*gin.Engine) *gin.Engine) httpserver.HttpServer {
	wire.Build(NewHttpServer, NewDoc, app.WireSet, post.WireSet, auth.WireSet, common.WireSet, views.WireSet)
	return httpserver.HttpServer{}
}
