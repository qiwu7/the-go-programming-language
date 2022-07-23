package dedup

import (
	"fmt"
	"testing"
)

func TestDedup(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{
			input:  []string{"a", "a", "b", "b", "c", "c", "c"},
			output: []string{"a", "b", "c"},
		},
		{
			input:  []string{"a", "b"},
			output: []string{"a", "b"},
		},
		{
			input:  []string{"a", "a"},
			output: []string{"a"},
		},
		{
			input:  []string{"a"},
			output: []string{"a"},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("input: %v", test.input)
		t.Run(name, func(t *testing.T) {
			res := Dedup(test.input)
			if !Equal(res, test.output) {
				t.Errorf("got %v, expected: %v", res, test.output)
			}
		})
	}
}
