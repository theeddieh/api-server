package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "jubilant-tribble/app"
)

func main() {
        fmt.Println("api-server starting up")
        fmt.Println(os.Getenv("TEST_ENV"))
        
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