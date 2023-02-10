package registry

import "github.com/topgs/trinit/admin-service/services"

func (r *registry) NewMailService() services.MailService {
	return services.NewMailService()
}
