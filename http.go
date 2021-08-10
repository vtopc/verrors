package verrors

import (
	"net/http"
)

// ToHTTPStatusCode maps err into HTTP status code
func ToHTTPStatusCode(err error) int {
	switch {
	case IsInvalidArgument(err):
		return http.StatusBadRequest
	case IsNotFound(err):
		return http.StatusNotFound
	case IsAlreadyExists(err):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
