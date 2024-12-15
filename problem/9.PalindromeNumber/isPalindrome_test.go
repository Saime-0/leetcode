package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	table := []struct {
		x        int
		expected bool
	}{
		{x: 10, expected: false},
		{x: 101, expected: true},
		{x: 1011, expected: false},
		{x: 10110, expected: false},
		{x: 101101, expected: true},
		{x: -101101, expected: false},
		{x: 1010101, expected: true},
	}
	for i, tt := range table {
		actual := isPalindrome(tt.x)
		if actual != tt.expected {
			t.Errorf("№%d. x:%d, expected:%t, actual:%t", i+1, tt.x, tt.expected, actual)
		}
	}
}
