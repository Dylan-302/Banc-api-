package dao

import (
	"fmt"
	"sync"

	"banc-api/src/modules/user/domain/entities"
)

// InMemoryUserRepository provides a simple in-memory implementation for local development.
type InMemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[uint]entities.User
	nextID uint
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[uint]entities.User)}
}

func (r *InMemoryUserRepository) Create(user entities.User) (*entities.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.nextID++
	user.ID = r.nextID
	r.users[user.ID] = user
	return &user, nil
}

func (r *InMemoryUserRepository) FindByID(id uint) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %d not found", id)
	}

	return &user, nil
}

func (r *InMemoryUserRepository) FindAll() ([]entities.User, error) {
	return r.GetAll()
}

func (r *InMemoryUserRepository) GetAll() ([]entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]entities.User, 0, len(r.users))
	for _, user := range r.users {
		result = append(result, user)
	}
	return result, nil
}

func (r *InMemoryUserRepository) Update(user entities.User) (*entities.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return nil, fmt.Errorf("user with id %d not found", user.ID)
	}

	r.users[user.ID] = user
	return &user, nil
}

func (r *InMemoryUserRepository) Delete(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return fmt.Errorf("user with id %d not found", id)
	}

	delete(r.users, id)
	return nil
}
