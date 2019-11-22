package calculon

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

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