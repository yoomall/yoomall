package main

import (
	"os"
	"yoomall/cmd/http/api"
)

func main() {
	server := api.NewApp()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8900"
	}

	server.Start(port)
}
