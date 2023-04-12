package app

import (
	"api-server/database"
	"fmt"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     database.PostgresDB
}

func New() *App {
	fmt.Println("initializing app")

	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()

	return a
}

func (a *App) initRoutes() {
	fmt.Println("initializing handlers")

	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/healthcheck", a.HealthCheckHandler()).Methods("GET")

	a.Router.HandleFunc("/chefs", a.GetChefsHandler()).Methods("GET")
	a.Router.HandleFunc("/chef/{uuid}", a.GetChefHandler()).Methods("GET")

	a.Router.HandleFunc("/allergies", a.GetAllergiesHandler()).Methods("GET")
	a.Router.HandleFunc("/dietaryRestrictions", a.GetDietaryRestrictionsHandler()).Methods("GET")
	a.Router.HandleFunc("/favoriteIngredients", a.GetFavoriteIngredientsHandler()).Methods("GET")
}
