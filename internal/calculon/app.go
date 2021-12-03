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
}

func printAddress(ip net.IP) {
	fmt.Println(
		PadFirstColumn("Address:"),
		Blue(PadSecondColumn(ip.String())),
		Yellow(FormatIpAsDottedBits(ip)))
}

func printNetmask(mask net.IPMask) {
	netmask := FormatMaskAsDottedDecimal(mask)
	decimalNetmask, _ := mask.Size()
	fmt.Println(
		PadFirstColumn("Netmask:"),
		Blue(PadSecondColumn(fmt.Sprintf("%s = %d", netmask, decimalNetmask))),
		Red(FormatMaskAsDottedBits(mask)))
}

func printWildcard(mask net.IPMask) {
	wildcard := DeriveWildcard(mask)

	fmt.Println(
		PadFirstColumn("Wildcard:"),
		Blue(PadSecondColumn(wildcard.String())),
		Yellow(FormatWildcardAsDottedBits(wildcard)))
}

func printNetwork(network *net.IPNet) {
	decimalNetmask, _ := network.Mask.Size()
	fmt.Println(
		PadFirstColumn("Network:"),
		Blue(PadSecondColumn(fmt.Sprintf("%s/%d", network.IP, decimalNetmask))),
		Yellow(FormatIpAsDottedBits(network.IP)))
}

func printBroadcast(ip net.IP, network *net.IPNet) {
	broadcast := DeriveBroadCast(ip, network)
	fmt.Println(
		PadFirstColumn("Broadcast:"),
		Blue(PadSecondColumn(broadcast.String())),
		Yellow(FormatIpAsDottedBits(broadcast)))
}

func printHostMin(network *net.IPNet) {
	hostMin := HostMin(network)
	fmt.Println(
		PadFirstColumn("Host min.:"),
		Blue(PadSecondColumn(hostMin)),
		Yellow(FormatIpAsDottedBits(hostMin)))
}

func printHostMax(ip net.IP, network *net.IPNet) {
	hostMax := HostMax(ip, network)
	fmt.Println(
		PadFirstColumn("Host max.:"),
		Blue(PadSecondColumn(hostMax)),
		Yellow(FormatIpAsDottedBits(hostMax)))
}

func printHostsPerNet(ip net.IP, network *net.IPNet) {
	number := HostsPerNetwork(ip, network)
	fmt.Println(
		PadFirstColumn("Hosts/Net:"),
		Blue(PadSecondColumn(fmt.Sprintf("%d", number))))
}
