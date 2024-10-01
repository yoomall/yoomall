package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	authmiddleware "lazyfury.github.com/yoomall-server/apps/auth/middleware"
	"lazyfury.github.com/yoomall-server/apps/common/model"
	commonservice "lazyfury.github.com/yoomall-server/apps/common/service"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/curd"
	"lazyfury.github.com/yoomall-server/core/helper/response"
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
		curd:      curd.New(db, &model.SystemConfig{}),
		authMidds: authMiddlewareGroup,
		groupCurd: curd.New(db, &model.SystemConfigGroup{}),
	}
}

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
		auth.WithDoc(&core.DocItem{
			Method: http.MethodGet,
			Path:   "/list",
		}).GET("/list", s.curd.GetListHandler(&[]model.SystemConfig{}))

		auth.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/create",
		}).POST("/create", func(ctx *gin.Context) {
			s.curd.CreateHandler(ctx, &model.SystemConfig{}, func(model interface{}) error {
				return nil
			})
		})

		auth.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/update",
		}).POST("/update", func(ctx *gin.Context) {
			s.curd.UpdateHandler(ctx, &model.SystemConfig{}, func(model interface{}) error {
				return nil
			})
		})

		auth.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/delete",
		}).POST("/delete", func(ctx *gin.Context) {
			s.curd.DeleteHandler(ctx, nil)
		})

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
