package util

import (
	"encoding/json"
	"fmt"
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

func NotFound(w http.ResponseWriter) {
	ErrorResponse(w, 404, 404, "Not Found", "Perhaps you have an incorrect id?")
}

func ServerError(w http.ResponseWriter, err error) {
	ErrorResponse(w, 500, 500, "Server Error", fmt.Sprintf("%s", err))
}

func InvalidMethod(w http.ResponseWriter) {
	ErrorResponse(w, 405, 405, "Invalid Method", "Perhaps you meant to GET instead of POST? Or vice versa?")
}
