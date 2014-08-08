package models

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Status           int    `json:"status"`
	Code             int    `json:"code"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
}

func (e *ApiError) ToJson() ([]byte, error) {
	// Always pretty-print JSON
	return json.MarshalIndent(*e, "", "    ")
}

func ErrorResponse(w http.ResponseWriter, status int, code int, message string, dMessage string) {
	w.WriteHeader(status)

	apiError := ApiError{
		Status:           status,
		Code:             code,
		Message:          message,
		DeveloperMessage: dMessage,
	}

	bytes, _ := apiError.ToJson()

	w.Write(bytes)
}
