package httpserver

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	Engine *gin.Engine
	Config *viper.Viper
}

func (h *HttpServer) Start() *gin.Engine {
	port := h.Config.GetString("http.port")
	if port == "0" {
		port = "8900"
	}
	log.Info("[server on] http://127.0.0.1:" + port)
	h.Engine.Run(":" + port)
	return h.Engine
}
