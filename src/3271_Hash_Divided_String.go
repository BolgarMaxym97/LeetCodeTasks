package main

import "fmt"

func stringHash(s string, k int) string {
	result := ""
	for i := 0; i < len(s); i += k {
		substring := s[i : i+k]
		sum := 0
		for j := 0; j < len(substring); j++ {
			sum += alphaNum(rune(substring[j]))
		}

		result += string(numAlpha(sum % 26))
	}

	return result
}

func alphaNum(letter rune) int {
	return int(letter - 'a')
}

func numAlpha(num int) rune {
	return rune(num + 'a')

}

func main() {
	fmt.Println(stringHash("abcd", 2))
}
