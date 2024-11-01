// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"yoomall/apps/app"
	"yoomall/apps/app/handler"
	"yoomall/apps/auth"
	handler2 "yoomall/apps/auth/handler"
	"yoomall/apps/auth/middleware"
	"yoomall/apps/auth/service"
	"yoomall/apps/common"
	handler3 "yoomall/apps/common/handler"
	"yoomall/apps/common/service"
	"yoomall/apps/post"
	"yoomall/apps/views"
	"yoomall/core/driver"
	"yoomall/core/http"
)

// Injectors from wire.go:

func NewApp(conf *viper.Viper, db *driver.DB, setupEngine func(*gin.Engine) *gin.Engine) httpserver.HttpServer {
	dtkHandler := handler.NewDtkHandler(conf)
	authMiddlewareGroup := authmiddleware.NewAuthMiddlewareGroup(db)
	menuHandler := handler.NewMenuHandler(db, authMiddlewareGroup)
	jtkHandler := handler.NewJtkHandler(conf)
	defaultApp := app.NewWireDefaultApp(conf, db, dtkHandler, menuHandler, jtkHandler)
	authService := authservice.NewAuthService(db)
	userHandler := handler2.NewUserHandler(db, conf, authService, authMiddlewareGroup)
	userRoleHandler := handler2.NewUserRoleHandler(db, authMiddlewareGroup)
	userTokenHandler := handler2.NewUserTokenHandler(db, authMiddlewareGroup)
	permissionHandler := handler2.NewPermissionHandler(db, authMiddlewareGroup)
	authApp := auth.NewAuthApp(conf, db, authService, userHandler, userRoleHandler, userTokenHandler, permissionHandler)
	postApp := post.NewDefaultApp(conf, db)
	notFoundRecordService := commonservice.NewNotFoundRecordService(db)
	notFoundRecordHandler := handler3.NewNotFoundRecordHandler(db, notFoundRecordService)
	systemConfigService := commonservice.NewSystemConfigService(db)
	systemConfigHandler := handler3.NewSystemConfigHandler(db, systemConfigService, authMiddlewareGroup)
	commonApp := common.NewCommonApp(conf, db, notFoundRecordHandler, systemConfigHandler)
	viewsApp := views.NewViewApp(db, conf)
	doc := NewDoc()
	httpServer := NewHttpServer(conf, defaultApp, authApp, postApp, commonApp, viewsApp, notFoundRecordService, doc, setupEngine)
	return httpServer
}
