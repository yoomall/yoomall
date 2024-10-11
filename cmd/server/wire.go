//go:build wireinject
// +build wireinject

package main

import (
	"yoomall/apps/app"
	"yoomall/apps/auth"
	"yoomall/apps/common"
	"yoomall/apps/post"
	"yoomall/config"
	httpserver "yoomall/core/http"

	"github.com/google/wire"
)

func NewApp() httpserver.HttpServer {
	wire.Build(NewHttpServer, NewDB, NewDoc, config.NewConfig, app.WireSet, post.WireSet, auth.WireSet, common.WireSet)
	return httpserver.HttpServer{}
}
