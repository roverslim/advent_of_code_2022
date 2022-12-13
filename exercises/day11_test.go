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

func TestDay11Part2Sample(t *testing.T) {
	expected := 2713310158
	result := Day11Part2("./day11_sample_input.txt")
	if result != expected {
		t.Errorf("Day11Part2() returns %d, want %d", result, expected)
	}
}

func TestDay11Part2(t *testing.T) {
	expected := -1
	result := Day11Part2("./day11_input.txt")
	if result != expected {
		t.Errorf("Day11Part2() returns %d, want %d", result, expected)
	}
}
