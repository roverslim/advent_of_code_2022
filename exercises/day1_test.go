package exercises

import "testing"

func TestDay1Part1Sample(t *testing.T) {
	expected := 24000
	result := Day1Part1("./day1_sample_input.txt")
	if result != expected {
		t.Errorf("Day1Part1() returns %d, want %d", result, expected)
	}
}

func TestDay1Part1(t *testing.T) {
	expected := 71780
	result := Day1Part1("./day1_input.txt")
	if result != expected {
		t.Errorf("Day1Part1() returns %d, want %d", result, expected)
	}
}

func TestDay1Part2Sample(t *testing.T) {
	expected := 45000
	result := Day1Part2("./day1_sample_input.txt")
	if result != expected {
		t.Errorf("Day1Part2() returns %d, want %d", result, expected)
	}
}

func TestDay1Part2(t *testing.T) {
	expected := 212489
	result := Day1Part2("./day1_input.txt")
	if result != expected {
		t.Errorf("Day1Part2() returns %d, want %d", result, expected)
	}
}
