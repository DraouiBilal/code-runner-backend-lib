package main

import (
    "fmt"
    "code-runner/backend/lib/routing"
)

func main () {
    server := routing.CreateServer(&routing.Options{
        Host: "0.0.0.0",
        Port: 8080,
    })

    err:=server.ListenAndServe()

    if err != nil {
        fmt.Println(err)
    } 

    fmt.Println("Test")
}
