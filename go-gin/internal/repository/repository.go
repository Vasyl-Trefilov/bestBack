package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(name, email string) (int, error) {
	var userId int
	err := r.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", name, email).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
