package user

import (
	"github.com/pawan1210/ticket-booking-app/internal/stores"
)

type UserService struct {
	store stores.UserStore
}

func New(store stores.UserStore) *UserService {
	return &UserService{
		store,
	}
}
