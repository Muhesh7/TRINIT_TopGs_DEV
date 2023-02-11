package router

import (
	"github.com/labstack/echo/v4"
	"github.com/topgs/trinit/admin-service/controllers"
	"github.com/topgs/trinit/admin-service/middlewares"
)

func AppRoutes(e *echo.Group, c controllers.AppController) {
	app := e.Group("/app")

	app.POST("/registerapp", middlewares.Authorizer(c.RegisterApp))
	app.POST("/registerrule", middlewares.Authorizer(c.RegisterRule))
	app.GET("/getrules/:id", middlewares.Authorizer(c.GetRules))
	app.DELETE("/deleterule/:id", middlewares.Authorizer(c.DeleteRule))
	app.PUT("/updaterule/:id", middlewares.Authorizer(c.UpdateRule))
}
