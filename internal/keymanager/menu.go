package keymanager

import (
	"vietnguyen.io.vn/ssh_manager/internal/input"
	"vietnguyen.io.vn/ssh_manager/internal/messages"
	"vietnguyen.io.vn/ssh_manager/internal/output"
	"vietnguyen.io.vn/ssh_manager/internal/window"
)

func RunMenu() {
	for {

		window.ClearScreen()
		output.PrintKeyManager()

		switch input.ChoiceMenu(messages.MesssChoiceMneu) {
		case 0:
			window.ClearScreen()
			return
		case 2:
			window.ClearScreen()
			ListKeys()

		default:
			window.ClearScreen()
		}

	}

}
