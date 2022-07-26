package toposort

import (
	"57/the-go-programming-language/utils"
	"fmt"
	"testing"
)

func TestTopoSort(t *testing.T) {
	tests := []struct {
		prereqs map[string][]string
		order   []string
	}{
		{
			prereqs: map[string][]string{
				"algorithms": {"data structures"},
				"calculus":   {"linear algebra"},
				"compilers": {
					"data structures",
					"formal languages",
					"computer organization",
				},
				"data structures":       {"discrete math"},
				"databases":             {"data structures"},
				"discrete math":         {"intro to programming"},
				"formal languages":      {"discrete math"},
				"networks":              {"operating systems"},
				"operating systems":     {"data structures", "computer organization"},
				"programming languages": {"data structures", "computer organization"},
			},
			order: []string{
				"intro to programming", "discrete math", "data structures", "algorithms",
				"linear algebra", "calculus", "formal languages", "computer organization",
				"compilers", "databases", "operating systems", "networks", "programming languages"},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("prerequisits: %v", test.prereqs)
		t.Run(name, func(t *testing.T) {
			res := TopoSort(test.prereqs)
			if !utils.Equal(res, test.order) {
				t.Errorf("got %v, expected: %v", res, test.order)
			}
		})
	}
}
