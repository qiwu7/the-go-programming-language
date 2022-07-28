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

func TestLen(t *testing.T) {
	i := IntSet{}
	tests := []struct {
		item     int
		expected int
	}{
		{
			item:     1,
			expected: 1,
		},
		{
			item:     2,
			expected: 2,
		},
		{
			item:     1,
			expected: 2,
		},
		{
			item:     4,
			expected: 3,
		},
		{
			item:     65,
			expected: 4,
		},
		{
			item:     129,
			expected: 5,
		},
	}

	fmt.Println(i.String())
	for _, test := range tests {
		name := fmt.Sprintf("len after adding %d", test.item)
		t.Run(name, func(t *testing.T) {
			i.Add(test.item)
			res := i.Len()
			if res != test.expected {
				t.Errorf("Has result %d, expected: %d", res, test.expected)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	i := IntSet{}
	i.Add(1)
	i.Add(2)
	i.Add(4)
	i.Add(65)
	i.Add(129)
	tests := []struct {
		item     int
		expected string
	}{
		{
			item:     2,
			expected: "{1 4 65 129}",
		},
		{
			item:     2,
			expected: "{1 4 65 129}",
		},
		{
			item:     2000,
			expected: "{1 4 65 129}",
		},
		{
			item:     129,
			expected: "{1 4 65}",
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("remove %d", test.item)
		t.Run(name, func(t *testing.T) {
			i.Remove(test.item)
			res := i.String()
			if res != test.expected {
				t.Errorf("Has result %s, expected: %s", res, test.expected)
			}
		})
	}
}

func TestClear(t *testing.T) {
	i := IntSet{}
	i.Add(1)
	i.Add(2)
	i.Add(4)
	i.Add(65)
	i.Add(129)

	t.Run("clear set", func(t *testing.T) {
		i.Clear()
		if i.Has(1) {
			t.Error("should not include 1 since cleared")
		}
		if i.Len() != 0 {
			t.Errorf("expect 0 length, got %d", i.Len())
		}
	})
}

func TestCopy(t *testing.T) {
	i := IntSet{}
	i.Add(1)
	i.Add(2)
	i.Add(4)
	i.Add(65)
	i.Add(129)

	i2 := i.Copy()

	t.Run("Copy", func(t *testing.T) {
		i.Remove(1)
		if i.Has(1) {
			t.Error("original set should not have 1")
		}
		if !i2.Has(1) {
			t.Error("copy set should still have 1")
		}
		i2.Add(10)
		if i.Has(10) {
			t.Error("original set should not have 10")
		}
		if !i2.Has(10) {
			t.Error("copy set should have 10")
		}
		i.Add(1000)
		if !i.Has(1000) {
			t.Error("original set should have 1000")
		}
		if i2.Has(1000) {
			t.Error("copy set should not have 1000")
		}

	})
}
