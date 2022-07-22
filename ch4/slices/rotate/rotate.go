package rotate

import "57/the-go-programming-language/ch4/slices/rev"

// Rotate a slice left by n elements
func Rotate(n int, s []int) {
	if len(s) < 2 {
		return
	}
	n = n % len(s)
	// reverse the leading n
	rev.Reverse(s[:n])
	// reverse the remaining
	rev.Reverse(s[n:])
	// reverse whole slice
	rev.Reverse(s)
}
