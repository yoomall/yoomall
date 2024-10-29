package api

import (
	"net/http"
	app "yoomall/cmd/api"
)

func Listen(w http.ResponseWriter, r *http.Request) {
	server := app.NewApp()
	server.Engine.ServeHTTP(w, r)
}
