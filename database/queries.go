package database

import (
	"api-server/models"
	"fmt"
)

// func (d *DB) GetChefDetails(id string) (*models.User, error) {
func (d *DB) GetChef(id string) (*models.User, error) {

	fmt.Println("checf: ", id)

	users, err := d.GetChefs()
	if err != nil {
		return nil, err
	}

	allergies, err := d.GetAllergies(id)
	if err != nil {
		return nil, err
	}

	restrictions, err := d.GetDietaryRestrictions(id)
	if err != nil {
		return nil, err
	}

	favorites, err := d.GetFavoriteIngredients(id)
	if err != nil {
		return nil, err
	}

	var user *models.User

	// find user
	for _, u := range users {
		fmt.Println(u.ID)

		if u.ID == id {
			user = u
		}
	}

	// "join" other fields
	for _, a := range allergies {
		fmt.Println(a.ID, " : ", a.Allergy)
		if a.ID == id {
			user.Allergies = append(user.Allergies, a.Allergy)
		}
	}
	for _, r := range restrictions {
		fmt.Println(r.ID, " : ", r.Restriction)
		if r.ID == id {
			user.DietaryRestrictions = append(user.DietaryRestrictions, r.Restriction)
		}
	}
	for _, f := range favorites {
		fmt.Println(f.ID, " : ", f.Ingredient)
		if f.ID == id {
			user.FavoriteIngredients = append(user.FavoriteIngredients, f.Ingredient)
		}
	}
	user.CreatedAt = ""
	user.UpdatedAt = ""
	return user, nil
}

func (d *DB) GetChefs() ([]*models.User, error) {

	var u []*models.User
	query := "SELECT * FROM users"

	err := d.db.Select(&u, query)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// func (d *DB) GetChef(id string) (*models.User, error) {

// 	var u *models.User
// 	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s ", id)

// 	fmt.Println(query)

// 	err := d.db.Select(&u, query)
// 	if err != nil {
// 		fmt.Println(err)
// 		return u, err
// 	}

// 	return u, nil
// }

func (d *DB) GetAllergies(id string) ([]*models.Allergy, error) {

	var a []*models.Allergy
	// fmt.Println(id)
	// query := fmt.Sprintf("SELECT * FROM \"allergies\" WHERE user_id = %s", id)

	err := d.db.Select(&a, "SELECT * FROM allergies")
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (d *DB) GetDietaryRestrictions(id string) ([]*models.DietaryRestriction, error) {

	var dr []*models.DietaryRestriction
	query := fmt.Sprintf("SELECT * FROM dietary_restrictions")

	err := d.db.Select(&dr, query)
	if err != nil {
		return nil, err
	}

	return dr, nil
}

func (d *DB) GetFavoriteIngredients(id string) ([]*models.FavoriteIngredient, error) {

	var fi []*models.FavoriteIngredient
	query := fmt.Sprintf("SELECT * FROM favorite_ingredients")

	err := d.db.Select(&fi, query)
	if err != nil {
		return nil, err
	}

	return fi, nil
}
