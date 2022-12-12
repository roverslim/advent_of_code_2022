package exercises

import (
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
	head        *knot
	tail        *knot
	tailHistory []*knot
}

func newropeBridge() *ropeBridge {
	return &ropeBridge{
		head:        &knot{x: 0, y: 0},
		tail:        &knot{x: 0, y: 0},
		tailHistory: []*knot{},
	}
}
func (rb *ropeBridge) addTailHistory(k *knot) {
	for _, each_knot := range rb.tailHistory {
		if each_knot.x == k.x && each_knot.y == k.y {
			return
		}
	}

	newKnot := &knot{x: k.x, y: k.y}
	rb.tailHistory = append(rb.tailHistory, newKnot)
}
func (rb *ropeBridge) moveRight(steps int) {
	// fmt.Printf("move right %d\n", steps)
	for i := 0; i < steps; i++ {
		rb.head.x += 1

		if rb.head.x-rb.tail.x >= 2 && int(math.Abs(float64(rb.head.y-rb.tail.y))) >= 1 {
			rb.tail.y = rb.head.y
		}
		if rb.head.y == rb.tail.y && rb.head.x-rb.tail.x >= 2 {
			rb.tail.x = rb.head.x - 1
		}

		// fmt.Printf("\th:%+v t:%+v\n", *rb.head, *rb.tail)
		rb.addTailHistory(rb.tail)
	}
}
func (rb *ropeBridge) moveUp(steps int) {
	// fmt.Printf("move up %d\n", steps)
	for i := 0; i < steps; i++ {
		rb.head.y += 1

		if rb.head.y-rb.tail.y >= 2 && int(math.Abs(float64(rb.head.x-rb.tail.x))) >= 1 {
			rb.tail.x = rb.head.x
		}
		if rb.head.x == rb.tail.x && rb.head.y-rb.tail.y >= 2 {
			rb.tail.y = rb.head.y - 1
		}

		// fmt.Printf("\th:%+v t:%+v\n", *rb.head, *rb.tail)
		rb.addTailHistory(rb.tail)
	}
}
func (rb *ropeBridge) moveLeft(steps int) {
	// fmt.Printf("move left %d\n", steps)
	for i := 0; i < steps; i++ {
		rb.head.x -= 1

		if rb.tail.x-rb.head.x >= 2 && int(math.Abs(float64(rb.head.y-rb.tail.y))) >= 1 {
			rb.tail.y = rb.head.y
		}
		if rb.head.y == rb.tail.y && rb.tail.x-rb.head.x >= 2 {
			rb.tail.x = rb.head.x + 1
		}

		// fmt.Printf("\th:%+v t:%+v\n", *rb.head, *rb.tail)
		rb.addTailHistory(rb.tail)
	}

}
func (rb *ropeBridge) moveDown(steps int) {
	// fmt.Printf("move down %d\n", steps)
	for i := 0; i < steps; i++ {
		rb.head.y -= 1

		if rb.tail.y-rb.head.y >= 2 && int(math.Abs(float64(rb.head.x-rb.tail.x))) >= 1 {
			rb.tail.x = rb.head.x
		}
		if rb.head.x == rb.tail.x && rb.tail.y-rb.head.y >= 2 {
			rb.tail.y = rb.head.y + 1
		}

		// fmt.Printf("\th:%+v t:%+v\n", *rb.head, *rb.tail)
		rb.addTailHistory(rb.tail)
	}
}

func Day9Part1(filepath string) int {
	text := file_reader.Read(filepath)

	rb := newropeBridge()
	for _, each_line := range text {
		parseMotion(each_line, rb)
	}

	return len(rb.tailHistory)
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
