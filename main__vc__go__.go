package main

import (
	"net/http"
	"os"
	"syscall"

	"api/api"
	vc "github.com/vercel/go-bridge/go/bridge"
)

func checkForLambdaWrapper() {
	wrapper := os.Getenv("AWS_LAMBDA_EXEC_WRAPPER")
	if wrapper == "" {
		return
	}

	// Removing the env var doesn't work
	// Set it to empty string to override the previous value
	os.Setenv("AWS_LAMBDA_EXEC_WRAPPER", "")
	argv := append([]string{wrapper}, os.Args...)
	err := syscall.Exec(wrapper, argv, os.Environ())
	if err != nil {
		panic(err)
	}
}

func main() {
	checkForLambdaWrapper()
	vc.Start(http.HandlerFunc(api.Handler_api_router_go))
}
