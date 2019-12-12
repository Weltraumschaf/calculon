package calculon

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func FormatIpAsDottedBits(ip net.IP) string {
	var result []string

	for _, b := range ip.To4() {
		result = append(result, FormatByteAsBits(b))
	}

	return joinBytes(result)
}

func FormatMaskAsDottedBits(mask net.IPMask) string {
	return formatBytesAsBits(mask)
}

func FormatMaskAsDottedDecimal(mask net.IPMask) string {
	var result []string

	for _, b := range mask {
		result = append(result, fmt.Sprintf("%d", b))
	}

	return joinBytes(result)
}

func FormatWildcardAsDottedBits(ip net.IP) string {
	return formatBytesAsBits(ip.To4())
}

func formatBytesAsBits(bytes []byte) string {
	var result []string

	for _, b := range bytes {
		result = append(result, FormatByteAsBits(b))
	}

	return joinBytes(result)
}

func FormatByteAsBits(input byte) string {
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
