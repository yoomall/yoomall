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
	DB        *driver.DB
	authMidds *authmiddleware.AuthMiddlewareGroup
}

var _ core.Handler = (*MenuHandler)(nil)

func NewMenuHandler(db *driver.DB, authMidds *authmiddleware.AuthMiddlewareGroup) *MenuHandler {
	return &MenuHandler{
		DB:        db,
		authMidds: authMidds,
	}
}

func (m *MenuHandler) Register(router *core.RouterGroup) {
	router.Use(m.authMidds.MustAuthMiddleware)
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
	return ui.NewRouter("overview", "控制台/首页", "ant-design:home-outlined", "/overview", nil, &ui.Page{
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
	}).AddChildren(
		ui.NewRouter("user-list", "用户列表", "ant-design:user-outlined", "user-list", nil, &ui.Page{
			Title:     "Users",
			Component: "HomeView",
			Widgets:   []ui.IWidget{},
		}),
	).AddChildren(
		ui.NewRouter("role-list", "角色列表", "ant-design:user-outlined", "role-list", nil, &ui.Page{
			Title:     "Roles",
			Component: "TableView",
			Widgets:   []ui.IWidget{},
		}),
	)
}
