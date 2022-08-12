package Problem_201_300

func isHappy(n int) bool {
	visited := make(map[int]bool)
	visited[n] = true
	for true {
		nn := 0
		for n != 0 {
			nn += (n % 10) * (n % 10)
			n = n / 10
		}
		n = nn
		if n == 1 {
			return true
		} else if visited[n] {
			break
		} else {
			visited[n] = true
		}
	}
	return false
}
