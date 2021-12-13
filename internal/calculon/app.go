package calculon

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"net"
	color "weltraumschaf.de/calculon/internal/calculon/color"
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
		PadFirstColumn("Address:"),
		color.Blue(PadSecondColumn(ip.String())),
		color.Yellow(FormatIpAsDottedBits(ip)))
}

func printNetmask(mask net.IPMask) {
	netmask := FormatMaskAsDottedDecimal(mask)
	decimalNetmask, _ := mask.Size()
	fmt.Println(
		PadFirstColumn("Netmask:"),
		color.Blue(PadSecondColumn(fmt.Sprintf("%s = %d", netmask, decimalNetmask))),
		color.Red(FormatMaskAsDottedBits(mask)))
}

func printWildcard(mask net.IPMask) {
	wildcard := DeriveWildcard(mask)

	fmt.Println(
		PadFirstColumn("Wildcard:"),
		color.Blue(PadSecondColumn(wildcard.String())),
		color.Yellow(FormatWildcardAsDottedBits(wildcard)))
}

func printNetwork(network *net.IPNet) {
	decimalNetmask, _ := network.Mask.Size()
	fmt.Println(
		PadFirstColumn("Network:"),
		color.Blue(PadSecondColumn(fmt.Sprintf("%s/%d", network.IP, decimalNetmask))),
		color.Yellow(FormatIpAsDottedBits(network.IP)))
}

func printBroadcast(ip net.IP, network *net.IPNet) {
	broadcast := DeriveBroadCast(ip, network)
	fmt.Println(
		PadFirstColumn("Broadcast:"),
		color.Blue(PadSecondColumn(broadcast.String())),
		color.Yellow(FormatIpAsDottedBits(broadcast)))
}

func printHostMin(network *net.IPNet) {
	hostMin := HostMin(network)
	fmt.Println(
		PadFirstColumn("HostMin:"),
		color.Blue(PadSecondColumn(hostMin)),
		color.Yellow(FormatIpAsDottedBits(hostMin)))
}

func printHostMax(ip net.IP, network *net.IPNet) {
	hostMax := HostMax(ip, network)
	fmt.Println(
		PadFirstColumn("HostMax:"),
		color.Blue(PadSecondColumn(hostMax)),
		color.Yellow(FormatIpAsDottedBits(hostMax)))
}

func printHostsPerNet(ip net.IP, network *net.IPNet) {
	number := HostsPerNetwork(ip, network)
	fmt.Print(
		PadFirstColumn("Hosts/Net:"),
		color.Blue(PadSecondColumn(fmt.Sprintf("%d", number))))
}
