package roman

import (
	"testing"
)

func TestRomanToDecimalConverter(t *testing.T) {
	testCases := []struct {
		roman    string // the roman numeral to convert
		expected int
	}{
		{"I", 1},
		{"MMMMMMMMMDCLXVI", 9666},
		{"L", 50},
		{"MCMXCVI", 1996},
		{"MMMM", 4000},
		{"C", 100},
	}
	// initialize a variable of type Roman
	r := NewRoman()
	for _, tc := range testCases {
		got := r.ToDecimal(tc.roman)
		if got != tc.expected {
			t.Errorf("got: %v, expected: %v", got, tc.expected)
		}
	}
}
