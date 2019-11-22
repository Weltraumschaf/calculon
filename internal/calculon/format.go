package calculon

import (
    "fmt"
    "net"
    "strconv"
    "strings"
)

func FormatByteAsBits(input byte) string {
    binary := strconv.FormatUint(uint64(input), 2)
    return fmt.Sprintf("%08s", binary)
}

func FormatIpAsDottedBits(ip net.IP) (string, error) {
    var result []string

    for _, b := range ip.To4() {
        result = append(result, FormatByteAsBits(b))
    }

    return strings.Join(result, "."), nil
}