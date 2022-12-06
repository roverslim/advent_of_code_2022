package exercises

import (
	file_reader "playground/advent_of_code_2022/helpers"
	"strings"
)

func Day6Part1(filepath string) []int {
	return common(filepath, 4)
}

func Day6Part2(filepath string) []int {
	return common(filepath, 14)
}

func common(filepath string, numDistinctChars int) []int {
	text := file_reader.Read(filepath)

	indexes := []int{}
	for _, each_ln := range text {
		indexes = append(indexes, startOfMarker(each_ln, numDistinctChars))
	}

	return indexes
}

func startOfMarker(signal string, numDistinctChars int) int {
	index := 0

	for i := 0; i < len(signal); i++ {
		marker := signal[index : index+numDistinctChars]

		if !repeatedCharacters(marker) {
			return index + numDistinctChars
		}

		index++
	}

	return index + numDistinctChars
}

func repeatedCharacters(marker string) bool {
	for i := 0; i < len(marker); i++ {
		if strings.Count(marker, marker[i:i+1]) > 1 {
			return true
		}
	}
	return false
}
