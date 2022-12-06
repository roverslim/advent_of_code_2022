package exercises

import (
	"fmt"
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
	"strings"
)

func Day5Part1(filepath string) string {
	text := file_reader.Read(filepath)

	stackLines := []string{}
	var supplyStack *supplyStack
	parsedStackLines := false
	for _, each_ln := range text {
		if len(each_ln) == 0 {
			// The first empty line signifies the end of the drawing of the starting stacks
			index := len(stackLines) - 1
			previousLine := stackLines[index]
			numStacks := parseNumStacks(previousLine)
			supplyStack = newSupplyStack(numStacks)
			parseStacks(stackLines[:index], supplyStack)
			parsedStackLines = true
		} else if !parsedStackLines {
			stackLines = append(stackLines, each_ln)
		} else {
			// fmt.Printf("before %+v\n", *supplyStack)
			// fmt.Println(each_ln)
			parseRearrangementProcedure(each_ln, supplyStack)
			// fmt.Printf("after. %+v\n\n", *supplyStack)
		}
	}

	return fmt.Sprint(supplyStack.cratesOnTop())
}

func parseRearrangementProcedure(line string, supplyStack *supplyStack) {
	split := strings.Split(line, " ")

	times, _ := strconv.Atoi(split[1])
	from, _ := strconv.Atoi(split[3])
	to, _ := strconv.Atoi(split[5])

	for i := 0; i < times; i++ {
		crate := supplyStack.pop(from - 1)
		supplyStack.push(crate, to-1)
		// fmt.Printf(".")
	}
	// fmt.Printf("%d %d %d\n", times, from, to)
}

func parseStacks(lines []string, supplyStack *supplyStack) {
	for _, each_line := range lines {
		split := strings.Split(each_line, " ")
		count := 0
		index := 0
		for _, crate := range split {
			if count >= 3 {
				count = 0
				// fmt.Printf("' '")
				index += 1
			} else if crate == "" {
				count += 1
			} else {
				supplyStack.addCrate(crate[1:2], index)
				// fmt.Printf("'%s'", crate[1:2])
				index += 1
			}
		}
	}
}

type supplyStack struct {
	numStacks int
	stacks    map[int][]string
}

func newSupplyStack(numStacks int) *supplyStack {
	return &supplyStack{
		numStacks: numStacks,
		stacks:    map[int][]string{},
	}
}
func (s *supplyStack) addCrate(crate string, stack int) {
	s.stacks[stack] = append(s.stacks[stack], crate)
}
func (s *supplyStack) pop(stack int) string {
	head := s.stacks[stack][0]
	s.stacks[stack] = s.stacks[stack][1:]
	return head
}
func (s *supplyStack) push(crate string, stack int) {
	s.stacks[stack] = append([]string{crate}, s.stacks[stack]...)
}
func (s *supplyStack) cratesOnTop() string {
	top := ""
	for i := 0; i < s.numStacks; i++ {
		each_stack := s.stacks[i]
		// fmt.Println(each_stack)
		if len(each_stack) > 0 {
			top = fmt.Sprintf("%s%s", top, each_stack[0])
		}
	}

	return top
}

func parseNumStacks(line string) int {
	split := strings.Split(line, " ")
	stacks := []int{}
	for _, each_stack := range split {
		if len(each_stack) > 0 {
			column, _ := strconv.Atoi(each_stack)
			stacks = append(stacks, column)
		}
	}

	return len(stacks)
}
