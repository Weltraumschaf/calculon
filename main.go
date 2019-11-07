package main

import (
    "fmt"
    "log"
    "os"
    "github.com/urfave/cli"
    "errors"
)

var app = cli.NewApp()

func main() {
    app.Name = "calculon"
    app.Usage = "Tool to calculate IP stuff"
    app.Author = "Sven Strittmatter"
    app.Email = "ich@weltraumschaf.de"
    app.Description = "TODO Description"
    app.Version = "1.0.0"
    app.Action = func(c *cli.Context) error {
        firstArg := c.Args().First();

        if len(firstArg) == 0 {
            return errors.New("no IP given")
        }

        fmt.Printf("Given: %q\n", firstArg)
        return nil
    }
    
    err := app.Run(os.Args)

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}
