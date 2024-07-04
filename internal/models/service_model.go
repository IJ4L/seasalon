package models

import (
	"database/sql"
	db "gitlab/go-prolog-api/example/db/sqlc"
)

type CreateServiceRequest struct {
	Idbranch int32  `form:"idbranch"`
	Name     string `form:"name"`
	Type     string `form:"type"`
	Detail   string `form:"detail"`
	Pricing  string `form:"pricing"`
	Duration int32  `form:"duration"`
}

type ServiceResponse struct {
	ID        int32        `json:"id"`
	Idbranch  int32        `json:"idbranch"`
	Name      string       `json:"name"`
	Type      string       `json:"type"`
	Detail    string       `json:"detail"`
	Pricing   string       `json:"pricing"`
	Duration  int32        `json:"duration"`
	Createdat sql.NullTime `json:"createdat"`
	Updatedat sql.NullTime `json:"updatedat"`
}

func NewServiceResponse(service db.Service) ServiceResponse {
	return ServiceResponse{
		ID:        service.ID,
		Idbranch:  service.Idbranch,
		Name:      service.Name,
		Type:      service.Type,
		Detail:    service.Detail,
		Pricing:   service.Pricing,
		Duration:  service.Duration,
		Createdat: service.Createdat,
		Updatedat: service.Updatedat,
	}
}

func NewServicesResponse(service []db.Service) []ServiceResponse {
	response := make([]ServiceResponse, 0)
	for _, service := range service {
		response = append(response, ServiceResponse{
			ID:        service.ID,
			Idbranch:  service.Idbranch,
			Name:      service.Name,
			Type:      service.Type,
			Detail:    service.Detail,
			Pricing:   service.Pricing,
			Duration:  service.Duration,
			Createdat: service.Createdat,
			Updatedat: service.Updatedat,
		})
	}
	return response
}

type UpdateServicesRequest struct {
	ID       int32  `form:"id"`
	Name     string `form:"name"`
	Type     string `form:"type"`
	Detail   string `form:"detail"`
	Pricing  string `form:"pricing"`
	Duration int32  `form:"duration"`
}
