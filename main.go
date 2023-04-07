package main

import (
        "fmt"
        "os"
)

func main() {
        fmt.Println("API server starting up...")
        fmt.Println(os.Getenv("TEST_ENV"))
}