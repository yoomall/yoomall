package core

import (
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type App interface {
	GetName() string
	GetDB() *driver.DB
	Migrate()
	Middleware() []gin.HandlerFunc
	GetConfig() *config.Config
	Register(router *gin.RouterGroup)
}

type AppImpl struct {
	AppName   string
	AppConfig *config.Config
	db        *driver.DB
	Handlers  []Handler
}

func NewAppImpl(name string, config *config.Config, db *driver.DB, handlers []Handler) *AppImpl {
	return &AppImpl{
		AppName:   name,
		AppConfig: config,
		db:        db,
		Handlers:  handlers,
	}
}

func (a *AppImpl) GetConfig() *config.Config {
	return a.AppConfig
}

func (a *AppImpl) GetName() string {
	return a.AppName
}

func (a *AppImpl) GetDB() *driver.DB {
	return a.db
}
