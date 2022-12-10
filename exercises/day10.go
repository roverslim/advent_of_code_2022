package exercises

import (
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
	"strings"
)

func Day10Part1(filepath string) int {
	text := file_reader.Read(filepath)

	pc := newProgramCounter()
	for _, each_ln := range text {
		parseProgram(each_ln, pc)
	}

	return pc.sumOfSignalStrengths()
}

type ProgramCounter struct {
	cycles []int
}

func newProgramCounter() *ProgramCounter {
	return &ProgramCounter{
		cycles: []int{1},
	}
}
func (pc *ProgramCounter) currentValue() int {
	return pc.cycles[len(pc.cycles)-1]
}
func (pc *ProgramCounter) noop() {
	currentValue := pc.currentValue()
	pc.cycles = append(pc.cycles, []int{currentValue}...)
}
func (pc *ProgramCounter) add(value int) {
	currentValue := pc.currentValue()
	pc.cycles = append(pc.cycles, []int{currentValue, currentValue + value}...)
}
func (pc *ProgramCounter) sumOfSignalStrengths() int {
	return 20*pc.cycles[20-1] +
		60*pc.cycles[60-1] +
		100*pc.cycles[100-1] +
		140*pc.cycles[140-1] +
		180*pc.cycles[180-1] +
		220*pc.cycles[220-1]

}

func parseProgram(line string, pc *ProgramCounter) {
	if strings.Contains(line, "noop") {
		pc.noop()
	} else {
		split := strings.Split(line, " ")
		value, _ := strconv.Atoi(split[1])
		pc.add(value)
	}
}
