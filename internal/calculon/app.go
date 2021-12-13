package calculon

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"net"
	"weltraumschaf.de/calculon/internal/calculon/calc"
	"weltraumschaf.de/calculon/internal/calculon/color"
	"weltraumschaf.de/calculon/internal/calculon/format"
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

	printResult(ip.To4(), network)
	return nil
}

func printResult(ip net.IP, network *net.IPNet) {
	printAddress(ip)
	printNetmask(network.Mask)
	printWildcard(network.Mask)
	fmt.Println("=>")
	printNetwork(network)
	printHostMin(network)
	printHostMax(ip, network)
	printBroadcast(ip, network)
	printHostsPerNet(ip, network)

	if ip.IsPrivate() {
		fmt.Println("Private Internet")
	}
}

func printAddress(ip net.IP) {
	fmt.Println(
		format.PadFirstColumn("Address:"),
		color.Blue(format.PadSecondColumn(ip.String())),
		color.Yellow(format.IpAsDottedBits(ip)))
}

func printNetmask(mask net.IPMask) {
	netmask := format.MaskAsDottedDecimal(mask)
	decimalNetmask, _ := mask.Size()
	fmt.Println(
		format.PadFirstColumn("Netmask:"),
		color.Blue(format.PadSecondColumn(fmt.Sprintf("%s = %d", netmask, decimalNetmask))),
		color.Red(format.MaskAsDottedBits(mask)))
}

func printWildcard(mask net.IPMask) {
	wildcard := calc.DeriveWildcard(mask)

	fmt.Println(
		format.PadFirstColumn("Wildcard:"),
		color.Blue(format.PadSecondColumn(wildcard.String())),
		color.Yellow(format.WildcardAsDottedBits(wildcard)))
}

func printNetwork(network *net.IPNet) {
	decimalNetmask, _ := network.Mask.Size()
	fmt.Println(
		format.PadFirstColumn("Network:"),
		color.Blue(format.PadSecondColumn(fmt.Sprintf("%s/%d", network.IP, decimalNetmask))),
		color.Yellow(format.IpAsDottedBits(network.IP)))
}

func printBroadcast(ip net.IP, network *net.IPNet) {
	broadcast := calc.DeriveBroadCast(ip, network)
	fmt.Println(
		format.PadFirstColumn("Broadcast:"),
		color.Blue(format.PadSecondColumn(broadcast.String())),
		color.Yellow(format.IpAsDottedBits(broadcast)))
}

func printHostMin(network *net.IPNet) {
	hostMin := calc.HostMin(network)
	fmt.Println(
		format.PadFirstColumn("HostMin:"),
		color.Blue(format.PadSecondColumn(hostMin)),
		color.Yellow(format.IpAsDottedBits(hostMin)))
}

func printHostMax(ip net.IP, network *net.IPNet) {
	hostMax := calc.HostMax(ip, network)
	fmt.Println(
		format.PadFirstColumn("HostMax:"),
		color.Blue(format.PadSecondColumn(hostMax)),
		color.Yellow(format.IpAsDottedBits(hostMax)))
}

func printHostsPerNet(ip net.IP, network *net.IPNet) {
	number := calc.HostsPerNetwork(ip, network)
	fmt.Print(
		format.PadFirstColumn("Hosts/Net:"),
		color.Blue(format.PadSecondColumn(fmt.Sprintf("%d", number))))
}
