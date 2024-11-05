package views

import (
	"yoomall/modules/views/render"

	"github.com/gin-gonic/gin"
)

func aboutView(ctx *gin.Context) {
	render.Html("about.html", nil).Done(ctx)
}
