package output

import (
	"fmt"

	"vietnguyen.io.vn/ssh_manager/internal/messages"
)

func PrintMenu() {
	Divider()
	fmt.Println(messages.Welcome)
	fmt.Println(messages.MenuPrompt)
	Divider()

	for _, item := range messages.MenuItems {
		fmt.Print(item)
	}

}

func PrintKeyManager()  {
	Divider()
	fmt.Println(messages.MenuPrompt)
	Divider()
	for _, item := range messages.MenuItemKeyManager {
		fmt.Print(item)
	}

}