package api

import (
	"net/http"
	"yoomall/cmd/api"
)

func Listen(w http.ResponseWriter, r *http.Request) {
	server := api.NewApp()
	server.Engine.ServeHTTP(w, r)
}
