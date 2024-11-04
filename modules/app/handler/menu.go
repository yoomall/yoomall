package handler

import (
	"net/http"

	"yoomall/core"
	"yoomall/core/driver"
	"yoomall/core/helper/response"
	"yoomall/core/ui"
	authmiddleware "yoomall/modules/auth/middleware"

	"github.com/gin-gonic/gin"
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
			systemConfigUI(),
		}).Done(ctx)
	})
}

func (m *MenuHandler) GetRouterGroupName() string {
	return "menus"
}

func overviewUI() ui.IRouter {
	return ui.NewRouter("overview", "控制台", "ant-design:home-outlined", "/overview", nil, &ui.Page{
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
		Component: "RouterView",
		Widgets:   []ui.IWidget{},
	}).AddChildren(
		ui.NewRouter("user-list", "用户列表", "ant-design:user-outlined", "user-list", nil, &ui.Page{
			Title:     "Users",
			Component: "TableView",
			Widgets:   []ui.IWidget{},
			Table: ui.NewTable().WithColumns([]ui.TableColumn{
				{
					Prop:  "username",
					Label: "用户名",
					Width: "160px",
					Props: map[string]any{
						"class": "line-clamp-1",
					},
				},
				{
					Prop:  "email",
					Label: "邮箱",
					Width: "160px",
					Props: map[string]any{
						"class": "line-clamp-1 font-bold",
					},
				},
				{
					Prop:  "phone",
					Label: "手机",
					Width: "125px",
					Props: map[string]any{
						"class": "line-clamp-1 font-bold !text-primary",
					},
				},
				{
					Prop:  "created_at",
					Label: "创建时间",
					Width: "240px",
					Props: nil,
				},
				{
					Prop:  "updated_at",
					Label: "更新时间",
					Width: "240px",
					Props: nil,
				},
				{
					Prop:  "last_login_at",
					Label: "最后登录时间",
					Width: "240px",
					Props: nil,
				},
			}).WithActions([]*ui.Action{
				ui.NewEditAction().WithApiKey("edit"),
			}),
		}).WithDescription("用户列表, 系统用户列表管理").WithApis(
			map[string]string{
				"list": "/auth/users/user-list",
			},
		),
	).AddChildren(
		ui.NewRouter("role-list", "角色列表", "ant-design:user-outlined", "role-list", nil, &ui.Page{
			Title:     "Roles",
			Component: "TableView",
			Widgets:   []ui.IWidget{},
			Table: ui.NewTable().WithColumns([]ui.TableColumn{
				{
					Prop:  "role_name",
					Label: "角色",
					Width: "160px",
					Props: map[string]any{
						"class": "line-clamp-1",
					},
				},
				{
					Prop:     "role_code",
					Label:    "code",
					Sortable: true,
					Width:    "160px",
				},
				{
					Prop:     "created_at",
					Label:    "创建时间",
					Width:    "240px",
					Props:    nil,
					Sortable: true,
				},
				{
					Prop:  "updated_at",
					Label: "更新时间",
					Width: "240px",
					Props: nil,
				},
			},
			).WithForms(map[string]*ui.Form{
				"create": ui.NewForm("create", "添加/编辑角色", "/auth/user-roles/create-role").WithRows([][]*ui.FormItem{
					{
						ui.NewFormItem("角色名称", "role_name", "text", "请输入角色名称").WithRequired(true),
						ui.NewFormItem("Code", "role_code", "text", "请输入角色 Code 定义"),
					},
				}),
				"update": ui.NewForm("create", "添加/编辑角色", "/auth/user-roles/update-role").WithRows([][]*ui.FormItem{
					{
						ui.NewFormItem("角色名称", "role_name", "text", "请输入角色名称").WithRequired(true),
						ui.NewFormItem("Code", "role_code", "text", "请输入角色 Code 定义"),
					},
				}),
			}).WithSearch(ui.NewForm("search", "搜索", "/auth/user-roles/search").WithRows([][]*ui.FormItem{
				{
					ui.NewFormItem("角色名称", "role_name__like", "text", "请输入角色名称"),
				},
			})).WithActions([]*ui.Action{
				ui.NewEditAction().WithFormKey("update"),
				ui.NewDeleteAction(),
			}),
		}).WithDescription("角色列表, 用于给用户分配角色").WithApis(
			map[string]string{
				"list":   "/auth/user-roles/role-list",
				"delete": "/auth/user-roles/delete-role",
			},
		),
	).AddChildren(
		ui.NewRouter("user-logs", "登录日志", "ant-design:login-outlined", "user-logs", nil, &ui.Page{
			Title:     "Login Logs",
			Component: "TableView",
			Widgets:   []ui.IWidget{},
			Table: ui.NewTable().WithColumns([]ui.TableColumn{
				{
					Prop:  "user.username",
					Label: "用户名",
					Width: "180px",
					Props: map[string]any{
						"class": "line-clamp-1",
					},
				},
				{
					Prop:  "ip",
					Label: "IP",
					Width: "160px",
					Props: map[string]any{
						"class": "line-clamp-1",
					},
				},
				{
					Prop:  "agent",
					Label: "User Agent",
					Width: "160px",
					Props: map[string]any{
						"class": "line-clamp-1",
					},
				},
				// browser
				{
					Prop:  "browser",
					Label: "浏览器",
					Width: "160px",
					Props: map[string]any{
						"class": "line-clamp-1",
					},
				},
				// os
				{
					Prop:  "os",
					Label: "OS",
					Width: "160px",
					Props: map[string]any{
						"class": "line-clamp-1",
					},
				},
				// device
				{
					Prop:  "device",
					Label: "设备",
					Width: "160px",
					Props: map[string]any{
						"class": "line-clamp-1",
					},
				},
				// expire_time
				{
					Prop:  "expire_time",
					Label: "过期时间",
					Width: "240px",
				},
				{
					Prop:  "created_at",
					Label: "创建时间",
					Width: "240px",
				},
				{
					Prop:  "updated_at",
					Label: "更新时间",
					Width: "240px",
				},
			})}).WithDescription("登录日志列表").WithApis(
			map[string]string{
				"list": "/auth/user-logs/logs",
			},
		),
	).AddChildren(
		ui.NewRouter("permissions-list", "权限管理", "ep:unlock", "permissions-list", nil, &ui.Page{
			Title:     "permissions",
			Component: "TableView",
			Table: ui.NewTable().WithColumns([]ui.TableColumn{
				{
					Prop:  "id",
					Label: "ID",
					Width: "160px",
				},
				{
					Prop:  "name",
					Label: "名称",
					Width: "160px",
				},
			}),
		}).WithDescription("权限列表").WithApis(
			map[string]string{
				"list": "/auth/permissions/list",
			},
		),
	)
}

