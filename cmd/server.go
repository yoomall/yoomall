package main

import (
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lazyfury.github.com/yoomall-server/app"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/constants"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/docs"
	"lazyfury.github.com/yoomall-server/driver"
)

//	@title						Nunu Example API
//	@version					1.0.0
//	@description				This is a sample server celler server.
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:8900
//	@securityDefinitions.apiKey	Bearer
//	@in							header
//	@name						Authorization
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	log.Info("hello world")
	instance := gin.Default()
	instance.SetTrustedProxies(nil)

	// config
	constants.CONFIG = config.GetConfig("./config.yaml")
	log.Info("config load success", "config", constants.CONFIG)

	// db
	constants.DB = driver.NewDB(constants.CONFIG.MysqlDsn())
	constants.DB.AutoMigrate() // 自动迁移
	log.Info("auto migrate success")
	// db

	docs.SwaggerInfo.BasePath = "/api"
	instance.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true)))

	// apps
	defApp := app.NewDefaultApp(instance, instance.Group("/api"), constants.CONFIG)

	register(defApp)

	log.Info("[server on] http://127.0.0.1:" + strconv.Itoa(constants.CONFIG.HTTP.Port))
	instance.Run(":" + strconv.Itoa(constants.CONFIG.HTTP.Port))
}

func register(apps ...core.App) {

	for _, app := range apps {
		log.Info("register app", "app", app.GetName())
		app.Migrate()
		app.GetRouter().Use(app.Middleware()...)
		app.Register()
		log.Info("register app success", "app", app.GetName())
	}
}
