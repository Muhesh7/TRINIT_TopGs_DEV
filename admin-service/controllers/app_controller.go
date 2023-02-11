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

// App godoc
//
//	@Summary		Register app
//	@Description	Register app
//	@Tags			App
//	@Accept					json
//	@Produce		json
//	@Param					app	body		models.AppRequest	true	"Regsiter app"
//	@Success		200	{object}	models.AppResponse
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
//
//	@Router			/v1/app/registerapp [post]
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

// App godoc
//
//	@Summary		Register rule
//	@Description	Register rule
//	@Tags			App
//	@Accept					json
//	@Produce		json
//	@Param					rule	body		models.RuleRequest	true	"Regsiter app"
//	@Success		200	{object}	models.AppResponse
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
//
//	@Router			/v1/app/registerrule [post]
func (ac *appController) RegisterRule(c echo.Context) error {
	request := new(models.RuleRequest)
	if err := c.Bind(request); err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	err := ac.as.AddAppRule(request)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occured")
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

// App godoc
//
//	@Summary		Update rule
//	@Description	Update rule
//	@Tags			App
//	@Accept					json
//	@Produce		json
//	@Param					rule	body		models.RuleRequest	true	"updated rule"
//	@Param					id	path		string	true	"rule id"
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
//
//	@Router			/v1/app/updaterule/{id} [put]
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

// App godoc
//
//	@Summary		Get rules
//	@Description	Get rules
//	@Tags			App
//	@Accept					json
//	@Produce		json
//	@Param					id	path		string	true	"app id"
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
//
//	@Router			/v1/app/getrules/{id} [get]
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

// App godoc
//
//	@Summary		Delete rules
//	@Description	Delete rules
//	@Tags			App
//	@Accept					json
//	@Produce		json
//	@Param					id	path		string	true	"rule id"
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
//
//	@Router			/v1/app/deleterule/{id} [delete]
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
