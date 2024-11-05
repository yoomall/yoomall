package views

import (
	"net/http"
	"yoomall/modules/views/render"

	"github.com/gin-gonic/gin"
)

func (v *ViewsApp) NotFoundHandler(ctx *gin.Context) {
	if ctx.Request.Header.Get("Accept") == "application/json" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "不存在的路由"})
		return
	}
	render.Html("404.html", nil).Done(ctx)
}
