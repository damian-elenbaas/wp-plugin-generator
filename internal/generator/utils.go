package generator

import (
	"errors"
	"strings"
	"unicode"
)

func splitOnCapitals(s string) []string {
	var parts []string
	last := 0

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			parts = append(parts, s[last:i])
			last = i
		}
	}
	parts = append(parts, s[last:])
	return parts
}

func toUpperWords(words *[]string) {
	for i := range *words {
		word := (*words)[i]
		(*words)[i] = strings.ToUpper(word)
	}
}

func toLowerWords(words *[]string) {
	for i := range *words {
		word := (*words)[i]
		(*words)[i] = strings.ToLower(word)
	}
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func capitalizeFirstLetterWords(words *[]string) {
	for i := range *words {
		word := (*words)[i]
		(*words)[i] = capitalizeFirstLetter(word)
	}
}

func getLastPartOfProjectName(projectName string) (*string, error) {
	parts := strings.Split(projectName, "/")
	if len(parts) <= 0 {
		return nil, errors.New("projectName is empty")
	}

	return &parts[len(parts)-1], nil
}
