package main

import (
	"fmt"
	"sort"
)

func mergeMeetings(meetings [][]int) [][]int {
	if len(meetings) == 0 {
		return meetings
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	merged := [][]int{meetings[0]}
	for _, m := range meetings[1:] {
		last := merged[len(merged)-1]
		if m[0] <= last[1] {
			last[1] = max(last[1], m[1])
		} else {
			merged = append(merged, m)
		}
	}
	return merged
}

func countDays(days int, meetings [][]int) int {
	merged := mergeMeetings(meetings)
	busyDays := 0

	for _, m := range merged {
		busyDays += m[1] - m[0] + 1
	}

	return days - busyDays
}

func main() {
	fmt.Println(countDays(5, [][]int{{2, 4}, {1, 3}}))
}
