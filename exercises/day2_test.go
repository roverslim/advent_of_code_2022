package exercises

import "testing"

func TestDay2Part1Sample(t *testing.T) {
	expected := 15
	result := Day2Part1("./day2_sample_input.txt")
	if result != expected {
		t.Errorf("Day2Part1() returns %d, want %d", result, expected)
	}
}

func TestDay2Part1(t *testing.T) {
	expected := 15632
	result := Day2Part1("./day2_input.txt")
	if result != expected {
		t.Errorf("Day2Part1() returns %d, want %d", result, expected)
	}
}

func TestDay2Part2Sample(t *testing.T) {
	expected := 12
	result := Day2Part2("./day2_sample_input.txt")
	if result != expected {
		t.Errorf("Day2Part2() returns %d, want %d", result, expected)
	}
}

func TestDay2Part2(t *testing.T) {
	expected := 14416
	result := Day2Part2("./day2_input.txt")
	if result != expected {
		t.Errorf("Day2Part2() returns %d, want %d", result, expected)
	}
}
