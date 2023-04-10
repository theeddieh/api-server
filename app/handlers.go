package app

import (
	"fmt"
	"net/http"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API server root")
	}
}

func (a *App) GetChefHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{firstName: Harry}")
	}
}

func (a *App) GetChefsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		chefs, err := a.DB.GetUsers()
		if err != nil {
			fmt.Println("err:", err)
		}

		fmt.Println("chef[0]:", chefs[0].FirstName)
	}
}
