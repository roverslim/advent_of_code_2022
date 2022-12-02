package exercises

import "testing"

func TestDay2Part1Sample(t *testing.T) {
	expected := 15
	result := Day2Part1("./day2_sample_input.txt")
	if result != expected {
		t.Errorf("Day2Part1() returns %d, want %d", result, expected)
	}
}
