package algorithms

func QuickSort(arr []int) []int {
	arrLength := len(arr)
	if arrLength < 2 {
		return arr
	}

	index := arrLength / 2
	pivot := arr[index]
	var less, greater []int
	for i := 0; i < arrLength; i++ {
		if i == index {
			continue
		}

		if arr[i] < pivot {
			less = append(less, arr[i])
			continue
		}
		greater = append(greater, arr[i])
	}

	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
