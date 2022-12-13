package exercises

import "testing"

func TestDay11Part1Sample(t *testing.T) {
	expected := 10605
	result := Day11Part1("./day11_sample_input.txt")
	if result != expected {
		t.Errorf("Day11Part1() returns %d, want %d", result, expected)
	}
}

func TestDay11Part1(t *testing.T) {
	expected := 62491
	result := Day11Part1("./day11_input.txt")
	if result != expected {
		t.Errorf("Day11Part1() returns %d, want %d", result, expected)
	}
}
