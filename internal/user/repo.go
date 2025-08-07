package user

import (
	"sync"

	"github.com/MalcolmFuchs/GoGuard/internal/common"
)

type Repository interface {
	GetByID(id string) (*User, error)
	ListRoles(userID string) ([]Role, error)
}

// InMemoryRepository is an in-memory implementation of the Repository interface.
type InMemoryRepository struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewInMemoryRepository(inital []*User) *InMemoryRepository {
	repo := &InMemoryRepository{
		users: make(map[string]*User),
	}
	for _, user := range inital {
		repo.users[user.ID] = user
	}
	return repo
}

func (r *InMemoryRepository) GetByID(id string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, common.ErrUserNotFound
	}
	return user, nil
}

func (r *InMemoryRepository) ListRoles(userID string) ([]Role, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[userID]
	if !exists {
		return nil, common.ErrUserNotFound
	}
	return user.Roles, nil
}
