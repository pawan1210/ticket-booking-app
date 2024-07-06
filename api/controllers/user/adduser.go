package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pawan1210/ticket-booking-app/internal/services/user/types"
)

func (controller *UserController) AddUser(request *gin.Context, requestBody *types.AddUserRequest) (any, error) {
	ctx := request.Request.Context()
	err := controller.service.AddUser(ctx, requestBody)

	if err != nil {
		return nil, err
	}

	return gin.H{}, nil
}
