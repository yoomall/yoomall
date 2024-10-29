package main

import "yoomall/cmd/api"

func main() {
	server := api.NewApp()
	server.Start()
}
