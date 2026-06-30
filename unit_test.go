package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		worlds []string
		prefix string
	}{
		{
			worlds: []string{"flower", "flow", "flight"},
			prefix: "fl",
		},
		{
			worlds: []string{"dog", "racecar", "car"},
			prefix: "",
		},
	}

	for _, tc := range tests {
		result := longestCommonPrefix(tc.worlds)
		if result != tc.prefix {
			t.Errorf("longestCommonPrefix(%v) = %q; expected %q", tc.worlds, result, tc.prefix)
		}
	}
}
