package exercises

import (
	"fmt"
	"math"
	file_reader "playground/advent_of_code_2022/helpers"
	"sort"
	"strconv"
	"strings"
)

func Day11Part1(filepath string) int {
	text := file_reader.Read(filepath)

	mim := newMonkeyInTheMiddle()

	batch := []string{}
	for _, each_line := range text {
		if len(each_line) == 0 {
			mim.addMonkey(parseMonkey(batch, mim))
			batch = []string{}
		} else {
			batch = append(batch, each_line)
		}
	}
	mim.addMonkey(parseMonkey(batch, mim))

	// mim.prettyPrint()
	for i := 0; i < 20; i++ {
		// fmt.Printf("Round %d\n", i+1)
		mim.runRound()
		// mim.prettyPrint()
	}

	return mim.monkeyBusiness()
}

type monkeyInTheMiddle struct {
	monkeys    []*monkey
	monkey_map map[int]*monkey
}

func newMonkeyInTheMiddle() *monkeyInTheMiddle {
	return &monkeyInTheMiddle{
		monkeys:    []*monkey{},
		monkey_map: map[int]*monkey{},
	}
}
func (mim *monkeyInTheMiddle) addMonkey(monkey *monkey) {
	mim.monkeys = append(mim.monkeys, monkey)
	mim.monkey_map[monkey.id] = monkey
}
func (mim *monkeyInTheMiddle) runRound() {
	for _, each_monkey := range mim.monkeys {
		each_monkey.takeTurn()
	}
}
func (mim *monkeyInTheMiddle) prettyPrint() {
	fmt.Println()
	for _, each_monkey := range mim.monkeys {
		fmt.Printf("\t%+v\n", *each_monkey)
	}
	fmt.Println("\t----------")
}
func (mim *monkeyInTheMiddle) monkeyBusiness() int {
	counts := []int{}
	for _, each_monkey := range mim.monkeys {
		counts = append(counts, each_monkey.inspectedItems)
	}
	sort.Ints(counts)

	return counts[len(counts)-1] * counts[len(counts)-2]
}

func parseMonkey(lines []string, game *monkeyInTheMiddle) *monkey {
	m := &monkey{
		id:             -1,
		items:          []int{},
		operation:      "",
		factor:         -1,
		divisible_by:   -1,
		if_true:        -1,
		if_false:       -1,
		game:           game,
		inspectedItems: 0,
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
		} else if strings.Contains(each_line, "Operation: new = old + ") {
			each_line = strings.Replace(each_line, "Operation: new = old + ", "", 1)
			each_line = strings.ReplaceAll(each_line, " ", "")
			m.operation = "+"
			factor, _ := strconv.Atoi(each_line)
			m.factor = factor
		} else if strings.Contains(each_line, "Operation: new = old * old") {
			m.operation = "^"
			m.factor = 2
		} else if strings.Contains(each_line, "Operation: new = old * ") {
			each_line = strings.Replace(each_line, "Operation: new = old * ", "", 1)
			each_line = strings.ReplaceAll(each_line, " ", "")
			m.operation = "*"
			factor, _ := strconv.Atoi(each_line)
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
	id             int
	items          []int
	operation      string
	factor         int
	divisible_by   int
	if_true        int
	if_false       int
	game           *monkeyInTheMiddle
	inspectedItems int
}

func (m *monkey) takeTurn() {
	// fmt.Printf("Monkey %d:\n", m.id)
	for _, worryLevel := range m.items {
		m.inspectedItems += 1
		// fmt.Printf("\tMonkey inspects an item with a worry level of %d\n", worryLevel)
		worryLevel = m.worryLevelOnInspection(worryLevel)
		// fmt.Printf("\t\tWorry level is %s by %d to %d\n", m.operation, m.factor, worryLevel)
		worryLevel = worryLevel / 3
		// fmt.Printf("\t\tMonkey gets bored with item. Worry level is divided by 3 to %d\n", worryLevel)
		if worryLevel%m.divisible_by == 0 {
			// fmt.Printf("\t\tCurrent worry level is divisible by %d\n", m.divisible_by)
			// fmt.Printf("\t\tItem with worry level %d is thrown to monkey %d\n", worryLevel, m.if_true)
			m.throwItemToMonkey(m.if_true, worryLevel)
		} else {
			// fmt.Printf("\t\tCurrent worry level is not divisible by %d\n", m.divisible_by)
			// fmt.Printf("\t\tItem with worry level %d is thrown to monkey %d\n", worryLevel, m.if_false)
			m.throwItemToMonkey(m.if_false, worryLevel)
		}
		// m.game.prettyPrint()
	}
	m.items = []int{}
}
func (m *monkey) worryLevelOnInspection(item int) int {
	switch m.operation {
	case "+":
		return item + m.factor
	case "*":
		return item * m.factor
	case "^":
		return int(math.Pow(float64(item), float64(m.factor)))
	}

	return -1
}
func (m *monkey) throwItemToMonkey(id int, item int) {
	m.game.monkey_map[id].items = append(m.game.monkey_map[id].items, item)
}
