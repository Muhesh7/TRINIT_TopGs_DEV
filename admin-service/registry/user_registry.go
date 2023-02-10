package registry

import (
	"github.com/topgs/trinit/admin-service/controllers"
	"github.com/topgs/trinit/admin-service/repositories"
	"github.com/topgs/trinit/admin-service/services"
)

func (r *registry) NewUserController() controllers.UserController {
	return controllers.NewUserController(r.NewUserService(), r.NewMailService())
}

func (r *registry) NewUserService() services.UserService {
	return services.NewUserService(r.NewUserRepository())
}

func (r *registry) NewUserRepository() repositories.UserRepository {
	return repositories.NewUserRepository(r.db)
}
