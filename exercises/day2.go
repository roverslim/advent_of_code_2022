package exercises

import (
	file_reader "playground/advent_of_code_2022/helpers"
	"strings"
)

func playPoints(hand string) int {
	points := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	return points[hand]
}

func roundPoints(opponent_hand string, your_hand string) int {
	opponent_points := playPoints(opponent_hand)
	your_points := playPoints(your_hand)

	if opponent_points < your_points {
		return 6
	} else if opponent_points > your_points {
		return 0
	} else {
		return 3
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

		// fmt.Printf(
		// 	"[%s %s] [%d %d] %d + %d = %d [%d]\n",
		// 	opponent_hand,
		// 	your_hand,
		// 	playPoints(your_hand),
		// 	playPoints(opponent_hand),
		// 	playPoints(your_hand),
		// 	roundPoints(opponent_hand, your_hand),
		// 	playPoints(your_hand)+roundPoints(opponent_hand, your_hand),
		// 	total_score,
		// )
	}

	return total_score
}
