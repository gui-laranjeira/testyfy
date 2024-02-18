package user

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	Password  string    `json:"password"`
}

type IUserRepository interface {
	CreateUser(user *User) (insertedID string, err error)
	FindByEmail(email string) (*User, error)
	FindByEmailWithPassword(email string) (*User, error)
	FindById(id string) (*User, error)
}

type IUserUseCase interface {
	CreateUser(CreateUserInput) (*User, error)
	Authenticate(AuthenticateInput) (*User, error)
}

type CreateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

type AuthenticateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(email string, password string, name string, phone string) (*User, error) {
	if validPassword := validatePassword(password); !validPassword {
		return nil, errors.New("password must have at least 8 characters, one upper case letter, one lower case letter, a number and a symbol")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now(),
		Password:  string(hashedPassword),
	}
	if validUser := user.validateUser(); !validUser {
		return nil, errors.New("invalid user entity")
	}

	return &user, nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) validateUser() bool {
	if u.Name == "" || u.Email == "" || u.Phone == "" || u.ID == uuid.Nil || u.CreatedAt.IsZero() {
		return false
	}
	if !validateEmail(u.Email) {
		return false
	}
	return true
}

func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasNumber, _ := regexp.MatchString(`[0-9]`, password)
	if !hasNumber {
		return false
	}
	hasUppercase, _ := regexp.MatchString(`[A-Z]`, password)
	if !hasUppercase {
		return false
	}
	hasLowercase, _ := regexp.MatchString(`[a-z]`, password)
	if !hasLowercase {
		return false
	}
	hasSymbol, _ := regexp.MatchString(`[^a-zA-Z0-9]`, password)
	return hasSymbol
}

func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
