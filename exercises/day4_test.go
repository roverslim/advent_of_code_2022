package exercises

import "testing"

func TestDay4Part1Sample(t *testing.T) {
	expected := 2
	result := Day4Part1("./day4_sample_input.txt")
	if result != expected {
		t.Errorf("Day4Part1() returns %d, want %d", result, expected)
	}
}

func TestDay4Part1(t *testing.T) {
	expected := 567
	result := Day4Part1("./day4_input.txt")
	if result != expected {
		t.Errorf("Day4Part1() returns %d, want %d", result, expected)
	}
}

func TestDay4Part4Sample(t *testing.T) {
	expected := 4
	result := Day4Part2("./day4_sample_input.txt")
	if result != expected {
		t.Errorf("Day4Part2() returns %d, want %d", result, expected)
	}
}

func TestDay4Part2(t *testing.T) {
	expected := 907
	result := Day4Part2("./day4_input.txt")
	if result != expected {
		t.Errorf("Day4Part2() returns %d, want %d", result, expected)
	}
}
