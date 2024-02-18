package guests

import "database/sql"

type GuestRepository struct {
	db *sql.DB
}

func NewGuestRepository(db *sql.DB) *GuestRepository {
	return &GuestRepository{db: db}
}

func (r *GuestRepository) CreateGuest(guest *Guest) (insertedID string, err error) {
	lastInsertId := ""

	_ = r.db.QueryRow("INSERT INTO guests (id, user_id, name, email, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		guest.ID, guest.UserId, guest.Name, guest.Email, guest.CreatedAt).Scan(&lastInsertId)

	return lastInsertId, nil
}

func (r *GuestRepository) FindGuestById(id string) (*Guest, error) {
	var guest Guest
	err := r.db.QueryRow("SELECT id, user_id, name, email, created_at FROM guests WHERE id = $1", id).
		Scan(&guest.ID, &guest.UserId, &guest.Name, &guest.Email, &guest.CreatedAt)

	return &guest, err
}

func (r *GuestRepository) FindGuestByEmail(email string) (*Guest, error) {
	var guest Guest
	err := r.db.QueryRow("SELECT id, user_id, name, email, created_at FROM guests WHERE email = $1", email).
		Scan(&guest.ID, &guest.UserId, &guest.Name, &guest.Email, &guest.CreatedAt)

	return &guest, err
}

func (r *GuestRepository) UpdateGuest(guest *Guest) (rowsAffected int, err error) {
	rows, err := r.db.Exec("UPDATE guests SET name = $1, email = $2 WHERE id = $3", guest.Name, guest.Email, guest.ID)
	if err != nil {
		return 0, err
	}
	ra, _ := rows.RowsAffected()
	rowsAffected = int(ra)

	return rowsAffected, nil
}

func (r *GuestRepository) DeleteGuest(id string) (rowsAffected int, err error) {
	rows, err := r.db.Exec("DELETE FROM guests WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	ra, _ := rows.RowsAffected()
	rowsAffected = int(ra)
	return rowsAffected, nil
}
