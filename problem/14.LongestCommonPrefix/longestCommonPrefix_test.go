package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	table := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"", ""}, ""},
	}

	for _, tt := range table {
		got := longestCommonPrefix(tt.strs)
		if got != tt.want {
			t.Errorf("%v, want: %v", got, tt.want)
		}
	}
}
