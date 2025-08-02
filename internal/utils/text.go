package utils

import "unicode/utf8"

const MaxLengthComments = 72

func truncateText(text string, maxLength int) string {
	if utf8.RuneCountInString(text) <= maxLength {
		return text
	}

	truncated := string([]rune(text)[:maxLength-3])
	return truncated + "..."
}

func TruncateComment(comment string) string {
	return truncateText(comment, MaxLengthComments)
}
