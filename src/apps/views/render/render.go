package render

import (
	"net/http"
	"yoomall/src/config"
	"yoomall/src/core/helper/response"

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

	a.WithExtra(map[string]any{
		"footer": map[string]any{
			"links": []map[string]string{
				{
					"href": "https://github.com/yoomall/yoomall",
					"text": "yoomall",
				},
				{
					"href": "https://github.com/yoomall/yoomall-ui",
					"text": "yoomall-ui",
				},
			},
		},
		"site": map[string]any{
			"title":       config.Config.GetString("site.title"),
			"description": config.Config.GetString("site.description"),
			"keywords":    config.Config.GetString("site.keywords"),
			"author":      config.Config.GetString("site.author"),
			"logo":        config.Config.GetString("site.logo"),
		},
	})

	a.Done(ctx)
}
