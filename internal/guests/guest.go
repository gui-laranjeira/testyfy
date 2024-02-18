package guests

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Guest struct {
	ID        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type IGuestRepository interface {
	CreateGuest(guest *Guest) (insertedID string, err error)
	FindGuestById(id string) (*Guest, error)
	FindGuestByEmail(email string) (*Guest, error)
	UpdateGuest(guest *Guest) (rowsAffected int, err error)
	DeleteGuest(id string) (rowsAffected int, err error)
}

type IGuestUseCase interface {
	CreateGuest(CreateGuestInput) (*Guest, error)
	FindGuestById(id string) (*Guest, error)
	FindGuestByEmail(email string) (*Guest, error)
	UpdateGuest(UpdateGuestInput) (*Guest, error)
	DeleteGuest(id string) error
}

type CreateGuestInput struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UpdateGuestInput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewGuest(userId string, name string, email string) (*Guest, error) {
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	guest := &Guest{
		ID:        uuid.New(),
		UserId:    userUUID,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
	validGuest := guest.validateGuest()
	if !validGuest {
		return nil, errors.New("invalid guest")
	}
	return guest, nil
}

func (u *Guest) validateGuest() bool {
	if u.Name == "" || u.Email == "" || u.ID == uuid.Nil || u.CreatedAt.IsZero() {
		return false
	}
	if !validateEmail(u.Email) {
		return false
	}
	return true
}

func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
