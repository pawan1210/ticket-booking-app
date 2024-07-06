package stores

import "github.com/pawan1210/ticket-booking-app/internal/entities"

type UserStore interface {
	Find(userID string) (*entities.User, error)
	Create(user *entities.User) (string, error)
}
