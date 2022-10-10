package helpers

import "strings"

func Quote(str string) string {
	return strings.ReplaceAll(str, "'", "''")
}

func Unquote(str string) string {
	return strings.ReplaceAll(str, "''", "'")
}
