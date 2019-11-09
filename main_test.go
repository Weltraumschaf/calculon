package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestExtractIpAddress(t *testing.T) {
    assert.Equal(t, "192.168.123.0", ExtractIpAddress("192.168.123.0/24"))
}

func TestExtractIpAddress_emptySubnet(t *testing.T) {
    assert.Equal(t, "192.168.123.0", ExtractIpAddress("192.168.123.0/"))
}

func TestExtractIpAddress_noSubnet(t *testing.T) {
    assert.Equal(t, "192.168.123.0", ExtractIpAddress("192.168.123.0"))
}

func TestExtractIpAddress_emptyIp(t *testing.T) {
    assert.Equal(t, "", ExtractIpAddress("/24"))
}

func TestExtractIpAddress_empty(t *testing.T) {
    assert.Equal(t, "", ExtractIpAddress(""))
}

func TestExtractNetMask(t *testing.T) {
    assert.Equal(t, "24", ExtractNetMask("192.168.123.0/24"))
}

func TestExtractNetMask_emptyMask(t *testing.T) {
    assert.Equal(t, "", ExtractNetMask("192.168.123.0/"))
}

func TestExtractNetMask_noMask(t *testing.T) {
    assert.Equal(t, "", ExtractNetMask("192.168.123.0"))
}

func TestExtractNetMask_emptyIp(t *testing.T) {
    assert.Equal(t, "24", ExtractNetMask("/24"))
}

func TestExtractNetMask_empty(t *testing.T) {
    assert.Equal(t, "", ExtractNetMask(""))
}
