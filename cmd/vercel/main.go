package main

import (
	"fmt"
	"net/http"
	"yoomall/api"
)

func main() {
	// 测试 vercel 服务配置是否正常
	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.Handler(w, r)
	})
	fmt.Println("start http server on port: http://127.1:8900")
	http.ListenAndServe(":8900", server)
}
