package views

import (
	"github.com/yoomall/yoomall/apps/views/render"

	"github.com/gin-gonic/gin"
)

func aboutView(ctx *gin.Context) {
	render.Html("about.html", nil).Done(ctx)
}
