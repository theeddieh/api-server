package main

import (
	"api-server/app"
	"api-server/database"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("api-server starting up")
	fmt.Println("TEST_ENV: ", os.Getenv("TEST_ENV"))

	app := app.New()
	app.DB = database.New()

	err := app.DB.Open()
	if err != nil {
		log.Println(err)
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
