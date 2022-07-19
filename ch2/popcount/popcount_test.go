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

func TestPopCountShift(t *testing.T) {
	testPopCount(t, PopCountShift)
}

func TestPopCountClear(t *testing.T) {
	testPopCount(t, PopCountClear)
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

/**
 * cpu: Intel(R) Core(TM) i5-8210Y CPU @ 1.60GHz
 * BenchmarkPopCount-4        	1000000000	         0.3269 ns/op	       0 B/op	       0 allocs/op
 * BenchmarkPopCountLoop-4    	1000000000	         0.2861 ns/op	       0 B/op	       0 allocs/op
 * BenchmarkPopCountShift-4   	1000000000	         0.2905 ns/op	       0 B/op	       0 allocs/op
 * BenchmarkPopCountClear-4   	1000000000	         0.2870 ns/op	       0 B/op	       0 allocs/op
 * PASS
 */
