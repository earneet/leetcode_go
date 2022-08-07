package Problem_301_400

func isSubsequence(s string, t string) bool {
	s_size := len(s)
	t_size := len(t)
	if s_size > t_size {
		return false
	}

	i := 0
	j := 0
	for i < s_size && j < t_size {
		if s[i] == t[j] {
			i += 1
			j += 1
		} else {
			j += 1
		}
	}
	return i == s_size
}
