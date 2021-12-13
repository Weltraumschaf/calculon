package calc

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func createNetmask(cidr string) net.IPMask {
	_, network, _ := net.ParseCIDR(cidr)
	return network.Mask
}

func TestDeriveWildcard(t *testing.T) {
	result := DeriveWildcard(createNetmask("0.0.0.0/24"))
	assert.Equal(t, net.ParseIP("0.0.0.255"), result)

	result = DeriveWildcard(createNetmask("0.0.0.0/16"))
	assert.Equal(t, net.ParseIP("0.0.255.255"), result)

	result = DeriveWildcard(createNetmask("0.0.0.0/8"))
	assert.Equal(t, net.ParseIP("0.255.255.255"), result)
}

func TestInvertByte(t *testing.T) {
	assert.Equal(t, byte(0x00), InvertByte(0xFF))
	assert.Equal(t, byte(0xFF), InvertByte(0x00))
	assert.Equal(t, byte(0x7F), InvertByte(0x80))
	assert.Equal(t, byte(0xF0), InvertByte(0x0F))
}
