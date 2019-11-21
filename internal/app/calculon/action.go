package calculon

import (
    "fmt"
    "github.com/urfave/cli"
)

func Execute(c *cli.Context) error {
    fmt.Println("Hello, World!")
    return nil
}
