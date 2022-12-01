package exercises

import (
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
)

type elf struct {
	total    int
	calories []int
}

func newElf() *elf {
	return &elf{total: 0, calories: make([]int, 0)}
}
func (e *elf) updateCalories(food int) {
	e.total += food
	e.calories = append(e.calories, food)
}

type expedition struct {
	elves []*elf
}

func newExpedition() *expedition {
	return &expedition{elves: make([]*elf, 0)}
}
func (e *expedition) addElf() *elf {
	new_elf := newElf()
	e.elves = append(e.elves, new_elf)
	return new_elf
}

func Day1(filepath string) int {
	text := file_reader.Read(filepath)

	expedition := newExpedition()
	current_elf := expedition.addElf()

	for _, each_ln := range text {
		// fmt.Println(each_ln)

		if len(each_ln) > 0 {
			food, _ := strconv.Atoi(each_ln)
			current_elf.updateCalories(food)
		} else {
			current_elf = expedition.addElf()
		}
	}

	// fmt.Printf("%+v\n", *expedition)
	most_calories := 0
	for _, each_elf := range expedition.elves {
		if each_elf.total > most_calories {
			most_calories = each_elf.total
		}
		// fmt.Printf("%+v\n", *each_elf)
	}

	return most_calories
}
