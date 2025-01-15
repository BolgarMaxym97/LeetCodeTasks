package main

import "fmt"

func arrangeCoins(n int) int {
	i := 1

	for n >= i {
		n -= i
		i++
	}

	return i - 1
}

func main() {
	fmt.Println(arrangeCoins(15))
}
