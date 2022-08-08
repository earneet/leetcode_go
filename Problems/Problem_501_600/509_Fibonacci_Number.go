package Problem_501_600

func fib(n int) int {
	pre := 0
	cur := 1
	for i := 0; i < n; i++ {
		tmp := cur
		cur += pre
		pre = tmp
	}
	return pre
}
