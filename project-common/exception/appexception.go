package appexception

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewAppError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func NotFoundError(message string) *AppError {
	return NewAppError(message, http.StatusNotFound)
}

func UnexpectedError(message string) *AppError {
	return NewAppError(message, http.StatusInternalServerError)
}

func ValidationError(message string) *AppError {
	return NewAppError(message, http.StatusUnprocessableEntity)
}

func AuthenticationError(message string) *AppError {
	return NewAppError(message, http.StatusUnauthorized)
}

func AuthorizationError(message string) *AppError {
	return NewAppError(message, http.StatusForbidden)
}
