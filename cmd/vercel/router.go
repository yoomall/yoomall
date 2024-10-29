package vercel

import "yoomall/cmd/api"

func Listen() {
	server := api.NewApp()
	server.Engine.ServeHTTP(nil, nil)
}
