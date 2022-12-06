package exercises

import "testing"

func TestDay6Part1Sample(t *testing.T) {
	expected := []int{7, 5, 6, 10, 11}
	result := Day6Part1("./day6_sample_input.txt")
	if len(result) != len(expected) {
		t.Errorf("Day6Part1() returns len()=%d, expect %d", len(result), len(expected))
	}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Day6Part1() returns %d, want %d", r, expected[i])
		}
	}
}

func TestDay6Part1(t *testing.T) {
	expected := []int{1987}
	result := Day6Part1("./day6_input.txt")
	if len(result) != len(expected) {
		t.Errorf("Day6Part2() returns len()=%d, expect %d", len(result), len(expected))
	}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Day6Part2() returns %d, want %d", r, expected[i])
		}
	}
}

func TestDay6Part2Sample(t *testing.T) {
	expected := []int{19, 23, 23, 29, 26}
	result := Day6Part2("./day6_sample_input.txt")
	if len(result) != len(expected) {
		t.Errorf("Day6Part2() returns len()=%d, expect %d", len(result), len(expected))
	}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Day6Part1() returns %d, want %d", r, expected[i])
		}
	}
}

func TestDay6Part2(t *testing.T) {
	expected := []int{3059}
	result := Day6Part2("./day6_input.txt")
	if len(result) != len(expected) {
		t.Errorf("Day6Part2() returns len()=%d, expect %d", len(result), len(expected))
	}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Day6Part2() returns %d, want %d", r, expected[i])
		}
	}
}
