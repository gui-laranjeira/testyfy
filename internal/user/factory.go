package user

import "database/sql"

func UserFactory(db *sql.DB) *UserHandler {
	repo := NewUserRepository(db)
	useCase := NewUserUseCase(repo)
	return NewUserHandler(useCase)
}
