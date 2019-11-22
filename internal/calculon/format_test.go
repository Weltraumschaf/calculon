package calculon

import (
    "github.com/stretchr/testify/assert"
    "net"
    "testing"
)

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
    result, _ := FormatIpAsDottedBits(net.ParseIP("192.168.1.1"))
    assert.Equal(t, "11000000.10101000.00000001.00000001", result)

    result, _ = FormatIpAsDottedBits(net.ParseIP("1.2.3.4"))
    assert.Equal(t, "00000001.00000010.00000011.00000100", result)
}
