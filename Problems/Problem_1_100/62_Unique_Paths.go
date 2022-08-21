package Problem_1_100

func uniquePaths(m int, n int) int {
	ans := 1
	if n > m {
		n = m ^ n
		m = m ^ n
		n = m ^ n
	}
	for i := 0; i < n-1; i++ {
		ans *= m + n - 2 - i
		ans /= i + 1
	}
	return ans
}
