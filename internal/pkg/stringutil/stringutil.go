package stringutil

import (
	"strings"
)

// IsBlank 检查一个字符串是否为空或只包含空白字符
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
