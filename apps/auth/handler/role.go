package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authmiddleware "lazyfury.github.com/yoomall-server/apps/auth/middleware"
	"lazyfury.github.com/yoomall-server/apps/auth/model"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/curd"
	"lazyfury.github.com/yoomall-server/core/helper/response"
)

type UserRoleHandler struct {
	DB        *driver.DB
	CRUD      *curd.CRUD
	authMidds *authmiddleware.AuthMiddlewareGroup
}

// GetRouterGroupName implements core.Handler.
func (u *UserRoleHandler) GetRouterGroupName() string {
	return "user-roles"
}

var _ core.Handler = (*UserRoleHandler)(nil)

func NewUserRoleHandler(db *driver.DB, authMidds *authmiddleware.AuthMiddlewareGroup) *UserRoleHandler {
	return &UserRoleHandler{
		DB:        db,
		CRUD:      curd.New(db, &model.UserRole{}),
		authMidds: authMidds,
	}
}

func (u *UserRoleHandler) Register(router *core.RouterGroup) {
	auth := router.Group("").Use(u.authMidds.MustAuthMiddleware)
	{
		auth.WithDoc(&core.DocItem{
			Method: http.MethodGet,
			Path:   "/role-list",
		}).GET("/role-list", u.CRUD.GetListHandler(&[]model.UserRole{}))

		auth.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/create-role",
		}).POST("/create-role", u.create)

		auth.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/update-role",
		}).POST("/update-role", u.update)

		auth.WithDoc(&core.DocItem{
			Method: http.MethodPost,
			Path:   "/delete-role",
		}).POST("/delete-role", u.delete)

	}
}

func (u *UserRoleHandler) create(ctx *gin.Context) {
	var role *model.UserRole = &model.UserRole{}
	if err := ctx.ShouldBindBodyWithJSON(role); err != nil {
		response.Error(response.ErrBadRequest, "获取参数错误:"+err.Error()).Done(ctx)
		return
	}

	var find []*model.UserRole
	u.DB.Where("role_code = ?", role.RoleCode).Find(&find)
	if len(find) > 0 {
		response.Error(response.ErrBadRequest, "角色编码已存在").Done(ctx)
		return
	}

	if err := u.DB.Create(role).Error; err != nil {
		response.Error(response.ErrInternalError, "保存失败").Done(ctx)
		return
	}

	response.Success(role).Done(ctx)
}

func (u *UserRoleHandler) update(ctx *gin.Context) {
	var role *model.UserRole = &model.UserRole{}
	if err := ctx.ShouldBindBodyWithJSON(role); err != nil {
		response.Error(response.ErrBadRequest, "获取参数错误:"+err.Error()).Done(ctx)
		return
	}

	var find []*model.UserRole
	u.DB.Where("role_code = ?", role.RoleCode).Where("id != ?", role.ID).Find(&find)
	if len(find) > 0 {
		response.Error(response.ErrBadRequest, "角色编码已存在").Done(ctx)
		return
	}

	if err := u.DB.Save(role).Error; err != nil {
		response.Error(response.ErrInternalError, "保存失败").Done(ctx)
		return
	}

	response.Success(role).Done(ctx)
}

func (u *UserRoleHandler) delete(ctx *gin.Context) {
	type data struct {
		ID uint `json:"id"`
	}

	var d = &data{}

	if err := ctx.ShouldBindBodyWithJSON(d); err != nil {
		response.Error(response.ErrBadRequest, "参数错误:"+err.Error()).Done(ctx)
		return
	}

	if err := u.DB.Delete(&model.UserRole{}, "id = ?", d.ID).Error; err != nil {
		response.Error(response.ErrBadRequest, "删除失败:"+err.Error()).Done(ctx)
		return
	}

	response.Success("ok").Done(ctx)

}
