package handler

import (
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/driver"
)

type Handler interface {
	Register(router *gin.RouterGroup)
}

type handler struct {
	App core.App
	DB  *driver.DB
}
