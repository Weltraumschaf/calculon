package main

import (
    "fmt"
    "os"
    "weltraumschaf.de/calculon/internal/calculon"
)

func main() {
    var app = calculon.Create()
    err := app.Run(os.Args)

    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }
}

