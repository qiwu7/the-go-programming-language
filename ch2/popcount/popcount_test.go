package popcount

import (
	"fmt"
	"testing"
)

func testPopCount(t *testing.T, fn func(x uint64) int) {
	tests := []struct {
		number   uint64
		expected int
	}{
		{number: 1, expected: 1},
		{number: 0xffef, expected: 15},
		{number: 0x1234, expected: 5},
	}

	for _, test := range tests {
		name := fmt.Sprintf("number:%d,expected:%d", test.number, test.expected)
		t.Run(name, func(t *testing.T) {
			res := fn(test.number)
			if res != test.expected {
				t.Errorf("got %d, expected: %d", res, test.expected)
			}
		})
	}
}

func TestPopCount(t *testing.T) {
	testPopCount(t, PopCount)
}

func TestPopCountLoop(t *testing.T) {
	testPopCount(t, PopCountLoop)
}
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1234567890)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(1234567890)
	}
}
