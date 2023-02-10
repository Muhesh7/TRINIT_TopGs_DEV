package router

import (
	"github.com/labstack/echo/v4"
	"github.com/topgs/trinit/admin-service/controllers"
)

func UserRoutes(e *echo.Group, c controllers.UserController) {
	user := e.Group("/user")

	user.POST("/signup", c.Register)
	user.GET("/verifyemail/:verificationCode", c.VerifyEmail)
	user.POST("/signin", c.Login)
}
