package virtualfilesystem

import (
	"unicode"
)

func isNameValid(str string) bool {
	for _, char := range str {
		if char != '-' && !unicode.Is(unicode.Latin, char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func isLengthExcessive(args string, limit int) bool {
	if len(args) > limit {
		return true
	}
	return false
}
