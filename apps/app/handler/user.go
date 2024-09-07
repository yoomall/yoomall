package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/app/model"
	"lazyfury.github.com/yoomall-server/apps/app/request"
	"lazyfury.github.com/yoomall-server/apps/app/service"
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

func (u *UserHandler) Register(router *gin.RouterGroup) {
	router.POST("/login", u.LoginWithUsernameAndPassword)
	router.GET("/list", u.CRUD.GetListHandler(&[]model.User{}))
}

func (u *UserHandler) GetRouterGroupName() string {
	return "users"
}

func (u *UserHandler) LoginWithUsernameAndPassword(ctx *gin.Context) {
	var data request.UserUserNameAndPasswordLoginRequest
	ctx.ShouldBindBodyWithJSON(&data)
	user, token, err := u.service.LoginWithUsernameAndPassword(data.UserName, data.Password)

	if err != nil {
		response.Error(response.ErrInternalError, err.Error()).Done(ctx)
		return
	}
	response.Success(map[string]any{
		"user":  user,
		"token": token,
	}).Done(ctx)
}
