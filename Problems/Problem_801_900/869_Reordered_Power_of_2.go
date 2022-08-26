package Problem_801_900

import "strconv"

func reorderedPowerOf2(n int) bool {
	nStr := strconv.Itoa(n)
	for i := 0; i < 31; i++ {
		p := strconv.Itoa(int(1 << i))
		if len(nStr) != len(p) {
			continue
		}

		dict := make(map[rune]int)
		for _, c := range p {
			dict[c] += 1
		}

		for _, c := range nStr {
			dict[c] -= 1
		}

		find := false
		for _, v := range dict {
			if v != 0 {
				find = true
				break
			}
		}
		if !find {
			return true
		}
	}

	return false
}
