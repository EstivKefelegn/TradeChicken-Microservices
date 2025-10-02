package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Password  string    `json:"password" db:"password"`
	Active    bool      `json:"active" db:"active"`
	CreatedAT time.Time `json:"created_at" db:"created_at"`
	UpdatedAT time.Time `json:"updated_at" db:"updated_at"`
}

type Models struct {
	User User
}

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool
	
	return Models{
		User: User{},
	}
}
