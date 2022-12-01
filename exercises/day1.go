package exercises

import (
	"fmt"
	file_reader "playground/advent_of_code_2022/helpers"
)

func Day1() int {
	text := file_reader.Read("exercises/day1_sample_input.txt")
	// and then a loop iterates through
	// and prints each of the slice values.
	for _, each_ln := range text {
		fmt.Println(each_ln)
	}

	return 1
}
