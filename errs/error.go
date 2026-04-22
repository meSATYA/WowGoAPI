package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func CustomerNotFound(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: "Customer does not exist"}
}

func CustomUnexpectedError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: "Unexpected Database Error"}
}
