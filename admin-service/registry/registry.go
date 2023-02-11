package registry

import (
	"github.com/topgs/trinit/admin-service/controllers"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewMainController() controllers.MainController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewMainController() controllers.MainController {
	return controllers.MainController{
		User: r.NewUserController(),
		App:  r.NewAppController(),
	}
}
