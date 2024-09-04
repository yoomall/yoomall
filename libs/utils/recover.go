package utils

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func RecoverHandlerFunc(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if err := recover(); err != nil {
			msg := ""
			if e, ok := err.(error); ok {
				msg = e.Error()
			}
			log.Error(err)
			ctx.JSON(500, map[string]any{"error": msg})
			ctx.Abort()
			panic(err)
		}
	}(ctx)

	ctx.Next()
}
