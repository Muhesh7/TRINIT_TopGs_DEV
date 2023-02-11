package models

type AppRequest struct {
	Name string `json:"name"`
}

type AppResponse struct {
	Secret string `json:"secret"`
}

type RuleRequest struct {
	AppID uint `json:"app_id"`
	Rule  Rule `json:"rule"`
}

type Rule struct {
	Parameter string `json:"parameter"`
	MatchType string `json:"match_type"`
}
