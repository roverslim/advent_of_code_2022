package exercises

import "testing"

func TestDay1Part1Sample(t *testing.T) {
	expected := 24000
	result := Day1("./day1_sample_input.txt")
	if result != expected {
		t.Errorf("Day1() returns %d, want %d", result, expected)
	}
}

func TestDay1Part1(t *testing.T) {
	expected := 71780
	result := Day1("./day1_input.txt")
	if result != expected {
		t.Errorf("Day1() returns %d, want %d", result, expected)
	}
}
