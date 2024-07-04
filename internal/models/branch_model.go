package models

import (
	"database/sql"
	db "gitlab/go-prolog-api/example/db/sqlc"
	"time"
)

type CreateBranchRequest struct {
	Name        string    `form:"name"`
	Location    string    `form:"location"`
	Openingtime time.Time `form:"openingtime"`
	Closingtime time.Time `form:"closingtime"`
}

type BranchResponse struct {
	ID          int32        `json:"id"`
	Name        string       `json:"name"`
	Location    string       `json:"location"`
	Openingtime time.Time    `json:"openingtime"`
	Closingtime time.Time    `json:"closingtime"`
	Createdat   sql.NullTime `json:"createdat"`
	Updatedat   sql.NullTime `json:"updatedat"`
}

func NewBranchResponse(branch db.Branch) BranchResponse {
	return BranchResponse{
		ID:          branch.ID,
		Name:        branch.Name,
		Location:    branch.Location,
		Openingtime: branch.Openingtime,
		Closingtime: branch.Closingtime,
		Createdat:   branch.Createdat,
		Updatedat:   branch.Updatedat,
	}
}

func NewBranchsResponse(branch []db.Branch) []BranchResponse {
	response := make([]BranchResponse, 0)
	for _, branch := range branch {
		response = append(response, BranchResponse{
			ID:          branch.ID,
			Name:        branch.Name,
			Location:    branch.Location,
			Openingtime: branch.Openingtime,
			Createdat:   branch.Createdat,
			Updatedat:   branch.Updatedat,
		})
	}
	return response
}

type UpdateBranchsRequest struct {
	ID          int32        `form:"id"`
	Name        string       `form:"name"`
	Location    string       `form:"location"`
	Openingtime time.Time    `form:"openingtime"`
	Closingtime time.Time    `form:"closingtime"`
}
