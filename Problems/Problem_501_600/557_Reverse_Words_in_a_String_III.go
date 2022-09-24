package Problem_501_600

func reverseWords(s string) string {
	ans := make([]rune, 0)
	stack := make([]rune, 0)
	for _, c := range s {
		if c == ' ' {
			for len(stack) > 0 {
				ans = append(ans, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			ans = append(ans, c)
		} else {
			stack = append(stack, c)
		}
	}
	for len(stack) > 0 {
		ans = append(ans, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return string(ans)
}
