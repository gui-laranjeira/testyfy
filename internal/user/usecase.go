package user

import (
	"database/sql"
	"errors"
)

type UserUseCase struct {
	repo IUserRepository
}

func NewUserUseCase(repo IUserRepository) *UserUseCase {
	return &UserUseCase{repo}
}

func (uc *UserUseCase) CreateUser(input CreateUserInput) (*User, error) {
	_, err := uc.repo.FindByEmail(input.Email)
	if err != sql.ErrNoRows {
		return nil, errors.New("email already in use")
	}

	user, err := NewUser(input.Email, input.Password, input.Name, input.Phone)
	if err != nil {
		return nil, err
	}
	_, err = uc.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) Authenticate(input AuthenticateInput) (*User, error) {
	user, err := uc.repo.FindByEmailWithPassword(input.Email)
	if err != nil {
		return nil, err
	}
	if !user.ComparePassword(input.Password) {
		return nil, errors.New("invalid password")
	}
	return user, nil
}
