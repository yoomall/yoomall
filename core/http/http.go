package httpserver

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	Engine *gin.Engine
	Config *viper.Viper
}

func (h *HttpServer) Start(port string) *gin.Engine {
	if port == "0" {
		port = "8900"
	}
	log.Info(fmt.Sprintf("start http server on port: http://127.1:%s", port))
	h.Engine.Run(fmt.Sprintf(":%s", port))
	return h.Engine
}
