package models

import "time"

type GetChefJSON struct {
	ID                  int64  `json:"id"`
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	Email               string `json:"email"`
	ExperienceLevel     string `json:"experienceLevel"`
	DietaryRestrictions string `json:"dietaryRestrictions"`
	Allergies           string `json:"allergies"`
	FavoriteIngredients string `json:"favoriteIngredients"`
}

type User struct {
	ID              string    `db:"id"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	Email           string    `db:"email"`
	ExperienceLevel string    `db:"experience"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type Allergies struct {
	ID      string `db:"user_id"`
	Allergy string `db:"allergy"`
}

type DietaryRestrictions struct {
	ID                 string `db:"user_id"`
	DietaryRestriction string `db:"dietary_restriction"`
}

type FavoriteIngredients struct {
	ID         string `db:"user_id"`
	Ingredient string `db:"ingredient"`
}
