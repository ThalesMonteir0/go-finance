package rest_err

import "net/http"

type RestErr struct {
	Code    int
	Message string
	Err     string
}

func (err RestErr) Error() string {
	return err.Message
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusBadRequest,
		Message: message,
		Err:     "bad_request",
	}
}

func NewInternalServeError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusInternalServerError,
		Message: message,
		Err:     "internal_server_error",
	}
}
