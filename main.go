package main

import (
	"fmt"

	"vietnguyen.io.vn/ssh_manager/internal/input"
	"vietnguyen.io.vn/ssh_manager/internal/keymanager"
	"vietnguyen.io.vn/ssh_manager/internal/messages"
	"vietnguyen.io.vn/ssh_manager/internal/output"
	"vietnguyen.io.vn/ssh_manager/internal/window"
)

func main() {

	//window.EnterAltScreen()
	//defer window.ExitAltScreen()

	for {
		output.PrintMenu()

		switch input.ChoiceMenu(messages.MesssChoiceMneu) {
		case 0:
			fmt.Println("\n", messages.Exit)
			return
		case 1:
			keymanager.RunMenu()
		default:
			window.ClearScreen()
		}
	}
}
