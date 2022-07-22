package rotate

import (
	"57/the-go-programming-language/ch4/slices/rev"
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input  []int
		n      int
		output []int
	}{
		{
			input:  []int{0, 1, 2, 3, 4, 5},
			n:      2,
			output: []int{2, 3, 4, 5, 0, 1},
		},
		{
			input:  []int{0, 1, 2},
			n:      0,
			output: []int{0, 1, 2},
		},
		{
			input:  []int{0},
			n:      1,
			output: []int{0},
		},
		{
			input:  []int{0, 1},
			n:      2,
			output: []int{0, 1},
		},
		{
			input:  []int{0, 1},
			n:      3,
			output: []int{1, 0},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("input: %v", test.input)
		t.Run(name, func(t *testing.T) {
			Rotate(test.n, test.input)
			if !rev.Equal(test.input, test.output) {
				t.Errorf("got %v, expected: %v", test.input, test.output)
			}
		})
	}
}
