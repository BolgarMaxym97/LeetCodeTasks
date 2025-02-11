package main

import (
	"fmt"
	"strings"
)

func removeOccurrences(s string, part string) string {
	index := strings.Index(s, part)
	for index >= 0 {
		s = s[:index] + s[index+len(part):]
		index = strings.Index(s, part)
	}
	return s
}

func main() {
	fmt.Println(removeOccurrences("daabcbaabcbc", "abc"))
}
