package repositories

import (
	"github.com/topgs/trinit/admin-service/models"
	"github.com/topgs/trinit/admin-service/schemas"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type appRepository struct {
	db *gorm.DB
}

type AppRepository interface {
	CreateApp(app *schemas.App) error
	CreateRule(rule *models.RuleRequest) error
	FindRulesByAppID(id uint) ([]models.Rule, error)
	UpdateRuleByID(ruleId uint, rule models.RuleRequest) error
	DeleteRuleByID(id uint) error
}

func NewAppRepository(db *gorm.DB) AppRepository {
	return &appRepository{db}
}

func (ar *appRepository) CreateApp(app *schemas.App) error {
	return ar.db.Create(app).Error
}

func (ar *appRepository) CreateRule(req *models.RuleRequest) error {

	parameter := schemas.Parameter{
		AppID: req.AppID,
		Name:  req.Rule.Parameter,
	}

	err := ar.db.Create(&parameter).Error

	if err != nil {
		return err
	}

	var match_type schemas.MatchType

	err = ar.db.Where("name = ?", req.Rule.MatchType).First(&match_type).Error

	rule := schemas.Rule{
		ParameterID: parameter.ID,
		MatchTypeID: match_type.ID,
	}

	return ar.db.Create(&rule).Error
}

func (ar *appRepository) UpdateRuleByID(ruleId uint, rule models.RuleRequest) error {
	var match_type schemas.MatchType
	query := ar.db.Where("name = ?", rule.Rule.MatchType).First(&match_type)
	if query.Error != nil {
		return query.Error
	}
	return ar.db.Model(&schemas.Rule{}).Where(
		"id = ?", ruleId).Update(
		"match_type_id", match_type.ID).Error
}

func (ar *appRepository) DeleteRuleByID(id uint) error {
	return ar.db.Unscoped().Where("id = ?", id).Delete(&schemas.Rule{}).Error
}

func (ar *appRepository) FindRulesByAppID(id uint) ([]models.Rule, error) {
	var rules []models.Rule
	err := ar.db.Preload(clause.Associations).Table("rules").Select(
		"parameters.name as parameter, match_types.name as match_type",
	).Joins(
		"JOIN parameters on parameters.id = rules.parameter_id and parameters.app_id = ?", id).Joins(
		"JOIN match_types on match_types.id = rules.match_type_id").Scan(&rules).Error
	return rules, err
}
