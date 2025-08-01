package repository

import "go-clean-architecture-playground/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id int) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int) error
	List() ([]*entity.User, error)
}