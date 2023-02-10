package schemas

import (
	"gorm.io/gorm"
)

type App struct {
	gorm.Model
	UserID uint `gorm:"not null;"`
	User   User
	Name   string `gorm:"not null;unique"`
	Secret string `gorm:"not null;"`
}

type Parameter struct {
	gorm.Model
	AppID uint `gorm:"not null;"`
	App   App
	Name  string `gorm:"not null"`
}

type MatchType struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}

type Rule struct {
	gorm.Model
	ParameterID uint `gorm:"not null;"`
	Parameter   Parameter
	MatchTypeID uint `gorm:"not null;"`
	MatchType   MatchType
}
