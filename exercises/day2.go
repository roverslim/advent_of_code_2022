package exercises

import (
	file_reader "playground/advent_of_code_2022/helpers"
	"strings"
)

func playPoints(hand string) int {
	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	return points[hand]
}

// A, X: rock
// B, Y: paper
// C, Z: scissors
//
// A, X > C
// C, Z > B
// B, Y > A

func roundPoints(opponent_hand string, your_hand string) int {
	if your_hand == "X" && opponent_hand == "A" {
		return 3
	} else if your_hand == "Z" && opponent_hand == "C" {
		return 3
	} else if your_hand == "Y" && opponent_hand == "B" {
		return 3
	} else if your_hand == "X" && opponent_hand == "C" {
		return 6
	} else if your_hand == "Z" && opponent_hand == "B" {
		return 6
	} else if your_hand == "Y" && opponent_hand == "A" {
		return 6
	} else {
		return 0
	}
}

func Day2Part1(filepath string) int {
	text := file_reader.Read(filepath)

	total_score := 0

	for _, each_ln := range text {
		split := strings.Split(each_ln, " ")

		opponent_hand := split[0]
		your_hand := split[1]

		total_score += playPoints(your_hand)
		total_score += roundPoints(opponent_hand, your_hand)
	}

	return total_score
}

func roundEnds(guide string) int {
	if guide == "X" {
		return 0
	} else if guide == "Y" {
		return 3
	} else {
		return 6
	}
}

func youPlay(move string) int {
	points := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	return points[move]
}

// A: rock (1)
// B: paper (2)
// C: scissors (3)
//
// A > C
// C > B
// B > A

func youChoose(opponent_hand string, expected_result string) int {
	if expected_result == "Z" { // Needs to win
		if opponent_hand == "A" {
			return youPlay("B")
		} else if opponent_hand == "B" {
			return youPlay("C")
		} else {
			return youPlay("A")
		}
	} else if expected_result == "Y" { // Needs to draw
		if opponent_hand == "A" {
			return youPlay("A")
		} else if opponent_hand == "B" {
			return youPlay("B")
		} else {
			return youPlay("C")
		}
	} else { // Needs to lose
		if opponent_hand == "A" {
			return youPlay("C")
		} else if opponent_hand == "B" {
			return youPlay("A")
		} else {
			return youPlay("B")
		}
	}
}

func Day2Part2(filepath string) int {
	text := file_reader.Read(filepath)

	total_score := 0

	for _, each_ln := range text {
		split := strings.Split(each_ln, " ")

		opponent_hand := split[0]
		expected_result := split[1]

		total_score += youChoose(opponent_hand, expected_result)
		total_score += roundEnds(expected_result)
	}

	return total_score
}
