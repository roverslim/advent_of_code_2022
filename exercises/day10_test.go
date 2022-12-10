package exercises

import "testing"

func TestDay10Part1Sample(t *testing.T) {
	expected := 13140
	result := Day10Part1("./day10_sample_input.txt")
	if result != expected {
		t.Errorf("Day10Part1() returns %d, want %d", result, expected)
	}
}

func TestDay10Part1(t *testing.T) {
	expected := 15680
	result := Day10Part1("./day10_input.txt")
	if result != expected {
		t.Errorf("Day10Part1() returns %d, want %d", result, expected)
	}
}
