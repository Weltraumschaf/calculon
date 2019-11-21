package calculon

import (
    "github.com/stretchr/testify/assert"
    "net"
    "testing"
)

func TestFormatBinary(t *testing.T) {
    result, err := FormatBinary(net.ParseIP("192.168.123.5"))

    assert.Empty(t, err)
    assert.Equal(t, "11000000.10101000.01111011.00000101", result)
}

func TestNetmaskToOctets(t *testing.T) {
    result, err := NetmaskToOctets(24)

    assert.Empty(t, err)
    assert.Equal(t, "255.255.255.0", result)
}

func TestNetmaskToBinary(t *testing.T) {
    result, err := NetmaskToBinary(24)

    assert.Empty(t, err)
    assert.Equal(t, "11111111.11111111.11111111.00000000", result)
}

func TestNetmaskToBytes(t *testing.T) {
    result, err := NetmaskToBytes(24)

    assert.Empty(t, err)
    assert.Equal(t, []uint64{0xFF, 0xFF, 0xFF, 0x00}, result)
}

func TestNetmaskToBytes_netmaskToSmall(t *testing.T) {
    result, err := NetmaskToBytes(0)

    assert.Nil(t, result)
    assert.NotNil(t, err)
}

func TestNetmaskToBytes_netmaskToBig(t *testing.T) {
    result, err := NetmaskToBytes(33)

    assert.Nil(t, result)
    assert.NotNil(t, err)
}

func TestFoo(t *testing.T) {
    ip, network, _ := net.ParseCIDR("192.168.123.5/24")
    bytes, _ := ip.To4().MarshalText()

    assert.Equal(t, nil, bytes)
    assert.Equal(t, "", ip.String())
    assert.Equal(t, "", network.String())
}

func TestFormatByteToBitString(t *testing.T) {
    assert.Equal(t, "00000000", FormatByteToBitString(0x00))
    assert.Equal(t, "00000001", FormatByteToBitString(0x01))
    assert.Equal(t, "00000010", FormatByteToBitString(0x02))
    assert.Equal(t, "00000011", FormatByteToBitString(0x03))
    assert.Equal(t, "00000100", FormatByteToBitString(0x04))
    assert.Equal(t, "00000101", FormatByteToBitString(0x05))
    assert.Equal(t, "00000110", FormatByteToBitString(0x06))
    assert.Equal(t, "00000111", FormatByteToBitString(0x07))
    assert.Equal(t, "00001000", FormatByteToBitString(0x08))
    assert.Equal(t, "00010000", FormatByteToBitString(0x10))
    assert.Equal(t, "00100000", FormatByteToBitString(0x20))
    assert.Equal(t, "01000000", FormatByteToBitString(0x40))
    assert.Equal(t, "10000000", FormatByteToBitString(0x80))
    assert.Equal(t, "11111111", FormatByteToBitString(0xFF))
}