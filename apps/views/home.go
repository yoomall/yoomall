package views

import (
	"github.com/yoomall/yoomall/apps/views/render"

	"github.com/gin-gonic/gin"
)

func homeView(ctx *gin.Context) {
	render.Html("index.html", map[string]any{
		"packages": []any{
			map[string]any{
				"price": 0,
			},
			map[string]any{
				"price": 680,
			},
			map[string]any{
				"price": 1980,
			},
		},
		"products": []any{
			map[string]any{
				"name": "yoomall",
			},
			map[string]any{
				"name": "yoomall-ui",
			},
			map[string]any{
				"name": "yoomall-admin",
			},
			map[string]any{
				"name": "yoomall-admin-ui",
			},
			map[string]any{
				"name": "yoomall-admin-api",
			},
			map[string]any{
				"name": "yoomall-admin-console",
			},
			map[string]any{
				"name": "yoomall-admin-ui",
			},
		},
	}).SEO("easyly make your design to production.", "yoomall", "yoomall").Done(ctx)
}
