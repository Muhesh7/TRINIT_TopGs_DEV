package registry

import (
	"github.com/topgs/trinit/admin-service/controllers"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		User: r.NewUserController(),
	}
}
