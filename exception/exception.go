package exception

import (
	"fmt"
	"net/http"
)

func NewNotFound(message string, a ...any) *ApiException {
	return New(404, message, a...)
}

func IsNotFound(err error) bool {

	if e, ok := err.(*ApiException); ok {
		if e.Code == 404 {
			return true
		}
	}
	return false
}

func NewAuthFailed(message string, a ...any) *ApiException {
	return New(5000, message, a...)

}

func NewPermissionDenied(message string, a ...any) *ApiException {
	return New(6000, message, a...)
}

func NewTokenExpired(message string, a ...any) *ApiException {

	return New(5001, message, a...)

}
func New(code int, message string, a ...any) *ApiException {
	httpCode := http.StatusInternalServerError
	if code > 0 && code < 600 {
		httpCode = code
	}
	return &ApiException{
		Code:     code,
		Message:  fmt.Sprintf(message, a...),
		HttpCode: httpCode,
	}
}

type ApiException struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
	HttpCode int    `json:"http_code"`
}

func (e *ApiException) Error() string {
	return e.Message
}
