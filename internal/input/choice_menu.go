package input

// ChoiceMenu đọc lựa chọn của người dùng (số nguyên >= 0), lặp lại cho tới khi hợp lệ
func ChoiceMenu(prompt string) int {
	choice, err := readStdInt(prompt)
	if err != nil || choice < 0 {
		//fmt.Println(messages.InvalidChoice)
		return -1
	}
	return choice
}
