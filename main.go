package main

import (
	"api-server/app"
	"api-server/database"
	"fmt"
	"log"
	"net/http"
	"os"
)

var version string

func main() {
	fmt.Println("api-server starting up")
	fmt.Println("version: ", version)

	app := app.New()
	app.DB = database.New()

	err := app.DB.Open()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		defer app.DB.Close()
	}

	fmt.Println("attaching router")
	http.HandleFunc("/", app.Router.ServeHTTP)

	fmt.Println("listening on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
