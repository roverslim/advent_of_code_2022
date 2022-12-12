package exercises

import (
	"fmt"
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
	"strings"
)

func Day11Part1(filepath string) int {
	text := file_reader.Read(filepath)

	batch := []string{}
	monkeys := []*monkey{}
	for _, each_line := range text {
		if len(each_line) == 0 {
			monkeys = append(monkeys, parseMonkey(batch))
			batch = []string{}
		} else {
			batch = append(batch, each_line)
		}
	}
	monkeys = append(monkeys, parseMonkey(batch))

	for _, each_monkey := range monkeys {
		fmt.Printf("%+v\n", *each_monkey)
	}

	return 0
}

func parseMonkey(lines []string) *monkey {
	m := &monkey{
		id:           -1,
		items:        []int{},
		operation:    "",
		factor:       -1,
		divisible_by: -1,
		if_true:      -1,
		if_false:     -1,
	}

	for _, each_line := range lines {
		if strings.Contains(each_line, "Monkey ") {
			each_line = strings.Replace(each_line, "Monkey ", "", 1)
			each_line = strings.ReplaceAll(each_line, " ", "")
			each_line = strings.ReplaceAll(each_line, ":", "")
			id, _ := strconv.Atoi(each_line)
			m.id = id
		} else if strings.Contains(each_line, "Starting items: ") {
			each_line = strings.Replace(each_line, "Starting items: ", "", 1)
			each_line = strings.ReplaceAll(each_line, " ", "")
			for _, each_item := range strings.Split(each_line, ",") {
				item, _ := strconv.Atoi(each_item)
				m.items = append(m.items, item)
			}
		} else if strings.Contains(each_line, "Operation: ") {
			each_line = strings.Replace(each_line, "Operation: new = old ", "", 1)
			split := strings.Split(each_line, " ")
			m.operation = split[2]
			factor, _ := strconv.Atoi(split[3])
			m.factor = factor
		} else if strings.Contains(each_line, "Test: ") {
			each_line = strings.Replace(each_line, "Test: divisible by ", "", 1)
			each_line = strings.ReplaceAll(each_line, " ", "")
			divisible_by, _ := strconv.Atoi(each_line)
			m.divisible_by = divisible_by
		} else if strings.Contains(each_line, "If true: ") {
			each_line = strings.Replace(each_line, "If true: throw to monkey", "", 1)
			each_line = strings.ReplaceAll(each_line, " ", "")
			next, _ := strconv.Atoi(each_line)
			m.if_true = next
		} else if strings.Contains(each_line, "If false: ") {
			each_line = strings.Replace(each_line, "If false: throw to monkey", "", 1)
			each_line = strings.ReplaceAll(each_line, " ", "")
			next, _ := strconv.Atoi(each_line)
			m.if_false = next
		}
	}

	return m
}

type monkey struct {
	id           int
	items        []int
	operation    string
	factor       int
	divisible_by int
	if_true      int
	if_false     int
}
