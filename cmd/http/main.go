package main

import (
	"os"
	"yoomall/cmd/http/api"
	"yoomall/config"
	"yoomall/core/driver"
)

func main() {
	conf := config.NewConfig()
	server := api.NewApp(conf, driver.NewPostgresDB(conf.GetString("postgres.dsn")))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8900"
	}

	server.Start(port)
}
