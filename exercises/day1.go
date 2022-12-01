package exercises

import (
	file_reader "playground/advent_of_code_2022/helpers"
	"sort"
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

type By func(e1, e2 *elf) bool

func (by By) Sort(elves []*elf) {
	s := &elfSorter{
		elves: elves,
		by:    by,
	}
	sort.Sort(s)
}

type elfSorter struct {
	elves []*elf
	by    func(e1, e2 *elf) bool
}

func (s *elfSorter) Len() int {
	return len(s.elves)
}
func (s *elfSorter) Swap(i, j int) {
	s.elves[i], s.elves[j] = s.elves[j], s.elves[i]
}
func (s *elfSorter) Less(i, j int) bool {
	return s.by(s.elves[i], s.elves[j])
}

func day1(filepath string) *expedition {
	text := file_reader.Read(filepath)

	expedition := newExpedition()
	current_elf := expedition.addElf()

	for _, each_ln := range text {
		if len(each_ln) > 0 {
			food, _ := strconv.Atoi(each_ln)
			current_elf.updateCalories(food)
		} else {
			current_elf = expedition.addElf()
		}
	}

	return expedition
}

func Part1(filepath string) int {
	expedition := day1(filepath)

	totalDecending := func(e1, e2 *elf) bool {
		return e1.total > e2.total
	}
	By(totalDecending).Sort(expedition.elves)

	total := 0
	count := 0
	for _, each_elf := range expedition.elves {
		if count >= 1 {
			break
		}

		total += each_elf.total
		count++
	}

	return total
}

func Part2(filepath string) int {
	expedition := day1(filepath)

	totalDecending := func(e1, e2 *elf) bool {
		return e1.total > e2.total
	}
	By(totalDecending).Sort(expedition.elves)

	total := 0
	count := 0
	for _, each_elf := range expedition.elves {
		if count >= 3 {
			break
		}

		total += each_elf.total
		count++
	}

	return total
}
