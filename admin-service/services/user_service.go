package services

import (
	"fmt"

	"github.com/topgs/trinit/admin-service/models"
	"github.com/topgs/trinit/admin-service/repositories"
	"github.com/topgs/trinit/admin-service/schemas"
	"github.com/topgs/trinit/admin-service/utils"
)

type userService struct {
	repo repositories.UserRepository
}

type UserService interface {
	RegisterUser(request models.RegisterRequest) (schemas.User, error)
	LoginUser(request models.LoginRequest) (string, error)
	AddVerificationCode(user *schemas.User, code string) error
	CompleteVerification(code string) error
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (us *userService) LoginUser(request models.LoginRequest) (string, error) {
	var user schemas.User
	token := ""
	reqEmail, err := utils.EmailValidator(request.Email)
	if err != nil {
		return token, err
	}
	user, err = us.repo.GetUserByEmail(reqEmail)
	if !user.Verified {
		return token, fmt.Errorf("Mail Not Verified")
	}
	DbPassword := user.Password
	reqPassword := utils.PasswordHasher(request.Password)

	if DbPassword == reqPassword {
		token, err = utils.CreateToken(user.Email)
	} else {
		err = fmt.Errorf("Wrong Password")
	}

	return token, err
}

func (us *userService) RegisterUser(request models.RegisterRequest) (schemas.User, error) {
	var user schemas.User

	email, err := utils.EmailValidator(request.Email)
	if err != nil {
		return user, err
	}

	password := utils.PasswordHasher(request.Password)

	user = schemas.User{
		Email:    email,
		Password: password,
	}

	err = us.repo.CreateUser(&user)

	return user, err
}

func (us *userService) AddVerificationCode(user *schemas.User, code string) error {
	return us.repo.UpdateVerificationCode(user, code)
}

func (us *userService) CompleteVerification(code string) error {
	return us.repo.UpdateVerification(code)
}
