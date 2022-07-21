package anagram

// Anagram reports whether 2 strings are anagrams of each other.
// Assumes only lower letters in both strings.
func Anagram(a string, b string) bool {
	aCounts := count(a)
	bCounts := count(b)
	return aCounts == bCounts
}

// count number of chars
func count(s string) [26]int {
	counts := [26]int{}
	for _, c := range s {
		counts[c-'a']++
	}
	return counts
}
