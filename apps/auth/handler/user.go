package handler

import (
	"net/http"

	authmiddleware "yoomall/apps/auth/middleware"
	"yoomall/apps/auth/model"
	"yoomall/apps/auth/request"
	authservice "yoomall/apps/auth/service"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/lazyfury/pulse/helper/curd"
	"github.com/lazyfury/pulse/helper/response"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserHandler struct {
	CRUD      *curd.CRUD
	service   *authservice.AuthService
	authMidds *authmiddleware.AuthMiddlewareGroup
}

var _ framework.Handler = (*UserHandler)(nil)

func NewUserHandler(db *driver.DB, config *viper.Viper, service *authservice.AuthService, authMiddlewareGroup *authmiddleware.AuthMiddlewareGroup) *UserHandler {
	return &UserHandler{
		CRUD:      curd.New(db, &model.User{}),
		service:   service,
		authMidds: authMiddlewareGroup,
	}
}

func (u *UserHandler) Register(router *framework.RouterGroup) {
	// 登录接口
	router.POST("/login", u.loginWithUsernameAndPassword).Doc(&framework.DocItem{
		Method: http.MethodPost,
		Path:   "/login",
		Body:   request.UserUserNameAndPasswordLoginRequest{},
	})

	// 用户列表
	auth := router.Group("").Use(u.authMidds.MustAuthMiddleware)
	{
		auth.GET("/user-list", u.userList).Doc(&framework.DocItem{
			Method:      http.MethodGet,
			Path:        "/user-list",
			Title:       "用户列表",
			Tag:         "auth",
			Description: "用户列表",
			Params:      nil,
			Success:     nil,
			Failure:     nil,
		})
	}

	authWithUser := router.Group("").Use(u.authMidds.MustAuthMiddlewareWithUser)
	{
		authWithUser.Doc(&framework.DocItem{
			Method: http.MethodGet,
			Path:   "/profile",
		}).GET("/profile", func(ctx *gin.Context) {
			response.Success(ctx.MustGet("user")).Done(ctx)
		})

		authWithUser.Doc(&framework.DocItem{
			Method: http.MethodGet,
			Path:   "/logout",
		}).POST("/logout", func(ctx *gin.Context) {
			u.service.Logout(ctx)
		})
	}

}

func (u *UserHandler) GetRouterGroupName() string {
	return "users"
}

func (u *UserHandler) userList(ctx *gin.Context) {
	u.CRUD.GetListHandlerWithWhere(&[]model.User{}, func(tx *gorm.DB) *gorm.DB {
		return tx
	})(ctx)
}

func (u *UserHandler) loginWithUsernameAndPassword(ctx *gin.Context) {
	var data request.UserUserNameAndPasswordLoginRequest
	ctx.ShouldBindBodyWithJSON(&data)
	result := u.service.LoginWithUsernameAndPassword(data.UserName, data.Password, ctx)

	if result.IsErr() {
		response.Error(response.ErrBadRequest, result.Error.Error()).Done(ctx)
		return
	}

	response.Success(result.Value).Done(ctx)
}
