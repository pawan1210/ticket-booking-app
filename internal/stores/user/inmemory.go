package user

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/pawan1210/ticket-booking-app/internal/entities"
)

type inMemoryStore struct {
	lock  sync.RWMutex
	users map[string]entities.User
}

func NewInMemoryStore() *inMemoryStore {
	return &inMemoryStore{
		users: make(map[string]entities.User),
		lock:  sync.RWMutex{},
	}
}

func (store *inMemoryStore) Create(user *entities.User) (string, error) {
	store.lock.Lock()
	defer store.lock.Unlock()

	id := uuid.New().String()

	user.SetID(id)

	store.users[id] = *user

	return id, nil
}

func (store *inMemoryStore) Find(userID string) (*entities.User, error) {
	store.lock.RLock()
	defer store.lock.RUnlock()

	user, exists := store.users[userID]

	if !exists {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
