package models

type User struct {
	ID                  string   `db:"id" json:"id"`
	FirstName           string   `db:"first_name" json:"firstName"`
	LastName            string   `db:"last_name" json:"lastName"`
	Email               string   `db:"email" json:"email"`
	ExperienceLevel     string   `db:"experience" json:"experienceLevel"`
	Allergies           []string `json:"dietaryRestrictions,omitempty"`
	DietaryRestrictions []string `json:"allergies,omitempty"`
	FavoriteIngredients []string `json:"favoriteIngredients,omitempty"`
	CreatedAt           string   `db:"created_at" json:",omitempty"`
	UpdatedAt           string   `db:"updated_at" json:",omitempty"`
	// CreatedAt           time.Time `db:"created_at" json:",omitempty"`
	// UpdatedAt           time.Time `db:"updated_at" json:",omitempty"`
}

type Allergy struct {
	ID      string `db:"user_id" json:"id"`
	Allergy string `db:"allergy" json:"allergy"`
}

type DietaryRestriction struct {
	ID          string `db:"user_id" json:"id"`
	Restriction string `db:"dietary_restriction" json:"dietaryRestriction"`
}

type FavoriteIngredient struct {
	ID         string `db:"user_id" json:"id"`
	Ingredient string `db:"ingredient" json:"favoriteIngredient"`
}
