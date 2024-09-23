package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authmiddleware "lazyfury.github.com/yoomall-server/apps/auth/middleware"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/response"
	"lazyfury.github.com/yoomall-server/core/ui"
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
	router.Use(authmiddleware.AuthMiddleware(m.DB, true, false))
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Path:   "",
	}).GET("", func(ctx *gin.Context) {
		response.Success([]any{
			overviewUI(),
			userManagementUI(),
		}).Done(ctx)
	})
}

func (m *MenuHandler) GetRouterGroupName() string {
	return "menus"
}

func overviewUI() ui.IRouter {
	return ui.NewRouter("overview", "Overview", "ant-design:home-outlined", "/overview", nil, &ui.Page{
		Title:     "Overview",
		Component: "HomeView",
		Widgets: []ui.IWidget{
			ui.NewTable().WithColumns([]ui.TableColumn{
				{
					Prop:  "id",
					Label: "ID",
					Width: "100px",
					Props: nil,
				},
			}),
		},
	})
}

func userManagementUI() ui.IRouter {
	return ui.NewRouter("user-management", "用户管理", "ant-design:user-outlined", "/user-management", nil, &ui.Page{
		Title:     "User Management",
		Component: "UserManagementView",
		Widgets:   []ui.IWidget{},
	})
}
