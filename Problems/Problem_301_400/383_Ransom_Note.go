package Problem_301_400

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	dict := make(map[rune]int)
	for _, c := range ransomNote {
		dict[c] -= 1
	}

	for _, c := range magazine {
		dict[c] += 1
	}

	for _, v := range dict {
		if v < 0 {
			return false
		}
	}
	return true
}
