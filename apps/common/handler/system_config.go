package handler

import (
	"net/http"
	"strconv"

	authmiddleware "yoomall/apps/auth/middleware"
	"yoomall/apps/common/model"
	commonservice "yoomall/apps/common/service"
	"yoomall/core"
	"yoomall/core/driver"
	"yoomall/core/helper/curd"
	"yoomall/core/helper/response"

	"github.com/gin-gonic/gin"
)

type SystemConfigHandler struct {
	DB        *driver.DB
	service   *commonservice.SystemConfigService
	authMidds *authmiddleware.AuthMiddlewareGroup
	curd      *curd.CRUD
	groupCurd *curd.CRUD
}

func NewSystemConfigHandler(db *driver.DB, service *commonservice.SystemConfigService, authMiddlewareGroup *authmiddleware.AuthMiddlewareGroup) *SystemConfigHandler {
	return &SystemConfigHandler{
		DB:        db,
		service:   service,
		curd:      curd.New(db, &model.SystemConfig{}).WithExportAttrs(model.SystemConfigExeclConfig),
		authMidds: authMiddlewareGroup,
		groupCurd: curd.New(db, &model.SystemConfigGroup{}),
	}
}

// 检查 handler 是否实现
var _ core.Handler = (*SystemConfigHandler)(nil)

// GetRouterGroupName implements core.Handler.
func (s *SystemConfigHandler) GetRouterGroupName() string {
	return "system-configs"
}

// Register implements core.Handler.
func (s *SystemConfigHandler) Register(router *core.RouterGroup) {

	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Path:   "/get/:groupId",
	}).GET("/get/:groupId", func(ctx *gin.Context) {
		groupId, err := strconv.ParseUint(ctx.Param("groupId"), 10, 32)
		if err != nil {
			response.Error(response.ErrBadRequest, "").Done(ctx)
		}
		response.Success(s.service.GetConfigByGroup(groupId)).Done(ctx)
	})

	auth := router.Group("").Use(s.authMidds.MustAuthMiddleware)
	{
		// 系统配置
		auth.WithDoc(&core.DocItem{
			Method: http.MethodGet,
			Path:   "/list",
		}).GET("/list", s.curd.GetListHandler(&[]model.SystemConfig{}))

		// 创建
		auth.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/create",
		}).POST("/create", func(ctx *gin.Context) {
			s.curd.CreateHandler(ctx, &model.SystemConfig{}, func(model interface{}) error {
				return nil
			})
		})

		// 更新
		auth.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/update",
		}).POST("/update", func(ctx *gin.Context) {
			s.curd.UpdateHandler(ctx, &model.SystemConfig{}, func(model interface{}) error {
				return nil
			})
		})

		// 删除
		auth.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/delete",
		}).POST("/delete", func(ctx *gin.Context) {
			s.curd.DeleteHandler(ctx, nil)
		})

		// 导出
		auth.WithDoc(&core.DocItem{
			Method: http.MethodGet,
			Path:   "/export",
		}).GET("/export", s.curd.ExportHanderWithWhere(&[]model.SystemConfig{}, nil))

	}

	// 设置分组
	groups := router.Group("/groups")
	auth_groups := groups.Group("").Use(s.authMidds.MustAuthMiddleware)
	{
		auth_groups.WithDoc(&core.DocItem{
			Method: http.MethodGet,
			Path:   "/list",
		}).GET("/list", s.groupCurd.GetListHandler(&[]model.SystemConfigGroup{}))

		auth_groups.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/create",
		}).POST("/create", func(ctx *gin.Context) {
			s.groupCurd.CreateHandler(ctx, &model.SystemConfigGroup{}, func(model interface{}) error {
				return nil
			})
		})

		auth_groups.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/update",
		}).POST("/update", func(ctx *gin.Context) {
			s.groupCurd.UpdateHandler(ctx, &model.SystemConfigGroup{}, func(model interface{}) error {
				return nil
			})
		})

		auth_groups.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/delete",
		}).POST("/delete", func(ctx *gin.Context) {
			s.groupCurd.DeleteHandler(ctx, func(model interface{}) error {
				return nil
			})
		})

	}
}
