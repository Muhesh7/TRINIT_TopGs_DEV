package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/topgs/trinit/admin-service/config"
	_ "github.com/topgs/trinit/admin-service/docs"
	"github.com/topgs/trinit/admin-service/gen/app"
	rpc "github.com/topgs/trinit/admin-service/grpc"
	"github.com/topgs/trinit/admin-service/middlewares"
	"github.com/topgs/trinit/admin-service/registry"
	"github.com/topgs/trinit/admin-service/router"
	"google.golang.org/grpc"
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

	appController := reg.NewMainController()

	// Seed database
	// appController.Seed.SeedDB()

	// Create and Setup Echo Server
	server := echo.New()

	// Create Router
	router.NewRouter(server, appController)

	var wg sync.WaitGroup
	wg.Add(2)
	// Start Server
	go func() {
		server.Logger.Fatal(server.Start(":" + config.ServerPort))
	}()

	grpcPort := config.RPCPort
	grpcServer := grpc.NewServer(middlewares.WithServerUnaryInterceptor())
	app.RegisterAuthServiceServer(grpcServer, &rpc.AuthRPCServer{})
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {
			log.Panic("grpc server running error on", err)
		}
		err1 := grpcServer.Serve(lis)
		if err1 != nil {
			log.Panic("grpc server running error on", err1)
		}
	}()

	wg.Wait()

}
