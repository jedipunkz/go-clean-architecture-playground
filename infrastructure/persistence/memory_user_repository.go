package persistence

import (
	"errors"
	"go-clean-architecture-playground/entity"
	"go-clean-architecture-playground/interface/repository"
	"sync"
)

type MemoryUserRepository struct {
	users  map[int]*entity.User
	lastID int
	mutex  sync.RWMutex
}

func NewMemoryUserRepository() repository.UserRepository {
	return &MemoryUserRepository{
		users:  make(map[int]*entity.User),
		lastID: 0,
	}
}

func (r *MemoryUserRepository) Create(user *entity.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.lastID++
	user.ID = r.lastID
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) GetByID(id int) (*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("ユーザーが見つかりません")
	}
	return user, nil
}

func (r *MemoryUserRepository) GetByEmail(email string) (*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (r *MemoryUserRepository) Update(user *entity.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("ユーザーが見つかりません")
	}

	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("ユーザーが見つかりません")
	}

	delete(r.users, id)
	return nil
}

func (r *MemoryUserRepository) List() ([]*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*entity.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}