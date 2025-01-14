package main

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {
	setA := make(map[int]bool)
	setB := make(map[int]bool)
	setCommon := make(map[int]bool)
	c := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		setA[A[i]] = true
		setB[B[i]] = true

		if setB[A[i]] {
			setCommon[A[i]] = true
		}
		if setA[B[i]] {
			setCommon[B[i]] = true
		}

		c[i] = len(setCommon)

	}

	return c
}

func main() {
	fmt.Println(findThePrefixCommonArray([]int{0, 3, 2, 4}, []int{3, 1, 4, 2}))

}
