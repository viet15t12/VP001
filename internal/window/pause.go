package window

import (
	"bufio"
	"fmt"
	"os"
)

func Pause() {
	fmt.Println("Nhấn Enter để thoát...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
