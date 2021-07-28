package errors

import "net/http"

type RestErrors struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewInternalServerError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal_server_error",
	}
}

func NewNotFoundError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not_found",
	}
}
