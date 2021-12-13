package format

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func IpAsDottedBits(ip net.IP) string {
	var result []string

	for _, b := range ip.To4() {
		result = append(result, ByteAsBits(b))
	}

	return joinBytes(result)
}

func MaskAsDottedBits(mask net.IPMask) string {
	return bytesAsBits(mask)
}

func MaskAsDottedDecimal(mask net.IPMask) string {
	var result []string

	for _, b := range mask {
		result = append(result, fmt.Sprintf("%d", b))
	}

	return joinBytes(result)
}

func WildcardAsDottedBits(ip net.IP) string {
	return bytesAsBits(ip.To4())
}

func bytesAsBits(bytes []byte) string {
	var result []string

	for _, b := range bytes {
		result = append(result, ByteAsBits(b))
	}

	return joinBytes(result)
}

func ByteAsBits(input byte) string {
	binary := strconv.FormatUint(uint64(input), 2)
	return fmt.Sprintf("%08s", binary)
}

func PadFirstColumn(input interface{}) string {
	return fmt.Sprintf("%-11s", input)
}

func PadSecondColumn(input interface{}) string {
	return fmt.Sprintf("%-22s", input)
}

func joinBytes(result []string) string {
	return strings.Join(result, ".")
}
