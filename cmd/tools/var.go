package main

import (
	"strings"
	"yoomall/cmd/server"

	"github.com/lazyfury/pulse/framework/config"
	"github.com/lazyfury/pulse/framework/driver"
)

var appConfig = config.NewConfig()

func getDB() *driver.DB {
	return server.NewDB(appConfig)
}

func trimShellInputString(input string) string {
	return strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r")
}
