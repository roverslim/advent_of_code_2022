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
