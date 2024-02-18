package user

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *User) (insertedID string, err error) {
	lastInsertId := ""

	_ = ur.db.QueryRow("INSERT INTO users (id, name, email, phone, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.ID, user.Name, user.Email, user.Phone, user.CreatedAt).Scan(&lastInsertId)

	_, err = ur.db.Exec("INSERT INTO passwords (user_id, password, code, created_at) VALUES ($1, $2, $3, $4)",
		user.ID, user.Password, 1, user.CreatedAt)

	return lastInsertId, err
}

func (ur *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := ur.db.QueryRow("SELECT id, name, email, phone, created_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt)
	return &user, err
}

func (ur *UserRepository) FindByEmailWithPassword(email string) (*User, error) {
	var user User

	err := ur.db.QueryRow(`
		SELECT u.id, u.name, u.email, u.phone, u.created_at, p.password 
		FROM "users" u 
		JOIN (
			SELECT DISTINCT ON ("user_id") "user_id", password
			FROM "passwords"
		    ORDER BY "user_id", "created_at" DESC) p ON u.id = p.user_id
		WHERE u.email = $1`,
		email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.Password)

	return &user, err
}

func (ur *UserRepository) FindById(id string) (*User, error) {
	var user User
	err := ur.db.QueryRow("SELECT id, name, email, phone, created_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.Password)
	return &user, err
}
