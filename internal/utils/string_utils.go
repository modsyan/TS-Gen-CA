package utils

import (
	"strings"
)

func ConvertToPascalCase(s string) string {
	words := strings.Split(s, "-")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

func ParseProperties(propertiesString string) []string {
	return strings.Fields(propertiesString)
}

func RemoveLastNewline(s string) string {
	if strings.HasSuffix(s, "\n") {
		return s[:len(s)-1]
	}
	return s
}
