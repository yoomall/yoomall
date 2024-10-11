//go:build wireinject
// +build wireinject

package main

import (
	"yoomall/apps/auth"
	"yoomall/config"
	httpserver "yoomall/core/http"

	"github.com/google/wire"
)

func NewApp() httpserver.HttpServer {
	wire.Build(NewHttpServer, NewDB, config.NewConfig, auth.WireSet)
	return httpserver.HttpServer{}
}
