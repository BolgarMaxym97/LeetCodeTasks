package main

import "fmt"

func search(nums []int, target int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2
		guessTarget := nums[middleIndex]

		if guessTarget == target {
			return middleIndex
		}

		if guessTarget > target {
			rightIndex = middleIndex - 1
			continue
		}

		leftIndex = middleIndex + 1
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
