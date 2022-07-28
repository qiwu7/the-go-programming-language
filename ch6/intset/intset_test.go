package intset

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	i := IntSet{}
	tests := []struct {
		item     int
		expected string
	}{
		{
			item:     1,
			expected: "{1}",
		},
		{
			item:     65,
			expected: "{1 65}",
		},
		{
			item:     4,
			expected: "{1 4 65}",
		},
		{
			item:     4,
			expected: "{1 4 65}",
		},
		{
			item:     129,
			expected: "{1 4 65 129}",
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("add %d", test.item)
		t.Run(name, func(t *testing.T) {
			i.Add(test.item)
			res := i.String()
			if res != test.expected {
				t.Errorf("add result %s, expected: %s", res, test.expected)
			}
		})
	}
}

func TestHas(t *testing.T) {
	i := IntSet{}
	i.Add(1)
	i.Add(2)
	i.Add(65)
	tests := []struct {
		item     int
		expected bool
	}{
		{
			item:     1,
			expected: true,
		},
		{
			item:     2,
			expected: true,
		},
		{
			item:     4,
			expected: false,
		},
		{
			item:     65,
			expected: true,
		},
		{
			item:     129,
			expected: false,
		},
	}

	fmt.Println(i.String())
	for _, test := range tests {
		name := fmt.Sprintf("has %d", test.item)
		t.Run(name, func(t *testing.T) {
			res := i.Has(test.item)
			if res != test.expected {
				t.Errorf("Has result %t, expected: %t", res, test.expected)
			}
		})
	}
}
