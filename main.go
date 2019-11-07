package main

import (
    "fmt"
    "log"
    "os"
    "github.com/urfave/cli"
)

var app = cli.NewApp()

// https://itnext.io/how-to-create-your-own-cli-with-golang-3c50727ac608
// http://jodies.de/ipcalc?host=10.0.0.0&mask1=16&mask2=

func main() {
    app.Name = "calculon"
    app.Usage = "Tool to calculate IP stuff"
    app.Author = "Sven Strittmatter"
    app.Email = "ich@weltraumschaf.de"
    app.Description = "TODO Description"
    app.Version = "1.0.0"
    app.Action = func(c *cli.Context) error {
        fmt.Println("boom! I say!")
        return nil
    }
    
    err := app.Run(os.Args)

    if err != nil {
        log.Fatal(err)
    }
}
