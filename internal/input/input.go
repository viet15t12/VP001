package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// readLine đọc chuỗi thô từ bàn phím, xóa khoảng trắng thừa
func readLine(text string) string {
	fmt.Print(text)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readStdInt(text string) (int, error) {
	return strconv.Atoi(readLine(text))
}

func readStdInt64(text string) (int64, error) {
	return strconv.ParseInt(readLine(text), 10, 64)
}

func readStdFloat32(text string) (float32, error) {
	n, err := strconv.ParseFloat(readLine(text), 32)
	return float32(n), err
}

func readStdFloat64(text string) (float64, error) {
	return strconv.ParseFloat(readLine(text), 64)
}

func readStdUint(text string) (uint, error) {
	n, err := strconv.ParseUint(readLine(text), 10, 64)
	return uint(n), err
}
