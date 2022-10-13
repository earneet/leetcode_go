package Problem_401_500

import "sort"

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	pre := intervals[0][1]
	ans := 0
	for _, interval := range intervals[1:] {
		if interval[0] < pre {
			ans += 1
			pre = min(pre, interval[1])
		} else {
			pre = interval[1]
		}
	}
	return ans
}
