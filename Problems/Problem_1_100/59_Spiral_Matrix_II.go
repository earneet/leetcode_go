package Problem_1_100

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	top, bottom, left, right := 0, n-1, 0, n-1

	idx := 1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			ans[top][i] = idx
			idx += 1
		}
		top += 1
		for i := top; i <= bottom; i++ {
			ans[i][right] = idx
			idx += 1
		}
		right -= 1
		for i := right; i >= left; i-- {
			ans[bottom][i] = idx
			idx += 1
		}
		bottom -= 1
		for i := bottom; i >= top; i-- {
			ans[i][left] = idx
			idx += 1
		}
		left += 1
	}
	return ans
}
