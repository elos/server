package util

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	Status           int    `json:"status"`
	Code             int    `json:"code"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
}

func ErrorResponse(w http.ResponseWriter, status int, code int, message string, dMessage string) {
	apiError := ApiError{
		Status:           status,
		Code:             code,
		Message:          message,
		DeveloperMessage: dMessage,
	}

	ResourceResponse(w, status, apiError)
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
