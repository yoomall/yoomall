package core

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type IApp interface {
	GetName() string
	GetDB() *driver.DB
	Migrate()
	Middleware() []gin.HandlerFunc
	GetConfig() *viper.Viper
	Register(router *RouterGroup)
	GetHandlers() []Handler
}

type App struct {
	AppName   string
	AppConfig *viper.Viper
	db        *driver.DB
	Handlers  []Handler
}

func NewApp(name string, config *viper.Viper, db *driver.DB, handlers []Handler) *App {
	return &App{
		AppName:   name,
		AppConfig: config,
		db:        db,
		Handlers:  handlers,
	}
}

func (a *App) GetConfig() *viper.Viper {
	return a.AppConfig
}

func (a *App) GetName() string {
	return a.AppName
}

func (a *App) GetDB() *driver.DB {
	return a.db
}

func (a *App) GetHandlers() []Handler {
	return a.Handlers
}

type RegisterApp struct {
	Router *RouterGroup
	App    IApp
}

func (instance *RegisterApp) Register() {
	log.Info(instance.App.GetName() + "====================================")
	log.Info("注册app", "app", instance.App.GetName())
	if config.Config.DEBUG {
		log.Info("迁移中", "app", instance.App.GetName())
		instance.App.Migrate()
		log.Info("迁移成功 success", "app", instance.App.GetName())
	}
	router := instance.Router.Group("")
	router.Use(instance.App.Middleware()...)
	instance.App.Register(router)

	for _, handler := range instance.App.GetHandlers() {
		handler.Register(router.Group(handler.GetRouterGroupName()))
	}

	log.Info("注册成功", "app", instance.App.GetName())
}
