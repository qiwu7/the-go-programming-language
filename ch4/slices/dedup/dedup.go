package dedup

func Dedup(s []string) []string {
	if len(s) == 0 {
		return s
	}
	i := 0
	for j := 0; j < len(s); {
		s[i] = s[j]
		for j < len(s) && s[i] == s[j] {
			j++
		}
		i++
	}
	return s[:i]
}

func Equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
