package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Intro")
	intro()

	fmt.Println()
	fmt.Println("Basename1")
	fmt.Println("a/b/c.go -> basename = ", basename1("a/b/c.go"))
	fmt.Println("c.d.go -> basename = ", basename1("c.d.go"))
	fmt.Println("abc -> basename = ", basename1("abc"))
	fmt.Println("Basename2")
	fmt.Println("a/b/c.go -> basename = ", basename2("a/b/c.go"))
	fmt.Println("c.d.go -> basename = ", basename2("c.d.go"))
	fmt.Println("abc -> basename = ", basename2("abc"))

	fmt.Println()
	fmt.Println("Print Ints")
	fmt.Println(intsToString([]int{1, 3, 5, 7, 0}))
}

func intro() {
	// len returns the number of btes (not runes) in a string
	// the index operations s[i] retrieves the ith byte of string s
	s := "你好啊"
	fmt.Printf("len(%s) = %d\n", s, len(s))
	fmt.Println(s[3])
}

// basename1 does all the work without any helper functions
func basename1(s string) string {
	// Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' && i < len(s)-1 {
			s = s[i+1:]
			break
		}
	}
	// Discard file extension if any
	for i := 0; i < len(s); i++ {
		if (s[i]) == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//  basename2 uses strings.lastIndex lib function
func basename2(s string) string {
	// Discard last '/' and everything before
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	// Discard file extension if any
	dot := strings.LastIndex(s, ".")
	if dot > 0 {
		return s[:dot-1]
	}
	return s
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
