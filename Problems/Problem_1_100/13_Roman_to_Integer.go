package Problem_1_100

func romanToInt(s string) int {
	tables := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}

	ans := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && tables[s[i]] < tables[s[i+1]] {
			ans += tables[s[i+1]] - tables[s[i]]
			i += 1
		} else {
			ans += tables[s[i]]
		}
	}
	return ans
}
