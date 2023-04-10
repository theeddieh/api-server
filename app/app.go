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
	a.Router.HandleFunc("/chefs", a.GetChefsHandler()).Methods("GET")
	a.Router.HandleFunc("/chef", a.GetChefHandler()).Methods("GET")
}
