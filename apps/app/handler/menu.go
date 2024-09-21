package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authmiddleware "lazyfury.github.com/yoomall-server/apps/auth/middleware"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/response"
)

type MenuHandler struct {
	DB *driver.DB
}

var _ core.Handler = (*MenuHandler)(nil)

func NewMenuHandler(db *driver.DB) *MenuHandler {
	return &MenuHandler{
		DB: db,
	}
}

func (m *MenuHandler) Register(router *core.RouterGroup) {
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Path:   "",
	}, func(ctx *gin.Context) {
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
	}, authmiddleware.AuthMiddleware(m.DB, true, false))
}

func (m *MenuHandler) GetRouterGroupName() string {
	return "menus"
}
