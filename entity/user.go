package entity

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("名前は必須です")
	}
	if email == "" {
		return nil, errors.New("メールアドレスは必須です")
	}

	now := time.Now()
	return &User{
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (u *User) UpdateInfo(name, email string) error {
	if name == "" {
		return errors.New("名前は必須です")
	}
	if email == "" {
		return errors.New("メールアドレスは必須です")
	}

	u.Name = name
	u.Email = email
	u.UpdatedAt = time.Now()
	return nil
}