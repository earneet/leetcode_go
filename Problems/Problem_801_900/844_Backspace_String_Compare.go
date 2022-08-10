package Problem_801_900

func backspaceCompare(s string, t string) bool {
	stackS := make([]rune, 0)
	stackT := make([]rune, 0)

	for _, v := range s {
		if v == '#' {
			if len(stackS) > 0 {
				stackS = stackS[:len(stackS)-1]
			}
		} else {
			stackS = append(stackS, v)
		}
	}

	for _, v := range t {
		if v == '#' {
			if len(stackT) > 0 {
				stackT = stackT[:len(stackT)-1]
			}
		} else {
			stackT = append(stackT, v)
		}
	}

	if len(stackT) != len(stackS) {
		return false
	}
	for i := 0; i < len(stackT); i++ {
		if stackT[i] != stackS[i] {
			return false
		}
	}
	return true
}
