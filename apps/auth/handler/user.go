package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	authmiddleware "lazyfury.github.com/yoomall-server/apps/auth/middleware"
	"lazyfury.github.com/yoomall-server/apps/auth/model"
	"lazyfury.github.com/yoomall-server/apps/auth/request"
	"lazyfury.github.com/yoomall-server/apps/auth/service"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/curd"
	"lazyfury.github.com/yoomall-server/core/helper/response"
)

type UserHandler struct {
	CRUD    *curd.CRUD
	service *service.AuthService
}

var _ core.Handler = (*UserHandler)(nil)

func NewUserHandler(db *driver.DB, config *viper.Viper, service *service.AuthService) *UserHandler {
	return &UserHandler{
		CRUD: &curd.CRUD{
			DB:    db,
			Model: &model.User{},
		},
		service: service,
	}
}

func (u *UserHandler) Register(router *core.RouterGroup) {
	router.WithDoc(&core.DocItem{
		Method: http.MethodPost,
		Path:   "/login",
		Body:   request.UserUserNameAndPasswordLoginRequest{},
	}).POST("/login", u.LoginWithUsernameAndPassword)

	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Path:   "/user-list",
	}).Use(authmiddleware.AuthMiddleware(u.CRUD.DB, true, false)).GET("/user-list", u.CRUD.GetListHandlerWithWhere(&[]model.User{}, func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("Ext")
	}))
}

func (u *UserHandler) GetRouterGroupName() string {
	return "users"
}

func (u *UserHandler) LoginWithUsernameAndPassword(ctx *gin.Context) {
	var data request.UserUserNameAndPasswordLoginRequest
	ctx.ShouldBindBodyWithJSON(&data)
	result := u.service.LoginWithUsernameAndPassword(data.UserName, data.Password)

	if result.IsErr() {
		response.Error(response.ErrBadRequest, result.Error.Error()).Done(ctx)
		return
	}

	response.Success(result.Value).Done(ctx)
}
