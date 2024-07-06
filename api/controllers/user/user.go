package user

import "github.com/pawan1210/ticket-booking-app/internal/services/user"

type UserController struct {
	service *user.UserService
}

func New(service *user.UserService) *UserController {
	return &UserController{
		service,
	}
}
