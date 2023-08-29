package stringutil

import (
	"strings"
	"unsafe"
)

// IsBlank 检查一个字符串是否为空或只包含空白字符
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func FromBytes(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func ToBytes(s string) (b []byte) {
	if s == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
