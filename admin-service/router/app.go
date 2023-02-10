package router

import (
	"github.com/labstack/echo/v4"
	"github.com/topgs/trinit/admin-service/controllers"
)

func AppRoutes(e *echo.Group, c controllers.AppController) {
	app := e.Group("/app")

	app.POST("/registerapp", c.RegisterApp)
	app.POST("/registerrule", c.RegisterRule)
	app.GET("/getrules:id", c.GetRules)
	app.DELETE("/deleterule:id", c.DeleteRule)
	app.PUT("/updaterule:id", c.DeleteRule)
}
