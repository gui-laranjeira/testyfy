package guests

import (
	"database/sql"
	"errors"

	"github.com/gui-laranjeira/testyfy/internal/user"
)

type GuestUseCase struct {
	guestRepo IGuestRepository
	userRepo  user.IUserRepository
}

func NewGuestUseCase(guestRepo IGuestRepository, userRepo user.IUserRepository) *GuestUseCase {
	return &GuestUseCase{guestRepo: guestRepo, userRepo: userRepo}
}

func (uc *GuestUseCase) CreateGuest(input CreateGuestInput) (*Guest, error) {
	_, err := uc.userRepo.FindById(input.UserId)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	_, err = uc.guestRepo.FindGuestByEmail(input.Email)
	if err != sql.ErrNoRows {
		return nil, errors.New("email already in use")
	}

	guest, err := NewGuest(input.UserId, input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	_, err = uc.guestRepo.CreateGuest(guest)
	if err != nil {
		return nil, err
	}

	return guest, nil
}

func (uc *GuestUseCase) FindGuestById(id string) (*Guest, error) {
	return uc.guestRepo.FindGuestById(id)
}

func (uc *GuestUseCase) FindGuestByUserId(userId string) ([]*Guest, error) {
	return uc.guestRepo.FindGuestByUserId(userId)
}

func (uc *GuestUseCase) FindGuestByEmail(email string) (*Guest, error) {
	return uc.guestRepo.FindGuestByEmail(email)
}

func (uc *GuestUseCase) UpdateGuest(input UpdateGuestInput) (*Guest, error) {
	guest, err := uc.guestRepo.FindGuestById(input.ID)
	if err != nil {
		return nil, err
	}

	guest.Name = input.Name
	guest.Email = input.Email

	_, err = uc.guestRepo.UpdateGuest(guest)
	if err != nil {
		return nil, err
	}

	return guest, nil
}

func (uc *GuestUseCase) DeleteGuest(id string) (int, error) {
	rowsAffected, err := uc.guestRepo.DeleteGuest(id)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
