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
    fmt.Println(
        PadFirstColumn("Address:"),
        PadSecondColumn(ip),
        FormatIpAsDottedBits(ip))
}

func printNetmask(mask net.IPMask) {
    netmask := FormatMaskAsDottedDecimal(mask)
    decimalNetmask, _ := mask.Size()
    fmt.Println(
        PadFirstColumn("Netmask:"),
        PadSecondColumn(fmt.Sprintf("%s = %d", netmask, decimalNetmask)),
        FormatMaskAsDottedBits(mask))
}

func printWildcard(mask net.IPMask) {
    wildcard := DeriveWildcard(mask)

    fmt.Println(
        PadFirstColumn("Wildcard:"),
        PadSecondColumn(wildcard), 
        FormatWildcardAsDottedBits(wildcard))
}

func printNetwork() {
    // TODO Implement print network.
    fmt.Printf("Network:   n/a\n")
}

func printBroadcast() {
    // TODO Implement print broadcast.
    fmt.Printf("Broadcast: n/a\n")
}

func printHostMin() {
    // TODO Implement print host min.
    fmt.Printf("HostMin:   n/a\n")
}

func printHostMax() {
    // TODO Implement print host max.
    fmt.Printf("HostMax:   n/a\n")
}

func printHostsPerNet() {
    // TODO Implement print hosts/net.
    fmt.Printf("Hosts/Net: n/a\n")
}
