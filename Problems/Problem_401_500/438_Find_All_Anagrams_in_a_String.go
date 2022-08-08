package Problem_401_500

func counter(s string) map[byte]int {
	res := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		res[s[i]] += 1
	}
	return res
}

func allZero(dict map[byte]int) bool {
	for _, v := range dict {
		if v != 0 {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	pSize := len(p)
	sSize := len(s)
	if pSize > sSize {
		return []int{}
	}
	cnt := counter(p)
	for i := 0; i < pSize; i++ {
		cnt[s[i]] -= 1
	}
	for i := 0; i < sSize-pSize+1; i++ {
		if allZero(cnt) {
			ans = append(ans, i)
		}
		if i+pSize == sSize {
			break
		}
		cnt[s[i]] += 1
		cnt[s[i+pSize]] -= 1
	}
	return ans
}
