package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/topgs/trinit/admin-service/controllers"
)

func NewRouter(e *echo.Echo, c controllers.MainController) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}}))

	api := e.Group("/v1")

	UserRoutes(api, c.User)
	AppRoutes(api, c.App)
	SwaggerRoutes(api)
}