func systemConfigCreateForm() *ui.Form {
	return ui.NewForm("key", "Key", "/common/system-configs/create").WithRows(
		[][]*ui.FormItem{
			{
				{
					Label:       "Key",
					Type:        "text",
					Prop:        "key",
					Placeholder: "Key",
				},
				{
					Label:       "Value",
					Type:        "text",
					Prop:        "value",
					Placeholder: "Value",
				},
			},
		},
	)
}

func systemConfigUI() ui.IRouter {
	return ui.NewRouter("system", "系统设置", "ant-design:setting-outlined", "/system", nil, &ui.Page{
		Title:     "System",
		Component: "RouterView",
		Widgets:   []ui.IWidget{},
	}).AddChildren(
		ui.NewRouter("system-configs", "系统配置", "ant-design:setting-outlined", "system-configs", nil, &ui.Page{
			Title:     "System Configs",
			Component: "TableView",
			Table: ui.NewTable().WithColumns([]ui.TableColumn{
				{
					Prop:  "key",
					Label: "Key",
					Width: "160px",
				},
				{
					Prop:  "value",
					Label: "Value",
					Width: "160px",
				},
			}).WithForms(map[string]*ui.Form{
				"create": systemConfigCreateForm().WithApi("/common/system-configs/create"),
				"update": systemConfigCreateForm().WithApi("/common/system-configs/update"),
			}).WithActions([]*ui.Action{
				ui.NewEditAction().WithFormKey("update"),
				ui.NewDeleteAction().WithApiKey("delete"),
			}),
		}).WithDescription("系统配置列表").WithApis(
			map[string]string{
				"list":   "/common/system-configs/list",
				"delete": "/common/system-configs/delete",
				"export": "/common/system-configs/export",
			},
		),
	).AddChildren(
		ui.NewRouter("system-config-groups", "系统配置组", "ant-design:setting-outlined", "system-config-groups", nil, &ui.Page{
			Title:     "System Config Groups",
			Component: "TableView",
			Table: ui.NewTable().WithColumns([]ui.TableColumn{
				{
					Prop:  "name",
					Label: "Name",
					Width: "160px",
				},
				{
					Prop:  "key",
					Label: "Key",
					Width: "160px",
				},
				{
					Prop:  "description",
					Label: "Description",
					Width: "160px",
				},
			}).WithForms(map[string]*ui.Form{
				"create": ui.NewForm("name", "Name", "/common/system-configs/groups/create").WithRows(
					[][]*ui.FormItem{
						{
							{
								Label:       "Name",
								Type:        "text",
								Prop:        "name",
								Placeholder: "Name",
							},
							{
								Label:       "Key",
								Type:        "text",
								Prop:        "key",
								Placeholder: "Key",
							},
						},
						{
							{
								Label:       "Description",
								Type:        "textarea",
								Prop:        "description",
								Placeholder: "Description",
								Width:       "100%",
							},
						},
					},
				),
			}),
		}).WithDescription("系统配置组列表").WithApis(
			map[string]string{
				"list": "/common/system-configs/groups/list",
			},
		),
	)
}
