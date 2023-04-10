package main

import (
	"api-server/app"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("api-server starting up")
	fmt.Println("TEST_ENV: ", os.Getenv("TEST_ENV"))

	app := app.New()

	fmt.Println("attaching router")
	http.HandleFunc("/", app.Router.ServeHTTP)

	fmt.Println("listening on port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
