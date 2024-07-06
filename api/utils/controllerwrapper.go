package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	customerrors "github.com/pawan1210/ticket-booking-app/internal/errors"
)

type controllerFunctionv2[Req any, Res any] func(ctx *gin.Context, reqType *Req) (Res, error)

type apiResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error,omitempty"`
	Data    any   `json:"data,omitempty"`
}

func ControllerWrapper[Req any, Res any](reqType Req, fn controllerFunctionv2[Req, Res]) func(*gin.Context) {
	return func(request *gin.Context) {
		// ctx := request.Request.Context()
		var response *apiResponse
		var statusCode int

		defer handlePanic(request)

		err := parseRequest(request, &reqType)

		if err != nil {
			response, statusCode := CreateApiResponse(nil, err)

			request.JSON(statusCode, response)

			request.Abort()

			return
		}

		data, err := fn(request, &reqType)

		if err != nil {
			response, statusCode = CreateApiResponse(nil, err)
		} else {
			response, statusCode = CreateApiResponse(data, nil)
		}

		request.JSON(statusCode, response)
	}
}

func handlePanic(request *gin.Context) {
	if r := recover(); r != nil {
		response, statusCode := CreateApiResponse(nil, errors.New("panic occured"))

		request.JSON(statusCode, response)

		request.Abort()

		return
	}
}

func CreateApiResponse(data any, err error) (*apiResponse, int) {
	if data != nil {
		return &apiResponse{
			Success: true,
			Data:    data,
		}, http.StatusOK
	}

	statusCode := customerrors.GetStatusCode(err)

	return &apiResponse{
		Success: false,
		Error:   err,
	}, statusCode
}
