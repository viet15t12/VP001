package output

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func TerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= 0 {
		return 80 // fallback khi không lấy được (vd chạy trong CI, redirect ra file)
	}
	return width
}

func Divider() {
	fmt.Println(strings.Repeat("─", TerminalWidth()))
}
