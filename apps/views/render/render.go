package render

import (
	"net/http"
	"yoomall/core/helper/response"

	"github.com/gin-gonic/gin"
)

type Render struct {
	*response.HtmlTemplateResponse
}

func Html(ctx *gin.Context, template string, data interface{}) {
	a := &Render{
		HtmlTemplateResponse: &response.HtmlTemplateResponse{
			HttpCode: http.StatusOK,
			Template: template,
			Data:     data,
		},
	}
	a.Done(ctx)
}
