package models

import "fmt"

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: err.Error(),
	}
}

func ResponseUnit(msg string, err error) error {
	return fmt.Errorf("%s: %v", msg, err.Error())
}

func ResponseUnitString(where string, err string) error {
	return fmt.Errorf("%v with %s not found", where, err)
}
