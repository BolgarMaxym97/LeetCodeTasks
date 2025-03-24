package main

import (
	"fmt"
	"math"
)

func maxContainers(n int, w int, maxWeight int) int {
	cellCount := n * n

	if cellCount*w <= maxWeight {
		return cellCount
	}
	return int(math.Floor(float64(maxWeight / w)))
}

func main() {
	fmt.Println(maxContainers(3, 5, 20))
}
