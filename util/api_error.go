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

// Error Generators {{{

func NewNotFoundError() *ApiError {
	return &ApiError{
		Status:           http.StatusNotFound,
		Code:             http.StatusNotFound,
		Message:          "Not Found",
		DeveloperMessage: "Perhaps you have an incorrect id?",
	}
}

func NewServerError() *ApiError {
	return &ApiError{
		Status:           http.StatusInternalServerError,
		Code:             http.StatusInternalServerError,
		Message:          http.StatusText(http.StatusInternalServerError),
		DeveloperMessage: "Server Error",
	}
}

func NewServerErrorWithError(err error) *ApiError {
	return &ApiError{
		Status:           http.StatusInternalServerError,
		Code:             http.StatusInternalServerError,
		Message:          http.StatusText(http.StatusInternalServerError),
		DeveloperMessage: fmt.Sprintf("%s", err),
	}
}

func NewInvalidMethodError() *ApiError {
	return &ApiError{
		Status:           http.StatusMethodNotAllowed,
		Code:             http.StatusMethodNotAllowed,
		Message:          http.StatusText(http.StatusMethodNotAllowed),
		DeveloperMessage: "Perhaps you meant to GET instead of POST? Or vice versa?",
	}
}

func NewUnauthorizedError() *ApiError {
	return &ApiError{
		Status:           http.StatusUnauthorized,
		Code:             http.StatusUnauthorized,
		Message:          http.StatusText(http.StatusUnauthorized),
		DeveloperMessage: "Check your key",
	}
}

func NewWebSocketFailedError() *ApiError {
	return &ApiError{
		Status:           http.StatusBadRequest,
		Code:             http.StatusBadRequest,
		Message:          http.StatusText(http.StatusBadRequest),
		DeveloperMessage: "We were unable to process your websocket request, perhaps it was not spec-valid?",
	}
}

func NewCustomError(status int, code int, msg string, dmsg string) *ApiError {
	return &ApiError{
		Status:           status,
		Code:             code,
		Message:          msg,
		DeveloperMessage: dmsg,
	}
}

// }}}

// Custom Error Writers {{{

func NotFound(w http.ResponseWriter) {
	WriteErrorResponse(w, NewNotFoundError())
}

func ServerError(w http.ResponseWriter, err error) {
	WriteErrorResponse(w, NewServerErrorWithError(err))
}

func InvalidMethod(w http.ResponseWriter) {
	WriteErrorResponse(w, NewInvalidMethodError())
}

func Unauthorized(w http.ResponseWriter) {
	WriteErrorResponse(w, NewUnauthorizedError())
}

func WebSocketFailed(w http.ResponseWriter) {
	WriteErrorResponse(w, NewWebSocketFailedError())
}

func CustomError(w http.ResponseWriter, status int, code int, message string, dMessage string) {
	apiError := NewCustomError(status, code, message, dMessage)
	WriteErrorResponse(w, apiError)
}

// }}}

func WriteErrorResponse(w http.ResponseWriter, apiError *ApiError) {
	WriteResourceResponse(w, apiError.Status, apiError)
}
