package Problem_301_400

func isPowerOfFour(n int) bool {
	for i := 0; n != 0 && i < 31; i += 2 {
		if n & ^(1<<i) == 0 {
			return true
		}
	}
	return false
}
