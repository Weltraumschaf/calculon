package calculon

import (
    "errors"
    "fmt"
    "github.com/urfave/cli"
    "net"
)

func Create() *cli.App {
    var app = cli.NewApp()
    app.Name = "calculon"
    app.Usage = "Tool to calculate IP stuff"
    app.UsageText = app.Name + " 192.168.123.0/24"
    app.Author = "Sven Strittmatter"
    app.Email = "ich@weltraumschaf.de"
    app.Description = "TODO Description"
    app.Version = "1.0.0"
    app.Action = Execute
    return app
}

func Execute(c *cli.Context) error {
    firstArg := c.Args().First()

    if len(firstArg) == 0 {
        return errors.New("No IP given!")
    }

    ip, network, err := net.ParseCIDR(firstArg)

    if err != nil {
        return errors.New("Invalid netmask given!")
    }

    return printResult(ip, network)
}

func printResult(ip net.IP, network *net.IPNet) error {
    fmt.Printf("Address:   %s       %s\n", ip, FormatIpAsDottedBits(ip))

    //fmt.Printf("Netmask:   %s = %d  %s\n",
    //    decimalNetmask, netmask, binaryNetmask)
    fmt.Printf("Wildcard:  \n")
    fmt.Printf("Network:   \n")
    fmt.Printf("Broadcast: \n")
    fmt.Printf("HostMin:   \n")
    fmt.Printf("HostMax:   \n")
    fmt.Printf("Hosts/Net: \n")
    return nil
}
