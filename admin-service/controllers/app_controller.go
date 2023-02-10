package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/topgs/trinit/admin-service/middlewares"
	"github.com/topgs/trinit/admin-service/models"
	"github.com/topgs/trinit/admin-service/services"
	"github.com/topgs/trinit/admin-service/utils"
)

type appController struct {
	as services.AppService
}

type AppController interface {
	RegisterApp(c echo.Context) error
	RegisterRule(c echo.Context) error
	UpdateRule(c echo.Context) error
	DeleteRule(c echo.Context) error
	GetRules(c echo.Context) error
}

func NewAppController(ts services.AppService) AppController {
	return &appController{ts}
}

func (ac *appController) RegisterApp(c echo.Context) error {
	user, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Invalid User")
	}
	request := new(models.AppRequest)
	if err := c.Bind(request); err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	res, err := ac.as.AddApp(user.ID, *request)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occured")
	}
	return middlewares.Responder(c, http.StatusOK, res)
}

func (ac *appController) RegisterRule(c echo.Context) error {
	_, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Invalid User")
	}
	request := new(models.RuleRequest)
	if err := c.Bind(request); err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	err = ac.as.AddAppRule(*request)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occured")
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

func (ac *appController) UpdateRule(c echo.Context) error {
	_, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Invalid User")
	}
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 32)
	request := new(models.RuleRequest)
	if err := c.Bind(request); err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	err = ac.as.UpdateAppRule(uint(id), *request)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occured")
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

func (ac *appController) GetRules(c echo.Context) error {
	_, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Invalid User")
	}
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 32)
	res, err := ac.as.GetAppRules(uint(id))
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occured")
	}
	return middlewares.Responder(c, http.StatusOK, res)
}

func (ac *appController) DeleteRule(c echo.Context) error {
	_, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Invalid User")
	}
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 32)
	err = ac.as.DeleteAppRule(uint(id))
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occured")
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}
