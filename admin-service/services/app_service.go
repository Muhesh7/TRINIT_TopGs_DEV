package services

import (
	"github.com/topgs/trinit/admin-service/models"
	"github.com/topgs/trinit/admin-service/repositories"
	"github.com/topgs/trinit/admin-service/schemas"
	"github.com/topgs/trinit/admin-service/utils"
)

type appService struct {
	repo repositories.AppRepository
}

type AppService interface {
	AddApp(userId uint,
		request models.AppRequest) (models.AppResponse, error)
	AddAppRule(rule models.RuleRequest) error
	GetAppRules(appId uint) ([]models.Rule, error)
	UpdateAppRule(ruleId uint, rule models.RuleRequest) error
	DeleteAppRule(ruleId uint) error
}

func NewAppService(repo repositories.AppRepository) AppService {
	return &appService{repo}
}

func (as *appService) AddApp(userId uint,
	request models.AppRequest) (models.AppResponse, error) {

	var response models.AppResponse
	secret := utils.GenerateSecret()
	app := schemas.App{
		UserID: userId,
		Name:   request.Name,
		Secret: secret,
	}

	err := as.repo.CreateApp(&app)
	if err != nil {
		return response, err
	}
	response.Secret = secret
	return response, nil
}

func (as *appService) AddAppRule(rule models.RuleRequest) error {
	return as.repo.CreateRule(rule)
}

func (as *appService) GetAppRules(appId uint) ([]models.Rule, error) {
	return as.repo.FindRulesByAppID(appId)
}

func (as *appService) UpdateAppRule(ruleId uint, rule models.RuleRequest) error {
	return as.repo.UpdateRuleByID(ruleId, rule)
}

func (as *appService) DeleteAppRule(ruleId uint) error {
	return as.repo.DeleteRuleByID(ruleId)
}
