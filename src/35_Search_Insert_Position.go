package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		guess := nums[mid]

		if guess == target {
			return mid
		}

		if guess < target {
			left = mid + 1
			continue
		}

		right = mid - 1
	}
	return left
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		guess := nums[mid]

		if guess == target {
			return mid
		}

		if guess < target {
			left = mid + 1
			continue
		}

		right = mid - 1
	}
	return 0
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6, 8, 11, 12, 14, 15, 17, 19, 22, 26, 24}, 13))
}
