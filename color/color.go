package color

import "fmt"

func Colorize(s string, r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, s)
}
