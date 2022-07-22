package rev

import (
	"fmt"
	"testing"
)

func testRev(t *testing.T, fn func([]int)) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{0, 1, 2, 3, 4, 5},
			output: []int{5, 4, 3, 2, 1, 0},
		},
		{
			input:  []int{0, 1, 2},
			output: []int{2, 1, 0},
		},
		{
			input:  []int{0},
			output: []int{0},
		},
		{
			input:  []int{0, 1},
			output: []int{1, 0},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("input: %v", test.input)
		t.Run(name, func(t *testing.T) {
			fn(test.input)
			if !Equal(test.input, test.output) {
				t.Errorf("got %v, expected: %v", test.input, test.output)
			}
		})
	}
}

func TestRev(t *testing.T) {
	testRev(t, Reverse)
}
