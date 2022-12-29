package color

import "fmt"

type Color string

const (
	Black           Color = "\033[0;30m"
	Red             Color = "\033[0;31m"
	Green           Color = "\033[0;32m"
	Yellow          Color = "\033[0;33m"
	Blue            Color = "\033[0;34m"
	Magenta         Color = "\033[0;35m"
	Cyan            Color = "\033[0;36m"
	White           Color = "\033[0;37m"
	BrightBlackGray Color = "\033[0;90m"
	BrightRed       Color = "\033[0;91m"
	BrightGreen     Color = "\033[0;92m"
	BrightYellow    Color = "\033[0;93m"
	BrightBlue      Color = "\033[0;94m"
	BrightMagenta   Color = "\033[0;95m"
	BrightCyan      Color = "\033[0;96m"
	BrightWhite     Color = "\033[0;97m"
	None            Color = "\033[0m"
)

func Colorize[T string | int](v T, to Color) string {
	return fmt.Sprintf("%s%s%v%s", None, to, v, None)
}

func ColorizeFirstChar(s string, to Color) string {
	return fmt.Sprintf("%s%s", Colorize(s[:1], to), s[1:])
}
