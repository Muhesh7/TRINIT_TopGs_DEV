package main

import (
	"github.com/labstack/echo/v4"
	"github.com/topgs/trinit/admin-service/config"
	_ "github.com/topgs/trinit/admin-service/docs"
	"github.com/topgs/trinit/admin-service/registry"
	"github.com/topgs/trinit/admin-service/router"
)

//	@title			trinit Admin
//	@version		1.0

//	@description	Admin Panel for trinit Application.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	topgs
//	@contact.email	hacksparrow169@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

func main() {

	// Load config
	config.InitApp()

	// Register app
	reg := registry.NewRegistry(config.GetDB())

	appController := reg.NewAppController()

	// Seed database
	// appController.Seed.SeedDB()

	// Create and Setup Echo Server
	server := echo.New()

	// Create Router
	router.NewRouter(server, appController)

	server.Logger.Fatal(server.Start(":" + config.ServerPort))

}
