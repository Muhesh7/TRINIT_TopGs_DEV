package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/topgs/trinit/admin-service/config"
	"github.com/topgs/trinit/admin-service/middlewares"
	"github.com/topgs/trinit/admin-service/models"
	"github.com/topgs/trinit/admin-service/services"
	"github.com/topgs/trinit/admin-service/utils"
)

type userController struct {
	us services.UserService
	ms services.MailService
}

type UserController interface {
	Register(c echo.Context) error
	VerifyEmail(c echo.Context) error
	Login(c echo.Context) error
}

func NewUserController(us services.UserService, ms services.MailService) UserController {
	return &userController{us, ms}
}

// Register godoc
//
//	@Summary		Register an user
//	@Description	register an user
//	@Tags			User
//	@Accept					json
//	@Produce		json
//	@Param					user	body		models.RegisterRequest	true	"Add user"
//	@Success		200	{object}	models.RegisterResponse
//	@Failure		400	{object}	models.Error
//	@Router			/v1/user/signup [post]
func (uc *userController) Register(c echo.Context) error {
	request := new(models.RegisterRequest)
	log.Print(request)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}

	user, err := uc.us.RegisterUser(*request)

	if err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Already Registered")
	}

	code := utils.GenerateVerificationCode()
	code = code + user.Email

	emailData := &models.EmailData{
		Name:    user.Email,
		URL:     config.Origin + "/v1/user/verifyemail/" + code,
		Subject: "user verification",
	}

	err = uc.us.AddVerificationCode(&user, code)

	if err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Unable to Register")
	}

	err = uc.ms.MailUser(user, *emailData)

	if err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Unable to Send Email")
	}

	return middlewares.Responder(c, http.StatusOK, "Check Your Email for verification mail")
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			User
//	@Accept					json
//	@Produce		json
//	@Param					user	body		models.LoginRequest	true	"Authenticate user"
//	@Success		200	{object}	models.LoginResponse
//	@Failure		400	{object}	models.Error
//	@Router			/v1/user/signin [post]
func (uc *userController) Login(c echo.Context) error {
	request := new(models.LoginRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}

	token, err := uc.us.LoginUser(*request)
	if err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, err.Error())
	}

	return middlewares.Responder(c, http.StatusOK, models.LoginResponse{Token: token, Message: "Login success"})
}

// VerifyEmail godoc
//
//	@Summary		verify an user
//	@Description	verify an user by sending email
//	@Tags			User
//	@Accept					json
//	@Produce		json
//	@Param					verificationCode	path  string true "Verify user"
//	@Success		200	{object}	string
//	@Failure		400	{object}	models.Error
//	@Router			/v1/user/verifyemail/{verificationCode} [post]
func (uc *userController) VerifyEmail(c echo.Context) error {

	code := c.Param("verificationCode")
	log.Println("VerificationCode", code)
	err := uc.us.CompleteVerification(code)
	if err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Unable to Verify")
	}
	return middlewares.Responder(c, http.StatusOK, "verified successfully")
}
