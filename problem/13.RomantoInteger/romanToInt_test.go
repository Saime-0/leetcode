package main

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	table := []struct {
		roman    string
		expected int
	}{
		{roman: "III", expected: 3},
		{roman: "IV", expected: 4},
		{roman: "IX", expected: 9},
		{roman: "LVIII", expected: 58},
		{roman: "MCMXCIV", expected: 1994},
	}
	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%d", tt.roman, tt.expected), func(t *testing.T) {
			result := romanToInt(tt.roman)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
