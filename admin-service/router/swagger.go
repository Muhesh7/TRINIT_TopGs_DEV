package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/topgs/trinit/admin-service/config"
)

func SwaggerRoutes(e *echo.Group) {
	origin := config.Origin
	url := echoSwagger.URL(origin + "/v1/admin/doc.json")
	e.GET("/admin/*", echoSwagger.EchoWrapHandler(url))
}
