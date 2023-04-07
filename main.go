package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "jubilant-tribble/app"
)

func main() {
        fmt.Println("API server starting up")
        fmt.Println(os.Getenv("TEST_ENV"))
        
        app.New()
        
        fmt.Println("API listening...")
        err := http.ListenAndServe(":9000", nil)
        if err != nil {
                log.Println(err)
                os.Exit(1)
        }
}