package httpserver

import (
	"strconv"

	"yoomall/config"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	Engine *gin.Engine
	Config *viper.Viper
}

func (h *HttpServer) Start() *gin.Engine {
	port := strconv.Itoa(config.Config.Port)
	if port == "0" {
		port = "8900"
	}
	log.Info("【服务正在运行】 http://127.0.0.1:" + port)
	h.Engine.Run(":" + port)
	return h.Engine
}
