package app

import (
	"fmt"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
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
	a.Router.HandleFunc("/chef", a.GetChefHandler()).Methods("GET")
}
