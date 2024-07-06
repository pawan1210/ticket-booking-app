package utils

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/pawan1210/ticket-booking-app/internal/errors"
)

func parseRequest[Req any](ctx *gin.Context, req *Req) error {
	if req == nil {
		return nil
	}

	var err error

	err = ctx.ShouldBindJSON(req)

	if err != nil && err.Error() != io.EOF.Error() {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}
