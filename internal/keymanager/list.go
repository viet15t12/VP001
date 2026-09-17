package keymanager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color" // Thư viện tạo màu
	"vietnguyen.io.vn/ssh_manager/internal/output"
	"vietnguyen.io.vn/ssh_manager/internal/window"
)

func ListKeys() {
	// 1. Định nghĩa bảng màu cho nhất quán
	titleColor := color.New(color.FgCyan, color.Bold)
	fileColor := color.New(color.FgGreen, color.Bold)
	pathColor := color.New(color.FgHiBlack) // Màu xám nhạt cho đường dẫn
	contentColor := color.New(color.FgWhite)
	errorColor := color.New(color.FgRed, color.Bold)
	//separatorColor := color.New(color.FgHiBlack)

	dir := KeyDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		errorColor.Fprintf(os.Stderr, "❌ Lỗi đọc thư mục key: %v\n", err)
		window.Pause()
		return
	}

	var keys []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".pub") {
			keys = append(keys, filepath.Join(dir, e.Name()))
		}
	}

	// 2. Xử lý trường hợp không có key nào
	if len(keys) == 0 {
		fmt.Println()
		titleColor.Println("⚠️  Không tìm thấy file key (.pub) nào trong thư mục.")
		pathColor.Printf("   Đường dẫn đã kiểm tra: %s\n", dir)
		fmt.Println()
		window.Pause()
		return
	}

	// 3. In tiêu đề tổng quan
	fmt.Println()
	titleColor.Printf("🔑 Tìm thấy %d key(s) trong: %s\n", len(keys), dir)
	output.Divider()
	// 4. In chi tiết từng key
	for _, k := range keys {
		fileName := filepath.Base(k)

		// In thông tin file
		fileColor.Printf("📄 File: %s\n", fileName)
		pathColor.Printf("   └─ Đường dẫn: %s\n", k)

		content, err := os.ReadFile(k)
		if err != nil {
			errorColor.Printf("   └─ ❌ Lỗi đọc file: %v\n", err)
		} else {
			keyStr := strings.TrimSpace(string(content))

			// Mẹo: Nếu key quá dài, bạn có thể chỉ hiển thị 1 phần hoặc in ra toàn bộ
			// Ở đây mình in toàn bộ nhưng thụt đầu dòng cho đẹp
			contentColor.Printf("   └─ Nội dung:\n      %s\n", keyStr)
		}

		// Đường kẻ phân cách giữa các key
		fmt.Println()
		output.Divider()
	}

	window.Pause()
}
