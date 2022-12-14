package exercises

import "testing"

func TestDay7Part1Sample(t *testing.T) {
	expected := 95437
	result := Day7Part1("./day7_sample_input.txt")
	if result != expected {
		t.Errorf("Day7Part1() returns %d, want %d", result, expected)
	}
}

func TestDay7Part1(t *testing.T) {
	expected := 1555642
	result := Day7Part1("./day7_input.txt")
	if result != expected {
		t.Errorf("Day7Part1() returns %d, want %d", result, expected)
	}
}

func TestDay7Part2Sample(t *testing.T) {
	expected := 24933642
	result := Day7Part2("./day7_sample_input.txt")
	if result != expected {
		t.Errorf("Day7Part2() returns %d, want %d", result, expected)
	}
}

func TestDay7Part2(t *testing.T) {
	expected := 5974547
	result := Day7Part2("./day7_input.txt")
	if result != expected {
		t.Errorf("Day7Part2() returns %d, want %d", result, expected)
	}
}
