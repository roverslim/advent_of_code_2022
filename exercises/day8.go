package exercises

import (
	"fmt"
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
	"strings"
)

func Day8Part1(filepath string) int {
	text := file_reader.Read(filepath)

	grid := newTreeGrid()
	for _, each_ln := range text {
		trees := parseTreeHeights(each_ln)
		grid.addTreeRow(trees)
	}

	grid.prettyPrint()
	grid.visibleTreeCount()

	return 0
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
}

func newTreeGrid() *treeGrid {
	return &treeGrid{
		heights: [][]int{},
	}
}
func (g *treeGrid) addTreeRow(treeHeights []int) {
	g.heights = append(g.heights, treeHeights)
}
func (g *treeGrid) prettyPrint() {
	for _, each_row := range g.heights {
		fmt.Println(each_row)
	}
}
func (g *treeGrid) visibleTreeCount() int {
	// visible := [][]int{}

	i_len := len(g.heights)
	for i := 0; i < i_len; i++ {
		j_len := len(g.heights[i])
		for j := 0; j < j_len; j++ {
			g.isVisible(i, j)
			fmt.Println(g.heights[i][j])
		}
	}

	return 0
}
func (g *treeGrid) isVisible(i int, j int) bool {
	return false
}
