package exercises

import (
	"fmt"
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
	"strings"
)

func day8(filepath string) *treeGrid {
	text := file_reader.Read(filepath)

	grid := newTreeGrid()
	for _, each_ln := range text {
		trees := parseTreeHeights(each_ln)
		grid.addTreeRow(trees)
	}

	return grid
}

func Day8Part1(filepath string) int {
	grid := day8(filepath)
	return grid.visibleTreeCount()
}

func Day8Part2(filepath string) int {
	grid := day8(filepath)
	return grid.highestScenicScore()
}

func parseTreeHeights(line string) []int {
	trees := strings.Split(line, "")
	heights := []int{}

	for _, each_tree := range trees {
		each_height, _ := strconv.Atoi(each_tree)
		heights = append(heights, each_height)
	}

	return heights
}

type treeGrid struct {
	heights [][]int
	x_max   int
	y_max   int
}

func newTreeGrid() *treeGrid {
	return &treeGrid{
		heights: [][]int{},
		x_max:   0,
		y_max:   0,
	}
}
func (g *treeGrid) addTreeRow(treeHeights []int) {
	g.heights = append(g.heights, treeHeights)
	g.x_max = len(treeHeights)
	g.y_max = len(g.heights)
}
func (g *treeGrid) prettyPrint() {
	for _, each_row := range g.heights {
		fmt.Println(each_row)
	}
}
func (g *treeGrid) visibleTreeCount() int {
	visible := [][]int{}

	for y := 0; y < len(g.heights); y++ {
		row := []int{}
		for x := 0; x < len(g.heights[y]); x++ {
			if g.isHidden(x, y) {
				row = append(row, 0)
			} else {
				row = append(row, 1)
			}
		}
		visible = append(visible, row)
		// fmt.Println(row)
	}

	count := 0
	for y := 0; y < len(visible); y++ {
		for x := 0; x < len(visible[y]); x++ {
			if visible[y][x] == 1 {
				count += 1
			}
		}
	}

	return count
}
func (g *treeGrid) isHidden(x int, y int) bool {
	if x <= 0 || y <= 0 || x >= g.x_max-1 || y >= g.y_max-1 {
		return false
	}

	treeHeight := g.heights[y][x]
	right := false
	left := false
	up := false
	down := false
	for i := 1; x+i < g.x_max; i++ {
		if g.heights[y][x+i] >= treeHeight {
			right = true
			break
		}
	}
	for i := 1; x-i >= 0; i++ {
		if g.heights[y][x-i] >= treeHeight {
			left = true
			break
		}
	}
	for i := 1; y+i < g.y_max; i++ {
		if g.heights[y+i][x] >= treeHeight {
			down = true
			break
		}
	}
	for i := 1; y-i >= 0; i++ {
		if g.heights[y-i][x] >= treeHeight {
			up = true
			break
		}
	}

	return right && left && up && down
}
func (g *treeGrid) scenicScore(x int, y int) int {
	treeHeight := g.heights[y][x]

	scoreRight := 0
	scoreLeft := 0
	scoreUp := 0
	scoreDown := 0
	for i := 1; x+i < g.x_max; i++ {
		if g.heights[y][x+i] >= treeHeight {
			scoreRight += 1
			break
		} else {
			scoreRight += 1
		}
	}
	for i := 1; x-i >= 0; i++ {
		if g.heights[y][x-i] >= treeHeight {
			scoreLeft += 1
			break
		} else {
			scoreLeft += 1
		}
	}
	for i := 1; y+i < g.y_max; i++ {
		if g.heights[y+i][x] >= treeHeight {
			scoreDown += 1
			break
		} else {
			scoreDown += 1
		}
	}
	for i := 1; y-i >= 0; i++ {
		if g.heights[y-i][x] >= treeHeight {
			scoreUp += 1
			break
		} else {
			scoreUp += 1
		}
	}

	return scoreLeft * scoreRight * scoreUp * scoreDown
}
func (g *treeGrid) highestScenicScore() int {
	scenicScores := [][]int{}

	for y := 0; y < len(g.heights); y++ {
		row := []int{}
		for x := 0; x < len(g.heights[y]); x++ {
			row = append(row, g.scenicScore(x, y))
		}
		scenicScores = append(scenicScores, row)
		// fmt.Println(row)
	}

	highestScore := 0
	for y := 0; y < len(scenicScores); y++ {
		for x := 0; x < len(scenicScores[y]); x++ {
			currentScore := scenicScores[y][x]
			if currentScore > highestScore {
				highestScore = currentScore
			}
		}
	}

	return highestScore
}
