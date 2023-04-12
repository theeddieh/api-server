package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/plain")
		fmt.Fprintln(w, "api-server root")
	}
}

func (a *App) HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := make(map[string]bool)
		m["alive"] = true

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(m)
		if err != nil {
			fmt.Println("unable to encode json:", m)
		}
	}
}

func (a *App) GetChefsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		chefs, err := a.DB.GetChefs()
		if err != nil {
			sendErrorResponse(w, r, err)
		} else {
			w.WriteHeader(http.StatusOK)
			e := json.NewEncoder(w).Encode(chefs)
			if err != nil {
				fmt.Println(e)
			}
		}
	}
}

func (a *App) GetChefHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		vars := mux.Vars(r)
		uuid := vars["uuid"]

		c, err := a.DB.GetChef(uuid)
		if err != nil {
			sendErrorResponse(w, r, err)
		} else {
			w.WriteHeader(http.StatusOK)
			e := json.NewEncoder(w).Encode(c)
			if err != nil {
				fmt.Println(e)
			}
		}
	}
}
func (a *App) GetAllergiesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		c, err := a.DB.GetAllergies()
		if err != nil {
			sendErrorResponse(w, r, err)
		} else {
			w.WriteHeader(http.StatusOK)
			e := json.NewEncoder(w).Encode(c)
			if err != nil {
				fmt.Println(e)
			}
		}
	}
}
func (a *App) GetDietaryRestrictionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		c, err := a.DB.GetDietaryRestrictions()
		if err != nil {
			sendErrorResponse(w, r, err)
		} else {
			w.WriteHeader(http.StatusOK)
			e := json.NewEncoder(w).Encode(c)
			if err != nil {
				fmt.Println(e)
			}
		}
	}
}
func (a *App) GetFavoriteIngredientsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		c, err := a.DB.GetFavoriteIngredients()
		if err != nil {
			sendErrorResponse(w, r, err)
		} else {
			w.WriteHeader(http.StatusOK)
			e := json.NewEncoder(w).Encode(c)
			if err != nil {
				fmt.Println(e)
			}
		}
	}
}
