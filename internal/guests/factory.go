package guests

import (
	"database/sql"

	"github.com/gui-laranjeira/testyfy/internal/user"
)

func GuestFactory(db *sql.DB) *GuestHandler {
	guestRepo := NewGuestRepository(db)
	userRepo := user.NewUserRepository(db)
	useCase := NewGuestUseCase(guestRepo, userRepo)
	return NewGuestHandler(useCase)
}
