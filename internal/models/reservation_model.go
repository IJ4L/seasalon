package models

import (
	"database/sql"
	db "gitlab/go-prolog-api/example/db/sqlc"
)

type CreateReservationRequest struct {
	Iduser    int32 `form:"iduser"`
	Idservice int32 `form:"idservice"`
}

type ReservationResponse struct {
	ID        int32        `json:"id"`
	Iduser    int32        `json:"iduser"`
	Fullname  string       `json:"fullname"`
	Photo     string       `json:"photo"`
	Email     string       `json:"email"`
	Idservice int32        `json:"idservice"`
	Name      string       `json:"name"`
	Type      string       `json:"type"`
	Detail    string       `json:"detail"`
	Pricing   string       `json:"pricing"`
	Duration  int32        `json:"duration"`
	Createdat sql.NullTime `json:"createdat"`
	Updatedat sql.NullTime `json:"updatedat"`
}

func NewReservationResponse(reservation db.Reservation, user db.User, service db.Service) ReservationResponse {
	return ReservationResponse{
		ID:        reservation.ID,
		Iduser:    reservation.Iduser,
		Name:      user.Fullname,
		Photo:     user.Photo,
		Email:     user.Email,
		Idservice: reservation.Idservice,
		Type:      service.Type,
		Detail:    service.Detail,
		Pricing:   service.Pricing,
		Duration:  service.Duration,
		Createdat: reservation.Createdat,
		Updatedat: reservation.Updatedat,
	}
}

func NewReservationsResponse(reservation []db.SelectReservationRow, user db.User, service db.Service) []ReservationResponse {
	response := make([]ReservationResponse, 0)
	for _, reservation := range reservation {
		response = append(response, ReservationResponse{
			ID:        reservation.ID,
			Name:      user.Fullname,
			Photo:     user.Photo,
			Email:     user.Email,
			Type:      service.Type,
			Detail:    service.Detail,
			Pricing:   service.Pricing,
			Duration:  service.Duration,
			Createdat: reservation.Createdat,
		})
	}
	return response
}

type UpdateReservationsRequest struct {
	ID        int32 `form:"id"`
	Iduser    int32 `form:"iduser"`
	Idservice int32 `form:"idservice"`
}
