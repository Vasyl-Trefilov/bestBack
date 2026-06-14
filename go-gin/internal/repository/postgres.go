<<<<<<< HEAD
package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(configDB string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", configDB)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL database successfully!")

	return db, nil
}
=======
package postgres
>>>>>>> bd3a36c591bd590609f477457d6e266c02a6b81c
