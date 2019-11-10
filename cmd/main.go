package cmd

import (
    "errors"
    "fmt"
    "github.com/urfave/cli"
    "net"
    "os"
    "strconv"
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
    app.Action = Action()

    err := app.Run(os.Args)

    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }
}

func Action() func(c *cli.Context) error {
    return func(c *cli.Context) error {
        firstArg := c.Args().First()

        if len(firstArg) == 0 {
            return errors.New("No IP given!")
        }

        ip := net.ParseIP(ExtractIpAddress(firstArg))
        netmask, err := strconv.Atoi(ExtractNetMask(firstArg))

        if err != nil {
            return errors.New("Invalid netmask given!")
        }

        return printResult(ip, netmask)
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
        return ipWithNetMask[position+1:]
    }

    return ""
}

func printResult(ip net.IP, netmask int) error {
    binaryIp, err := FormatBinary(ip)

    if err != nil {
        return nil
    }

    fmt.Printf("Address:   %s       %s\n", ip, binaryIp)

    binaryNetmask, err := NetmaskToBinary(netmask)

    if err != nil {
        return nil
    }

    decimalNetmask, err := NetmaskToOctets(netmask)

    if err != nil {
        return nil
    }

    fmt.Printf("Netmask:   %s = %d  %s\n",
        decimalNetmask, netmask, binaryNetmask)
    fmt.Printf("Wildcard:  \n")
    fmt.Printf("Network:   \n")
    fmt.Printf("Broadcast: \n")
    fmt.Printf("HostMin:   \n")
    fmt.Printf("HostMax:   \n")
    fmt.Printf("Hosts/Net: \n")
    return nil
}

func FormatBinary(ip net.IP) (string, error) {
    var result []string

    for _, octet := range strings.Split(ip.String(), ".") {
        o, err := strconv.Atoi(octet)

        if err != nil {
            return "", err
        }

        binaryOctet := strconv.FormatInt(int64(o), 2)
        binaryOctet = fmt.Sprintf("%08s", binaryOctet)
        result = append(result, binaryOctet)
    }

    return strings.Join(result, "."), nil
}

func NetmaskToOctets(netmask int) (string, error) {
    octets, err := NetmaskToBytes(netmask)

    if err != nil {
        return "", err
    }

    var result []string

    for _, octet := range octets {
        binaryOctet := strconv.FormatInt(int64(octet), 10)
        result = append(result, binaryOctet)
    }

    return strings.Join(result, "."), nil
}

func NetmaskToBinary(netmask int) (string, error) {
    octets, err := NetmaskToBytes(netmask)

    if err != nil {
        return "", err
    }

    var result []string

    for _, octet := range octets {
        binaryOctet := strconv.FormatInt(int64(octet), 2)
        binaryOctet = fmt.Sprintf("%08s", binaryOctet)
        result = append(result, binaryOctet)
    }

    return strings.Join(result, "."), nil
}

const maxNetmask = 32

func NetmaskToBytes(netmask int) ([]uint64, error) {
    if netmask < 1 {
        return nil, errors.New(
            fmt.Sprintf("Invalid netmask! Value is less than one: %d", netmask))
    }

    if netmask > maxNetmask {
        return nil, errors.New(
            fmt.Sprintf("Invalid netmask! Value is greater than %d: %d", maxNetmask, netmask))
    }

    var binaryNetmask uint64
    binaryNetmask = (0xFFFFFFFF << (maxNetmask - netmask)) & 0xFFFFFFFF
    var countDownBits uint64
    countDownBits = maxNetmask
    octets := make([]uint64, 0, 4)

    for i := 1; i <= 4; i++ {
        tmp := binaryNetmask >> (countDownBits - 8) & 0xFF
        octets = append(octets, tmp)
        countDownBits -= 8
    }

    return octets, nil
}

