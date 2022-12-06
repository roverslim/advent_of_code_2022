package exercises

import "testing"

func TestDay5Part1Sample(t *testing.T) {
	expected := "CMZ"
	result := Day5Part1("./day5_sample_input.txt")
	if result != expected {
		t.Errorf("Day5Part1() returns %s, want %s", result, expected)
	}
}

func TestDay5Part1(t *testing.T) {
	expected := "ZBDRNPMVH"
	result := Day5Part1("./day5_input.txt")
	if result != expected {
		t.Errorf("Day5Part1() returns %s, want %s", result, expected)
	}
}
