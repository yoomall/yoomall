//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"lazyfury.github.com/yoomall-server/apps/auth"
	"lazyfury.github.com/yoomall-server/config"
	httpserver "lazyfury.github.com/yoomall-server/core/http"
)

func NewApp() httpserver.HttpServer {
	wire.Build(NewHttpServer, NewDB, config.NewConfig, auth.WireSet)
	return httpserver.HttpServer{}
}
