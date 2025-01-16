package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	lessArr, biggerArr := getArraysOrderedBySize(nums1, nums2)
	numsMap := make(map[int]int, len(lessArr))

	for _, v := range lessArr {
		numsMap[v] = 1
	}

	var resultArr []int

	for _, v := range biggerArr {
		if _, ok := numsMap[v]; ok {
			resultArr = append(resultArr, v)
			delete(numsMap, v)
		}
	}

	return resultArr
}

func getArraysOrderedBySize(nums1 []int, nums2 []int) ([]int, []int) {
	if len(nums1) > len(nums2) {
		return nums2, nums1
	}
	return nums1, nums2
}

func main() {
	fmt.Println(intersection(
		[]int{1, 2, 2, 1, 9},
		[]int{2, 2, 9},
	))

}
