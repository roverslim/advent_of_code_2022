package exercises

import "testing"

func TestDay1(t *testing.T) {
	result := Day1()
	if result != 1 {
		t.Errorf("Day1() returns %q, want %q", result, 1)
	}
}
