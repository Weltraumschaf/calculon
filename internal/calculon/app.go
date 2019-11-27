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

    printResult(ip, network)
    return nil
}

func printResult(ip net.IP, network *net.IPNet) {
    printAddress(ip)
    printNetmask(network.Mask)
    printWildcard(network.Mask)
    printNetwork()
    printBroadcast()
    printHostMin()
    printHostMax()
    printHostsPerNet()
}

func printAddress(ip net.IP) {
    fmt.Printf("Address:   %s       %s\n", ip, FormatIpAsDottedBits(ip))
}

func printNetmask(mask net.IPMask) {
    netmask := FormatMaskAsDottedDecimal(mask)
    decimalNetmask, _ := mask.Size()
    binaryNetmask := FormatMaskAsDottedBits(mask)
    fmt.Printf("Netmask:   %s = %d  %s\n", netmask, decimalNetmask, binaryNetmask)
}

func printWildcard(mask net.IPMask) {
    wildcard := DeriveWildcard(mask)
    decimalWildcard := FormatWildcardAsDottedDecimal(wildcard)
    binaryWildcard := FormatWildcardAsDottedBits(wildcard)
    fmt.Printf("Wildcard:  %s       %s\n", decimalWildcard, binaryWildcard)
}

func printNetwork() {
    fmt.Printf("Network:   \n")
}

func printBroadcast() {
    fmt.Printf("Broadcast: \n")
}

func printHostMin() {
    fmt.Printf("HostMin:   \n")
}

func printHostMax() {
    fmt.Printf("HostMax:   \n")
}

func printHostsPerNet() {
    fmt.Printf("Hosts/Net: \n")
}
