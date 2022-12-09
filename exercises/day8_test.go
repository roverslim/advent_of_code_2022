package exercises

import "testing"

func TestDay8Part1Sample(t *testing.T) {
	expected := 21
	result := Day8Part1("./day8_sample_input.txt")
	if result != expected {
		t.Errorf("Day8Part1() returns %d, want %d", result, expected)
	}
}

func TestDay8Part1(t *testing.T) {
	expected := 1823
	result := Day8Part1("./day8_input.txt")
	if result != expected {
		t.Errorf("Day8Part1() returns %d, want %d", result, expected)
	}
}
