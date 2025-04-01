package gon

import (
	"regexp"
	"strings"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func Uncapitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

func PascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = Capitalize(parts[i])
	}
	return strings.Join(parts, "")
}

func SnakeCase(s string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}
