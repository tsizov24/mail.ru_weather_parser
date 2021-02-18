package helper

import (
	"regexp"
	"strconv"
	"strings"
)

func cutFrom(s, index string) string {
	return s[:strings.Index(s, index)]
}

func cutSpaces(s string) string {
	return strings.TrimSpace(regexp.MustCompile("\\s+").ReplaceAllString(s, " "))
}

func strToNum(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

