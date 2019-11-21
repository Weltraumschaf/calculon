package calculon

import (
    "errors"
    "fmt"
    "net"
    "strconv"
    "strings"
)

func PrintResult(ip net.IP, netmask *net.IPNet) error {
    binaryIp, err := FormatBinary(ip)

    if err != nil {
        return nil
    }

    fmt.Printf("Address:   %s       %s\n", ip, binaryIp)

    //binaryNetmask, err := NetmaskToBinary(netmask)
    //
    //if err != nil {
    //    return nil
    //}
    //
    //decimalNetmask, err := NetmaskToOctets(netmask)
    //
    //if err != nil {
    //    return nil
    //}

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

func FormatByteToBitString(input byte) string {
    binary := strconv.FormatUint(uint64(input), 2)
    return fmt.Sprintf("%08s", binary)
}