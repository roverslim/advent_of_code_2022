package exercises

import "testing"

func TestDay1Part1Sample(t *testing.T) {
	expected := 24000
	result := Part1("./day1_sample_input.txt")
	if result != expected {
		t.Errorf("Part1() returns %d, want %d", result, expected)
	}
}

func TestDay1Part1(t *testing.T) {
	expected := 71780
	result := Part1("./day1_input.txt")
	if result != expected {
		t.Errorf("Part1() returns %d, want %d", result, expected)
	}
}

func TestDay1Part2Sample(t *testing.T) {
	expected := 45000
	result := Part2("./day1_sample_input.txt")
	if result != expected {
		t.Errorf("Part2() returns %d, want %d", result, expected)
	}
}
