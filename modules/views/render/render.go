package render

import (
	"net/http"
	"yoomall/yoo/config"
	"yoomall/yoo/helper/response"
)

type Render struct {
	*response.HtmlTemplateResponse
}

func Html(template string, data interface{}) *Render {
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

	return a
}

func (r *Render) SEO(title string, description string, keywords string) *Render {
	r.WithExtra(map[string]any{
		"page": map[string]any{
			"title":       title,
			"description": description,
			"keywords":    keywords,
		},
	})

	return r
}
