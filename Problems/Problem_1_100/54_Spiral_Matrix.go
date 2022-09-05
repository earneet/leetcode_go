package Problem_1_100

func spiralOrder(matrix [][]int) []int {
	l, r, t, b := 0, len(matrix[0]), 0, len(matrix)
	ans := make([]int, 0)
	for {
		if r <= l || t >= b {
			break
		}
		for i := t; i < r; i++ {
			ans = append(ans, matrix[t][i])
		}
		t += 1
		if r <= l || t >= b {
			break
		}
		for j := t; j < b; j++ {
			ans = append(ans, matrix[j][r-1])
		}
		r -= 1

		if r <= l || t >= b {
			break
		}
		for i := r - 1; i >= l; i-- {
			ans = append(ans, matrix[b-1][i])
		}
		b -= 1

		if r <= l || t >= b {
			break
		}
		for j := b - 1; j >= t; j-- {
			ans = append(ans, matrix[j][l])
		}
		l += 1
	}
	return ans
}
