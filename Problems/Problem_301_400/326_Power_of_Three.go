package Problem_301_400

func isPowerOfThree(n int) bool {
	for n >= 3 {
		r := n % 3
		n /= 3
		if r != 0 {
			return false
		}
	}
	return n == 1
}
