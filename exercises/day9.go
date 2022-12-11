package exercises

import (
	"fmt"
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
	"strings"
)

type knot struct {
	x int
	y int
}
type ropeBridge struct {
	head *knot
	tail *knot
}

func newropeBridge() *ropeBridge {
	return &ropeBridge{
		head: &knot{x: 0, y: 0},
		tail: &knot{x: 0, y: 0},
	}
}
func (rb *ropeBridge) moveRight(steps int) {
	rb.head.x += steps
}
func (rb *ropeBridge) moveUp(steps int) {
	rb.head.y += steps
}
func (rb *ropeBridge) moveLeft(steps int) {
	rb.head.x -= steps
}
func (rb *ropeBridge) moveDown(steps int) {
	rb.head.y -= steps
}

func Day9Part1(filepath string) int {
	text := file_reader.Read(filepath)

	rb := newropeBridge()
	for _, each_line := range text {
		parseMotion(each_line, rb)
	}

	fmt.Printf("%+v\n", rb.head)

	return 0
}

func parseMotion(line string, rb *ropeBridge) {
	split := strings.Split(line, " ")
	direction := split[0]
	steps, _ := strconv.Atoi(split[1])
	if strings.Contains(direction, "R") {
		rb.moveRight(steps)
	} else if strings.Contains(direction, "U") {
		rb.moveUp(steps)
	} else if strings.Contains(direction, "L") {
		rb.moveLeft(steps)
	} else if strings.Contains(direction, "D") {
		rb.moveDown(steps)
	}
}
