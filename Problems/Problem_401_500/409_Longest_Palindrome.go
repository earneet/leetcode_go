package Problem_401_500

func longestPalindrome(s string) int {
	stat := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		v, e := stat[s[i]]
		if !e {
			v = 1
		} else {
			v += 1
		}
		stat[s[i]] = v
	}

	ans := 0
	odd := 0
	for _, v := range stat {
		if v&1 == 0 {
			ans += v
		} else {
			ans += v - 1
			odd = 1
		}
	}
	return ans + odd
}
