package main

import (
    "fmt"
    "github.com/urfave/cli"
    "os"
)

var app = cli.NewApp()

func main() {
    app.Name = "calculon"
    app.Usage = "Tool to calculate IP stuff"
    app.UsageText = app.Name + " 192.168.123.0/24"
    app.Author = "Sven Strittmatter"
    app.Email = "ich@weltraumschaf.de"
    app.Description = "TODO Description"
    app.Version = "1.0.0"
    app.Action = Action

    err := app.Run(os.Args)

    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }
}

func Action(c *cli.Context) error {
    fmt.Println("Hello, World!")
    return nil
}
