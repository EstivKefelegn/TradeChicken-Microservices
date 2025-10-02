package models

import "time"

type Orders struct {
	ID        string    `json:"id,omitempty" db:"id,omitempty"`
	Product   string    `json:"product" db:"product"`
	User      string    `json:"user" db:"user"`
	Price     float64   `json:"price" db:"price"`
	CreatedAT time.Time `json:"created_at" db:"created_at"`
}

