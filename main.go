package main

import (
    "errors"
    "fmt"
    "github.com/urfave/cli"
    "net"
    "os"
    "strings"
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
    app.Action = func(c *cli.Context) error {
        firstArg := c.Args().First()

        if len(firstArg) == 0 {
            return errors.New("No IP given!")
        }

        ip := net.ParseIP(firstArg)

        fmt.Printf("Given: %q\n", ip)
        return nil
    }

    err := app.Run(os.Args)

    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }
}

func ExtractIpAddress(ipWithNetMask string) string {
    if len(ipWithNetMask) == 0 {
        return ""
    }

    position := strings.Index(ipWithNetMask, "/")

    if position > -1 {
        return ipWithNetMask[:position]
    }

    return ipWithNetMask
}

func ExtractNetMask(ipWithNetMask string) string {
    if len(ipWithNetMask) == 0 {
        return ""
    }

    position := strings.Index(ipWithNetMask, "/")

    if position > -1 {
        return ipWithNetMask[position + 1:]
    }

    return ""
}
