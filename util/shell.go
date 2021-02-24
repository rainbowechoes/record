package util

import "strings"

// ReplaceUnixLine will replace unix line characters(\n) with blank characters("").
func ReplaceUnixLine(s string) string {
	return strings.Replace(s, "\n", "", -1)
}
