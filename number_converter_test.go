package main

import "testing"

var hexTestData = []struct {
	arg         string
	expectedVal uint32
	hasErr      bool
}{
	{"", 0, false},
	{"0", 0, false},
	{"00000", 0, false},
	{"FF", 0xFF, false},
	{"10", 0x10, false},
	{"0x1A", 0x1A, false},
	{"0xFFF1", 0xFFF1, false},
	{"DEADBEEF", 0xDEADBEEF, false},
	{"0xDEADBEEFDEADBEEF", 0, true},
	{"0b0010", 0, true},
}

var decTestData = []struct {
	arg         string
	expectedVal uint32
	hasErr      bool
}{
	{"", 0, false},
	{"0", 0, false},
	{"00000", 0, false},
	{"123", 123, false},
	{"1239490485", 1239490485, false},
	{"-11123", 0, true},
	{"0xFFF1", 0, true},
	{"214748364712", 0, true},
}

func TestGetFromHex(t *testing.T) {
	for _, tt := range hexTestData {
		out, err := getFromHex(tt.arg)
		if (out != tt.expectedVal) || ((err == nil) == tt.hasErr) {
			t.Errorf("hexTestData(%q) => 0x%X, expected 0x%X. Should have error: %t, actual error: %q", tt.arg, out, tt.expectedVal, tt.hasErr, err)
		}
	}
}

func TestGetFromDec(t *testing.T) {
	for _, tt := range decTestData {
		out, err := getFromDec(tt.arg)
		if (out != tt.expectedVal) || ((err == nil) == tt.hasErr) {
			t.Errorf("hexTestData(%q) => 0x%X, expected 0x%X. Should have error: %t, actual error: %q", tt.arg, out, tt.expectedVal, tt.hasErr, err)
		}
	}
}
func TestGetFromBin(t *testing.T) {
}
