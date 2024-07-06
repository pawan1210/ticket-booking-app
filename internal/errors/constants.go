package errors

import "net/http"

var ErrorNameToHTTPStatusCode = map[string]int{
	"BadRequest": http.StatusBadRequest,
}
