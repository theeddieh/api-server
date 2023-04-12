package database

import (
	"api-server/models"
	"fmt"
)

func (d *DB) GetChef(id string) (*models.User, error) {

	fmt.Println("checf: ", id)

	users, err := d.GetChefs()
	if err != nil {
		return nil, err
	}

	allergies, err := d.GetAllergies()
	if err != nil {
		return nil, err
	}

	restrictions, err := d.GetDietaryRestrictions()
	if err != nil {
		return nil, err
	}

	favorites, err := d.GetFavoriteIngredients()
	if err != nil {
		return nil, err
	}

	var user *models.User

	// find user
	for _, u := range users {
		if u.ID == id {
			user = u
		}
	}

	// "join" other fields
	for _, a := range allergies {
		if a.ID == id {
			user.Allergies = append(user.Allergies, a.Allergy)
		}
	}
	for _, r := range restrictions {
		if r.ID == id {
			user.DietaryRestrictions = append(user.DietaryRestrictions, r.Restriction)
		}
	}
	for _, f := range favorites {
		if f.ID == id {
			user.FavoriteIngredients = append(user.FavoriteIngredients, f.Ingredient)
		}
	}

	// zero out to hide from JSON response
	user.CreatedAt = ""
	user.UpdatedAt = ""

	return user, nil
}

func (d *DB) GetChefs() ([]*models.User, error) {

	var u []*models.User

	err := d.db.Select(&u, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (d *DB) GetAllergies() ([]*models.Allergy, error) {

	var a []*models.Allergy

	err := d.db.Select(&a, "SELECT * FROM allergies")
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (d *DB) GetDietaryRestrictions() ([]*models.DietaryRestriction, error) {

	var dr []*models.DietaryRestriction

	err := d.db.Select(&dr, "SELECT * FROM dietary_restrictions")
	if err != nil {
		return nil, err
	}

	return dr, nil
}

func (d *DB) GetFavoriteIngredients() ([]*models.FavoriteIngredient, error) {

	var fi []*models.FavoriteIngredient

	err := d.db.Select(&fi, "SELECT * FROM favorite_ingredients")
	if err != nil {
		return nil, err
	}

	return fi, nil
}
