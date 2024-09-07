package httpserver

import (
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/core/constants"
)

type HttpServer struct {
	Engine *gin.Engine
}

func (h *HttpServer) Start() *gin.Engine {
	port := strconv.Itoa(constants.CONFIG.HTTP.Port)
	if port == "0" {
		port = "8900"
	}
	log.Info("[server on] http://127.0.0.1:" + port)
	h.Engine.Run(":" + port)
	return h.Engine
}
