package comma

import (
	"bytes"
	"strings"
)

// Comma takes a string representatin of an integer,
// and insert commas every three places.
func Comma(s string) string {
	n := len(s)
	// base case
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

// Comma2 uses bytes.Buffer instead of straing concatenation
// in a non-recursive manner
func Comma2(s string) string {
	buf := bytes.Buffer{}

	// handle the first chunk
	offset := len(s) % 3
	buf.WriteString(s[:offset])

	// handle the remainding chunks with size 3
	chunks := len(s) / 3
	for i := 0; i < chunks; i++ {
		if offset > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[offset+i*3 : offset+i*3+3])
	}

	return buf.String()
}

// Comma3 deals with floating-point numbers and an optional sign
func Comma3(s string) string {
	prefix, suffix := "", ""
	// determines sign
	if len(s) > 0 && s[0] == '-' {
		prefix = "-"
		s = s[1:]
	}
	// find decimal numbers
	if i := strings.LastIndex(s, "."); i >= 0 {
		suffix = s[i:]
		s = s[:i]
	}
	return prefix + Comma(s) + suffix
}
