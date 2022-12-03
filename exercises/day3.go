package exercises

import (
	file_reader "playground/advent_of_code_2022/helpers"
	"strings"
)

func Day3Part1(filepath string) int {
	text := file_reader.Read(filepath)

	sum := 0
	for _, each_ln := range text {
		length := len(each_ln)
		half := length / 2

		list1 := each_ln[:half]
		list2 := each_ln[half:]

		item := matchingItem(list1, list2)
		sum += priority(item)
	}

	return sum
}

func matchingItem(list1 string, list2 string) string {
	for _, each_item := range list1 {
		if strings.Contains(list2, string(each_item)) {
			return string(each_item)
		}
	}

	return ""
}

func priority(item string) int {
	item_priority := []rune(item)[0]

	if item_priority >= 97 && item_priority <= 122 { // a-z
		return int(item_priority) - 96
	} else if item_priority >= 65 && item_priority <= 90 { // A-Z
		return int(item_priority) - 38
	}

	return 0
}
