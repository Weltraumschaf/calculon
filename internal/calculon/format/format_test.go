package format

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func createNetmask(cidr string) net.IPMask {
	_, network, _ := net.ParseCIDR(cidr)
	return network.Mask
}

func TestFormatByteAsBits(t *testing.T) {
	assert.Equal(t, "00000000", ByteAsBits(0x00))
	assert.Equal(t, "00000001", ByteAsBits(0x01))
	assert.Equal(t, "00000010", ByteAsBits(0x02))
	assert.Equal(t, "00000011", ByteAsBits(0x03))
	assert.Equal(t, "00000100", ByteAsBits(0x04))
	assert.Equal(t, "00000101", ByteAsBits(0x05))
	assert.Equal(t, "00000110", ByteAsBits(0x06))
	assert.Equal(t, "00000111", ByteAsBits(0x07))
	assert.Equal(t, "00001000", ByteAsBits(0x08))
	assert.Equal(t, "00010000", ByteAsBits(0x10))
	assert.Equal(t, "00100000", ByteAsBits(0x20))
	assert.Equal(t, "01000000", ByteAsBits(0x40))
	assert.Equal(t, "10000000", ByteAsBits(0x80))
	assert.Equal(t, "11111111", ByteAsBits(0xFF))
}

func TestFormatIpAsDottedBits(t *testing.T) {
	result := IpAsDottedBits(net.ParseIP("192.168.1.1"))
	assert.Equal(t, "11000000.10101000.00000001.00000001", result)

	result = IpAsDottedBits(net.ParseIP("1.2.3.4"))
	assert.Equal(t, "00000001.00000010.00000011.00000100", result)
}

func TestFormatMaskAsDottedBits(t *testing.T) {
	result := MaskAsDottedBits(createNetmask("192.168.1.1/24"))
	assert.Equal(t, "11111111.11111111.11111111.00000000", result)

	result = MaskAsDottedBits(createNetmask("192.168.1.1/18"))
	assert.Equal(t, "11111111.11111111.11000000.00000000", result)

	result = MaskAsDottedBits(createNetmask("192.168.1.1/16"))
	assert.Equal(t, "11111111.11111111.00000000.00000000", result)

	result = MaskAsDottedBits(createNetmask("192.168.1.1/8"))
	assert.Equal(t, "11111111.00000000.00000000.00000000", result)
}

func TestFormatMaskAsDottedDecimal(t *testing.T) {
	result := MaskAsDottedDecimal(createNetmask("192.168.1.1/24"))
	assert.Equal(t, "255.255.255.0", result)

	result = MaskAsDottedDecimal(createNetmask("192.168.1.1/18"))
	assert.Equal(t, "255.255.192.0", result)

	result = MaskAsDottedDecimal(createNetmask("192.168.1.1/16"))
	assert.Equal(t, "255.255.0.0", result)

	result = MaskAsDottedDecimal(createNetmask("192.168.1.1/8"))
	assert.Equal(t, "255.0.0.0", result)
}

func TestFormatWildcardAsDottedBits(t *testing.T) {
	assert.Equal(t,
		"00000000.00000000.00000000.11111111",
		WildcardAsDottedBits(net.ParseIP("0.0.0.255")))
	assert.Equal(t,
		"00000000.00000000.11111111.11111111",
		WildcardAsDottedBits(net.ParseIP("0.0.255.255")))
	assert.Equal(t,
		"00000000.11111111.11111111.11111111",
		WildcardAsDottedBits(net.ParseIP("0.255.255.255")))
}

func TestPadFirstColumn(t *testing.T) {
	assert.Equal(t,
		"foo        ",
		PadFirstColumn("foo"))
}

func TestPadSecondColumn(t *testing.T) {
	assert.Equal(t,
		"foo                   ",
		PadSecondColumn("foo"))
}
