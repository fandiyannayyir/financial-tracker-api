package users

import (
	"database/sql"
	"financial-tracker-api/config"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DB: config.DB}
}

func (r *UserRepository) CreateUser(u User) error {
	query := `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)`
	
	_, err := r.DB.Exec(query, u.Name, u.Email, u.Password)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (User, error) {
	var u User

	query := `
		SELECT id, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
		LIMIT 1
	`
	err := r.DB.QueryRow(query, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	
	return u, err
}