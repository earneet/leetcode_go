package Problem_201_300

func isBadVersion(_ int) bool {
	return false
}

func firstBadVersion(n int) int {
	low := 1
	high := n
	for low <= high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high + 1
}
