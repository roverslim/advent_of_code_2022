package exercises

import "testing"

func TestDay3Part1Sample(t *testing.T) {
	expected := 157
	result := Day3Part1("./day3_sample_input.txt")
	if result != expected {
		t.Errorf("Day3Part1() returns %d, want %d", result, expected)
	}
}

func TestDay3Part1(t *testing.T) {
	expected := 7889
	result := Day3Part1("./day3_input.txt")
	if result != expected {
		t.Errorf("Day3Part1() returns %d, want %d", result, expected)
	}
}
