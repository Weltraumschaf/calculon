package main

import (
    "fmt"
    "log"
    "os"
    "github.com/urfave/cli"
    "strings"
)

var app = cli.NewApp()

// https://itnext.io/how-to-create-your-own-cli-with-golang-3c50727ac608
// http://jodies.de/ipcalc?host=10.0.0.0&mask1=16&mask2=
var pizza = []string{"Enjoy your pizza with some delicous"}

func info() {
    app.Name = "ipcalc"
    app.Usage = "Tool to calculate IP stuff"
    app.Author = "Sven Strittmatter"
    app.Version = "1.0.0"
}

func commands() {
    app.Commands = []cli.Command{
        {
            Name:    "peppers",
            Aliases: []string{"p"},
            Usage:   "Add peppers to your pizza",
            Action: func(c *cli.Context) {
                pe := "peppers"
                peppers := append(pizza, pe)
                m := strings.Join(peppers, " ")
                fmt.Println(m)
            },
        },
        {
            Name:    "pineapple",
            Aliases: []string{"pa"},
            Usage:   "Add pineapple to your pizza",
            Action: func(c *cli.Context) {
                pa := "pineapple"
                pineapple := append(pizza, pa)
                m := strings.Join(pineapple, " ")
                fmt.Println(m)
            },
        },
        {
            Name:    "cheese",
            Aliases: []string{"c"},
            Usage:   "Add cheese to your pizza",
            Action: func(c *cli.Context) {
                ch := "cheese"
                cheese := append(pizza, ch)
                m := strings.Join(cheese, " ")
                fmt.Println(m)
            },
        },
    }
}

func main() {
    info()
    commands()

    err := app.Run(os.Args)

    if err != nil {
        log.Fatal(err)
    }
}
