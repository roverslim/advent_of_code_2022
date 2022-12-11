package exercises

import (
	"fmt"
	"math"
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
	for i := 0; i < steps; i++ {
		rb.head.x += 1

		if rb.head.x-rb.tail.x >= 2 && int(math.Abs(float64(rb.head.y-rb.tail.y))) >= 1 {
			rb.tail.y = rb.head.y
		}
		rb.tail.x = rb.head.x - 1
	}
}
func (rb *ropeBridge) moveUp(steps int) {
	for i := 0; i < steps; i++ {
		rb.head.y += 1

		if rb.head.y-rb.tail.y >= 2 && int(math.Abs(float64(rb.head.x-rb.tail.x))) >= 1 {
			rb.tail.x = rb.head.x
		}
		rb.tail.y = rb.head.y - 1
	}
}
func (rb *ropeBridge) moveLeft(steps int) {
	for i := 0; i < steps; i++ {
		rb.head.x -= 1

		if rb.tail.x-rb.head.x >= 2 && int(math.Abs(float64(rb.head.y-rb.tail.y))) >= 1 {
			rb.tail.y = rb.head.y
		}
		rb.tail.x = rb.head.x + 1
	}

}
func (rb *ropeBridge) moveDown(steps int) {
	for i := 0; i < steps; i++ {
		rb.head.y -= 1

		if rb.tail.y-rb.head.y >= 2 && int(math.Abs(float64(rb.head.x-rb.tail.x))) >= 1 {
			rb.tail.x = rb.head.x
		}
	}
}

func Day9Part1(filepath string) int {
	text := file_reader.Read(filepath)

	rb := newropeBridge()
	for _, each_line := range text {
		parseMotion(each_line, rb)

		fmt.Printf("%+v ", rb.head)
		fmt.Printf("%+v\n", rb.tail)
	}

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
