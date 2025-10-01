package models

type Product struct {
	ID           string  `json:"id,omitempty" db:"id,omitempty"`
	AgeWeek      int     `json:"age_week" db:"age_week"`
	PricePerHen  float32 `json:"price_per_hen" db:"price_per_hen"`
	IsVaccinated bool    `json:"is_vaccinated" db:"is_vaccinated"`
	LastUpdate   string  `json:"last_update,omitempty" db:"last_update,omitempty"`
	BreedID      string	 `json:"breed_id" db:"breed_id"`	
	Breed        *Breed	 	
}
