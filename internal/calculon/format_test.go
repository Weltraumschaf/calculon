package calculon

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
    assert.Equal(t, "00000000", FormatByteAsBits(0x00))
    assert.Equal(t, "00000001", FormatByteAsBits(0x01))
    assert.Equal(t, "00000010", FormatByteAsBits(0x02))
    assert.Equal(t, "00000011", FormatByteAsBits(0x03))
    assert.Equal(t, "00000100", FormatByteAsBits(0x04))
    assert.Equal(t, "00000101", FormatByteAsBits(0x05))
    assert.Equal(t, "00000110", FormatByteAsBits(0x06))
    assert.Equal(t, "00000111", FormatByteAsBits(0x07))
    assert.Equal(t, "00001000", FormatByteAsBits(0x08))
    assert.Equal(t, "00010000", FormatByteAsBits(0x10))
    assert.Equal(t, "00100000", FormatByteAsBits(0x20))
    assert.Equal(t, "01000000", FormatByteAsBits(0x40))
    assert.Equal(t, "10000000", FormatByteAsBits(0x80))
    assert.Equal(t, "11111111", FormatByteAsBits(0xFF))
}

func TestFormatIpAsDottedBits(t *testing.T) {
    result := FormatIpAsDottedBits(net.ParseIP("192.168.1.1"))
    assert.Equal(t, "11000000.10101000.00000001.00000001", result)

    result = FormatIpAsDottedBits(net.ParseIP("1.2.3.4"))
    assert.Equal(t, "00000001.00000010.00000011.00000100", result)
}

func TestFormatMaskAsDottedBits(t *testing.T) {
    result := FormatMaskAsDottedBits(createNetmask("192.168.1.1/24"))
    assert.Equal(t, "11111111.11111111.11111111.00000000", result)

    result = FormatMaskAsDottedBits(createNetmask("192.168.1.1/18"))
    assert.Equal(t, "11111111.11111111.11000000.00000000", result)

    result = FormatMaskAsDottedBits(createNetmask("192.168.1.1/16"))
    assert.Equal(t, "11111111.11111111.00000000.00000000", result)

    result = FormatMaskAsDottedBits(createNetmask("192.168.1.1/8"))
    assert.Equal(t, "11111111.00000000.00000000.00000000", result)
}

func TestFormatMaskAsDottedDecimal(t *testing.T) {
    result := FormatMaskAsDottedDecimal(createNetmask("192.168.1.1/24"))
    assert.Equal(t, "255.255.255.0", result)

    result = FormatMaskAsDottedDecimal(createNetmask("192.168.1.1/18"))
    assert.Equal(t, "255.255.192.0", result)

    result = FormatMaskAsDottedDecimal(createNetmask("192.168.1.1/16"))
    assert.Equal(t, "255.255.0.0", result)

    result = FormatMaskAsDottedDecimal(createNetmask("192.168.1.1/8"))
    assert.Equal(t, "255.0.0.0", result)
}
