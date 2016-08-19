package main

import "testing"

var hexTestData = []struct {
	arg      string
	expected int
}{
	{"", 0},
}

func TestGetFromHex(t *testing.T) {
	for _, tt := range hexTestData {
		out := getFromHex(tt.arg)
		if out != tt.expected {
			t.Errorf("hexTestData(%q) => %q, expected %q", tt.arg, out, tt.expected)
		}
	}

}
func TestGetFromDec(t *testing.T) {
}
func TestGetFromBin(t *testing.T) {
}
