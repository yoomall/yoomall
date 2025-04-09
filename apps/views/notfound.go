package views

import (
	"net/http"
	"yoomall/apps/views/render"

	"github.com/gin-gonic/gin"
)

func SortUrlHandler(ctx *gin.Context) (breaking bool) {
	// 处理排序的逻辑
	path := ctx.Request.URL.Path
	if path == "/some-content" {
		render.Html("some-content.html", nil).Done(ctx)
		return true
	}

	//todo: articles

	// todo:category

	// todo:tags

	// todo:products

	return false
}

func (v *ViewsApp) NotFoundHandler(ctx *gin.Context) {

	if breaking := SortUrlHandler(ctx); breaking {
		return
	}

	if ctx.Request.Header.Get("Accept") == "application/json" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "不存在的路由"})
		return
	}
	render.Html("404.html", nil).Done(ctx)
}
