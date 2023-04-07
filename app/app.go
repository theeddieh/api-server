package app

import (
	"fmt"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	fmt.Println("Initializing app")
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	fmt.Println("Initializing handler for route: /")
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
}
