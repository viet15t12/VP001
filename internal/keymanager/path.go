// internal/keymanager/path.go
package keymanager

import (
	"os"
	"path/filepath"
)

// KeyDir trả về đường dẫn thư mục chứa SSH key
func KeyDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".ssh")
}
