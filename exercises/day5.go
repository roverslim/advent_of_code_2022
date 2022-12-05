package exercises

import (
	"fmt"
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
	"strings"
)

func Day5Part1(filepath string) int {
	text := file_reader.Read(filepath)

	stackLines := []string{}
	var supplyStack *supplyStack
	for _, each_ln := range text {
		if len(each_ln) == 0 {
			// The first empty line signifies the end of the drawing of the starting stacks
			index := len(stackLines) - 1
			previousLine := stackLines[index]
			numStacks := parseNumStacks(previousLine)
			parseStacks(stackLines[:index], numStacks)
			supplyStack = newSupplyStack(numStacks)
			break
		}
		stackLines = append(stackLines, each_ln)
	}

	fmt.Printf("%+v\n", *supplyStack)

	return 0
}

func parseStacks(lines []string, numStacks int) {
	for _, each_line := range lines {
		split := strings.Split(each_line, " ")
		count := 0
		for _, each_stack := range split {
			if count >= 3 {
				count = 0
				fmt.Printf("'   '")
			} else if each_stack == "" {
				count += 1
			} else {
				fmt.Printf("'%s'", each_stack)
			}
		}
		fmt.Println()
	}
}

type supplyStack struct {
	numStacks int
}

func newSupplyStack(numStacks int) *supplyStack {
	return &supplyStack{numStacks: numStacks}
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
