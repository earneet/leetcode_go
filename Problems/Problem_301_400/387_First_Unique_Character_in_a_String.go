package Problem_301_400

func firstUniqChar(s string) int {
	stat := make(map[rune]int)
	for _, c := range s {
		stat[c] += 1
	}

	for i, c := range s {
		if stat[c] == 1 {
			return i
		}
	}
	return -1
}
