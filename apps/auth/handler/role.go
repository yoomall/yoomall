package handler

import (
	"net/http"

	authmiddleware "github.com/yoomall/yoomall/apps/auth/middleware"
	"github.com/yoomall/yoomall/apps/auth/model"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/lazyfury/pulse/helper/curd"
	"github.com/lazyfury/pulse/helper/response"
)

type UserRoleHandler struct {
	DB        *driver.DB
	CRUD      *curd.CRUD
	authMidds *authmiddleware.AuthMiddlewareGroup
}

// GetRouterGroupName implements yoo.Handler.
func (u *UserRoleHandler) GetRouterGroupName() string {
	return "user-roles"
}

var _ framework.Handler = (*UserRoleHandler)(nil)

func NewUserRoleHandler(db *driver.DB, authMidds *authmiddleware.AuthMiddlewareGroup) *UserRoleHandler {
	return &UserRoleHandler{
		DB:        db,
		CRUD:      curd.New(db, &model.UserRole{}),
		authMidds: authMidds,
	}
}

func (u *UserRoleHandler) Register(router *framework.RouterGroup) {
	auth := router.Group("").Use(u.authMidds.MustAuthMiddleware)
	{
		auth.GET("/role-list", u.CRUD.GetListHandler(&[]model.UserRole{})).Doc(&framework.DocItem{
			Method: http.MethodGet,
			Path:   "/role-list",
		})

		auth.POST("/create-role", u.create).Doc(&framework.DocItem{
			Method: http.MethodPost,
			Path:   "/create-role",
		})

		auth.POST("/update-role", u.update).Doc(&framework.DocItem{
			Method: http.MethodPost,
			Path:   "/update-role",
		})

		auth.POST("/delete-role", u.delete).Doc(&framework.DocItem{
			Method: http.MethodPost,
			Path:   "/delete-role",
		})

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
