package main

import "fmt"

func countNegatives(grid [][]int) int {
	negativeCnt := 0

	for _, row := range grid {

		left, right := 0, len(row)

		for left < right {
			mid := (left + right) / 2

			if row[mid] < 0 {
				right = mid
			} else {
				left = mid + 1
			}
		}

		negativeCnt += len(row) - left
	}

	return negativeCnt
}

func main() {
	fmt.Println(countNegatives([][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}))
}
