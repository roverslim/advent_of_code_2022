package exercises

import "testing"

func TestDay9Part1Sample(t *testing.T) {
	expected := 13
	result := Day9Part1("./day9_sample_input.txt")
	if result != expected {
		t.Errorf("Day9Part1() returns %d, want %d", result, expected)
	}
}

func TestDay9Part1(t *testing.T) {
	expected := 5878
	result := Day9Part1("./day9_input.txt")
	if result != expected {
		t.Errorf("Day9Part1() returns %d, want %d", result, expected)
	}
}
