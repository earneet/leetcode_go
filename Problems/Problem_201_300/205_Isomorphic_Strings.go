package Problem_201_300

func isIsomorphic(s string, t string) bool {
	dict := make(map[byte]byte)
	dict2 := make(map[byte]byte)
	for i := 0; i < len(t); i++ {
		if m, e := dict[s[i]]; !e {
			if _, e2 := dict2[t[i]]; e2 {
				return false
			}
			dict[s[i]] = t[i]
			dict2[t[i]] = s[i]
		} else if m != t[i] {
			return false
		}
	}
	return true
}
