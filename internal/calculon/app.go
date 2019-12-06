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
    app.Description = "Simple IPv4 calculator."
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
    printNetwork(network)
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

func printNetwork(network *net.IPNet) {
    // TODO Implement print network.
    fmt.Println(
        PadFirstColumn("Network:"),
        "n/a")
}

func printBroadcast() {
    // TODO Implement print broadcast.
    fmt.Println(
        PadFirstColumn("Broadcast:"),
        "n/a")
}

func printHostMin() {
    // TODO Implement print host min.
    fmt.Println(
        PadFirstColumn("Host min.:"),
        "n/a")
}

func printHostMax() {
    // TODO Implement print host max.
    fmt.Println(
        PadFirstColumn("Host max.:"),
        "n/a")
}

func printHostsPerNet() {
    // TODO Implement print hosts/net.
    fmt.Println(
        PadFirstColumn("Hosts/Net:"),
        "n/a")
}
