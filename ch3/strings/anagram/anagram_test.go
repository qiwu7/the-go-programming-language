package anagram

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		a       string
		b       string
		anagram bool
	}{
		{
			a:       "abc",
			b:       "cba",
			anagram: true,
		},
		{
			a:       "about",
			b:       "toubaa",
			anagram: false,
		},
		{
			a:       "aabbcc",
			b:       "ccbbaa",
			anagram: true,
		},
		{
			a:       "aaabcc",
			b:       "cccbaa",
			anagram: false,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("a:%s, b:%s", test.a, test.b)
		t.Run(name, func(t *testing.T) {
			res := Anagram(test.a, test.b)
			if res != test.anagram {
				t.Errorf("got %v, expected: %v", res, test.anagram)
			}
		})
	}
}
