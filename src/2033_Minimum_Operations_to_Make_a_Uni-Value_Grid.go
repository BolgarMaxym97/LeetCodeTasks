package main

import (
	"fmt"
	"math"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	operationsCount := 0
	var list []int
	for _, v := range grid {
		list = append(list, v...)
	}

	sort.Ints(list)
	median := list[len(list)/2]

	for _, v := range list {
		if (v-median)%x != 0 {
			return -1
		}

		operationsCount += int(math.Abs(float64(v-median))) / x
	}
	return operationsCount
}

func main() {
	fmt.Println(minOperations([][]int{{931, 128}, {639, 712}}, 73))
}
