package stringspretty

import "fmt"

// ANSI Color Codes
const (
	ColorDefault = "\x1b[39m"
	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[38;5;41m"
	ColorBlue    = "\x1b[38;5;117m"
	ColorGray    = "\x1b[90m"
	ColorPurple  = "\x1b[38;5;129m"
	ColorPink    = "\x1b[38;5;219m"
)

func Color(color string, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, ColorDefault)
}

func Red(s string) string {
	return Color(ColorRed, s)
}

func Green(s string) string {
	return Color(ColorGreen, s)
}

func Blue(s string) string {
	return Color(ColorBlue, s)
}

func Gray(s string) string {
	return Color(ColorGray, s)
}

func Purple(s string) string {
	return Color(ColorPurple, s)
}

func Pink(s string) string {
	return Color(ColorPink, s)
}
