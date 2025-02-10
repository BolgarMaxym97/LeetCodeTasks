package main

import (
	"fmt"
	"unicode"
)

func clearDigits(s string) string {
	var arrayS []rune
	for _, v := range s {
		if !unicode.IsDigit(v) {
			arrayS = append(arrayS, v)
			continue
		}

		if len(arrayS) > 0 {
			arrayS = arrayS[:len(arrayS)-1]
		}
	}

	return string(arrayS)
}

func main() {
	fmt.Println(clearDigits("cb3asdas4"))
}
