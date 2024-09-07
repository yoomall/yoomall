package handler

import (
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/helper/response"
)

type MenuHandler struct {
}

var _ core.Handler = (*MenuHandler)(nil)

func NewMenuHandler() *MenuHandler {
	return &MenuHandler{}
}

func (m *MenuHandler) Register(router *gin.RouterGroup) {
	router.GET("", func(ctx *gin.Context) {
		response.Success([]any{
			map[string]any{
				"key":       "overview",
				"name":      "overview",
				"path":      "/overview",
				"component": "HomeView",
				"icon":      "ant-design:home-outlined",
				"title":     "Overview",
				"meta": map[string]any{
					"noCache": true,
				},
			},
		}).Done(ctx)
	})
}

func (m *MenuHandler) GetRouterGroupName() string {
	return "menus"
}
