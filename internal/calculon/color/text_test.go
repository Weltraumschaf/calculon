package color

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRed(t *testing.T) {
	assert.Equal(t, "\u001B[31m|foo|\x1b[0m", Red("|foo|"))
}

func TestGreen(t *testing.T) {
	assert.Equal(t, "\u001B[32m|foo|\x1b[0m", Green("|foo|"))
}

func TestYellow(t *testing.T) {
	assert.Equal(t, "\u001B[33m|foo|\x1b[0m", Yellow("|foo|"))
}

func TestBlue(t *testing.T) {
	assert.Equal(t, "\u001B[34m|foo|\x1b[0m", Blue("|foo|"))
}

func TestPurple(t *testing.T) {
	assert.Equal(t, "\u001B[35m|foo|\x1b[0m", Purple("|foo|"))
}

func TestCyan(t *testing.T) {
	assert.Equal(t, "\u001B[36m|foo|\x1b[0m", Cyan("|foo|"))
}

func TestWhite(t *testing.T) {
	assert.Equal(t, "\u001B[37m|foo|\x1b[0m", White("|foo|"))
}

func TestColorize(t *testing.T) {
	var c Color = "color"
	assert.Equal(t, "color|foo|\x1b[0m", colorize(c, "|foo|"))
}
