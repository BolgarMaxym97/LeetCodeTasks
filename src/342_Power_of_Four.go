package main

import "fmt"

func isPowerOfFourLoop(n int) bool {
	if n <= 0 {
		return false
	}
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

func isPowerOfFourRecursion(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 || n%4 != 0 {
		return false
	}

	return isPowerOfFourRecursion(n / 4)
}

func main() {
	fmt.Println(isPowerOfFourLoop(16))
	fmt.Println(isPowerOfFourRecursion(16))
}
