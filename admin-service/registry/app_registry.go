package registry

import (
	"github.com/topgs/trinit/admin-service/controllers"
	"github.com/topgs/trinit/admin-service/repositories"
	"github.com/topgs/trinit/admin-service/services"
)

func (r *registry) NewAppController() controllers.AppController {
	return controllers.NewAppController(r.NewAppService())
}

func (r *registry) NewAppService() services.AppService {
	return services.NewAppService(r.NewAppRepository())
}

func (r *registry) NewAppRepository() repositories.AppRepository {
	return repositories.NewAppRepository(r.db)
}
