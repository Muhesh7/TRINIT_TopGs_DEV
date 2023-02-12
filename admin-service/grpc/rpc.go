package rpc

import (
	"context"

	"github.com/topgs/trinit/admin-service/config"
	"github.com/topgs/trinit/admin-service/gen/app"
	"github.com/topgs/trinit/admin-service/schemas"
	"gorm.io/gorm/clause"
)

type AuthRPCServer struct {
	app.UnimplementedAuthServiceServer
}

func (AuthRPCServer) AuthRPC(ctx context.Context, request *app.AppAuthRequest) (*app.AppAuthResponse, error) {
	db := config.GetDB()

	var app_details schemas.App
	err := db.Where("name = ?", request.AppId).First(&app_details).Error

	if err != nil {
		return &app.AppAuthResponse{
			IsSuccess: false,
		}, err
	}

	if app_details.Secret != request.AppSecret {
		return &app.AppAuthResponse{
			IsSuccess: false,
		}, err
	}
	var rules []*app.Rule

	err = db.Preload(clause.Associations).Table("rules").Select(
		"parameters.name as parameter, match_types.name as match_type",
	).Joins(
		"JOIN parameters on parameters.id = rules.parameter_id and parameters.app_id = ?", app_details.ID).Joins(
		"JOIN match_types on match_types.id = rules.match_type_id").Scan(&rules).Error

	return &app.AppAuthResponse{
		IsSuccess: true,
		Rule:      rules}, nil
}
