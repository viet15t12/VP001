package window

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // linux, darwin (macOS)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func EnterAltScreen() {
	fmt.Print("\x1b[?1049h")
}

func ExitAltScreen() {
	fmt.Print("\x1b[?1049l")
}
