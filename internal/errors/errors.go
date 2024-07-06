package errors

import (
	"net/http"
	"reflect"
)

func GetStatusCode(err error) int {
	errorName := reflect.TypeOf(err).Elem().Name()
	statusCode, exists := ErrorNameToHTTPStatusCode[errorName]

	if exists {
		return statusCode
	}

	return http.StatusInternalServerError
}
