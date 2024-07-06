package user

import (
	"context"

	"github.com/pawan1210/ticket-booking-app/internal/entities"
	"github.com/pawan1210/ticket-booking-app/internal/services/user/types"
)

func (service *UserService) AddUser(ctx context.Context, requestBody *types.AddUserRequest) error {
	newUser := &entities.User{}

	newUser.SetName(requestBody.Name)

	_, err := service.store.Create(newUser)

	if err != nil {
		return err
	}

	return nil
}
