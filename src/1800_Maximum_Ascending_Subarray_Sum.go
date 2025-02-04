package main

import (
	"fmt"
	"slices"
)

func maxAscendingSum(nums []int) int {
	prev, sum, lenNums := 0, 0, len(nums)
	var sumArr []int

	for i, v := range nums {
		isBigger := i == 0 || v > prev
		prev = v

		if isBigger {
			sum += v

			if i == lenNums-1 {
				sumArr = append(sumArr, sum)
			}
			continue
		}

		sumArr = append(sumArr, sum)
		sum = v
	}

	return slices.Max(sumArr)
}

func main() {
	fmt.Println(maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
}
