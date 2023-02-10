package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/topgs/trinit/admin-service/config"
	"github.com/topgs/trinit/admin-service/utils"
)

func Authorizer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if len(authHeader) == 0 {
			return Responder(c, http.StatusBadRequest, "please set Header Authorization")
		}
		user, err := utils.GetCurrentUserFromToken(authHeader, config.GetDB())
		if err != nil {
			return Responder(c, http.StatusBadRequest, "Invalid User-Token")
		}
		c.Set("user", user)
		return next(c)
	}
}
