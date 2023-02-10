package repositories

import (
	"fmt"

	"github.com/topgs/trinit/admin-service/schemas"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	CreateUser(user *schemas.User) error
	GetUserByEmail(email string) (schemas.User, error)
	UpdateVerificationCode(user *schemas.User, code string) error
	UpdateVerification(code string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(email string) (schemas.User, error) {
	var user schemas.User
	err := ur.db.First(&user, "email = ?", email).Error
	return user, err
}

func (ur *userRepository) CreateUser(user *schemas.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) UpdateVerificationCode(user *schemas.User, code string) error {
	return ur.db.Model(&user).Update("verification_code", code).Error
}

func (ur *userRepository) UpdateVerification(code string) error {
	var err error
	res := ur.db.Model(&schemas.User{}).Where(
		"verification_code = ?", code).Updates(
		map[string]interface{}{"verified": true, "verification_code": gorm.Expr("NULL")})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		err = fmt.Errorf("No Matching User Found")
	}
	return err
}
