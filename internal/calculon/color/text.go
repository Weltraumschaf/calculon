package color

import "weltraumschaf.de/calculon/internal/calculon/config"

type Color string

const (
	reset  Color = "\033[0m"
	red          = "\033[31m"
	green        = "\033[32m"
	yellow       = "\033[33m"
	blue         = "\033[34m"
	purple       = "\033[35m"
	cyan         = "\033[36m"
	white        = "\033[37m"
)

func Red(str string) string {
	return colorize(red, str)
}

func Green(str string) string {
	return colorize(green, str)
}

func Yellow(str string) string {
	return colorize(yellow, str)
}

func Blue(str string) string {
	return colorize(blue, str)
}

func Purple(str string) string {
	return colorize(purple, str)
}

func Cyan(str string) string {
	return colorize(cyan, str)
}

func White(str string) string {
	return colorize(white, str)
}


func colorize(color Color, input string) string {
	var useColors = config.ApplicationOptions().IsColor()

	if useColors {
		return string(color) + input + string(reset)
	}

	return input
}
