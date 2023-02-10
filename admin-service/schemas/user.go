package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"not null;unique"`
	Email            string `gorm:"not null;unique"`
	Password         string `gorm:"not null;"`
	VerificationCode string `gorm:"default:null;"`
	Verified         bool   `gorm:"default:false;"`
}
