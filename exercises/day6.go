package exercises

import (
	file_reader "playground/advent_of_code_2022/helpers"
	"strings"
)

func Day6Part1(filepath string) []int {
	text := file_reader.Read(filepath)

	indexes := []int{}
	for _, each_ln := range text {
		indexes = append(indexes, startOfPacketMarker(each_ln))
	}

	return indexes
}

func startOfPacketMarker(signal string) int {
	index := 0

	for i := 0; i < len(signal); i++ {
		marker := signal[index : index+4]

		if !repeatedCharacters(marker) {
			return index + 4
		}

		index++
	}

	return index + 4
}

func repeatedCharacters(marker string) bool {
	for i := 0; i < len(marker); i++ {
		if strings.Count(marker, marker[i:i+1]) > 1 {
			return true
		}
	}
	return false
}
