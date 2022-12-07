package exercises

import (
	"fmt"
	file_reader "playground/advent_of_code_2022/helpers"
	"strings"
)

func Day7Part1(filepath string) int {
	text := file_reader.Read(filepath)

	for _, each_ln := range text {
		parse(each_ln)
	}

	return 0
}

func parseCommand(input string) {
	args := strings.Split(input, " ")
	if args[1] == "cd" {
		move(args[2])
	} else {
		fmt.Println("list")
	}
}

func parseOutput(input string) {
	args := strings.Split(input, " ")
	if args[0] == "dir" {
		directory(args[1])
	} else {
		file(args[0], args[1])
	}
}

func directory(name string) {
	fmt.Printf("directory %s\n", name)
}

func file(size string, name string) {
	fmt.Printf("file %s with size %s\n", name, size)
}

func move(input string) {
	if input == ".." {
		fmt.Println("Move out one level up (cd ..)")
	} else if input == "/" {
		fmt.Println("Move to the outermost directory")
	} else {
		fmt.Printf("Move in one level down (cd %s)\n", input)
	}
}

func parse(input string) {
	if strings.Contains(input, "$") {
		parseCommand(input)
	} else {
		parseOutput(input)
	}
}
