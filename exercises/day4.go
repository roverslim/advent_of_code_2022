package exercises

import (
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
	"strings"
)

func Day4Part1(filepath string) int {
	text := file_reader.Read(filepath)

	count := 0
	for _, each_ln := range text {
		split := strings.Split(each_ln, ",")
		first_range := strings.Split(split[0], "-")
		second_range := strings.Split(split[1], "-")

		if fullyContainsOther(first_range, second_range) {
			count += 1
		}
	}

	return count
}

func fullyContainsOther(first_range []string, second_range []string) bool {
	first_section_start, _ := strconv.Atoi(first_range[0])
	first_section_end, _ := strconv.Atoi(first_range[1])
	second_section_start, _ := strconv.Atoi(second_range[0])
	second_section_end, _ := strconv.Atoi(second_range[1])

	if (first_section_start <= second_section_start && first_section_end >= second_section_end) ||
		(second_section_start <= first_section_start && second_section_end >= first_section_end) {
		return true
	}

	return false
}
