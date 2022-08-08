package Problem_1_100

func climbStairs(n int) int {
	pre := 1
	cur := 2

	for i := 1; i < n; i++ {
		temp := cur
		cur = pre + cur
		pre = temp
	}
	return pre
}
