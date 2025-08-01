package usecase

import (
	"errors"
	"go-clean-architecture-playground/entity"
	"go-clean-architecture-playground/interface/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) CreateUser(name, email string) (*entity.User, error) {
	existingUser, _ := u.userRepo.GetByEmail(email)
	if existingUser != nil {
		return nil, errors.New("このメールアドレスは既に使用されています")
	}

	user, err := entity.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) GetUser(id int) (*entity.User, error) {
	if id <= 0 {
		return nil, errors.New("無効なユーザーIDです")
	}

	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("ユーザーが見つかりません")
	}

	return user, nil
}

func (u *UserUsecase) UpdateUser(id int, name, email string) (*entity.User, error) {
	user, err := u.GetUser(id)
	if err != nil {
		return nil, err
	}

	existingUser, _ := u.userRepo.GetByEmail(email)
	if existingUser != nil && existingUser.ID != id {
		return nil, errors.New("このメールアドレスは既に使用されています")
	}

	err = user.UpdateInfo(name, email)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) DeleteUser(id int) error {
	_, err := u.GetUser(id)
	if err != nil {
		return err
	}

	return u.userRepo.Delete(id)
}

func (u *UserUsecase) ListUsers() ([]*entity.User, error) {
	return u.userRepo.List()
}