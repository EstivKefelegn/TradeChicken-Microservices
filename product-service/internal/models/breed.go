package models

import "time"

type Breed struct {
	ID        string    `json:"id,omitempty" db:"id,omitempty"`
	Name      string    `json:"name" db:"name"`
	Origin    string    `json:"origin" db:"origin"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
