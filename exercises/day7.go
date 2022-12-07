package exercises

import (
	"fmt"
	file_reader "playground/advent_of_code_2022/helpers"
	"strconv"
	"strings"
)

type fileSystem struct {
	rootDirectory    *directory
	currentDirectory *directory
}

func newFileSystem() fileSystem {
	rootDir := directory{
		name:      "/",
		parentDir: nil,
	}
	return fileSystem{
		rootDirectory:    &rootDir,
		currentDirectory: &rootDir,
	}
}
func (fs *fileSystem) prettyPrint() {
	fmt.Printf("\t%+v\n", *fs)
	fs.rootDirectory.prettyPrint()
}
func (fs *fileSystem) addDirectory(name string) {
	currentDir := fs.currentDirectory

	newDir := directory{
		name:      name,
		parentDir: currentDir,
	}
	currentDir.directories = append(currentDir.directories, &newDir)
}
func (fs *fileSystem) addFile(name string, size int) {
	currentDir := fs.currentDirectory

	newFile := file{
		name: name,
		size: size,
	}
	currentDir.files = append(currentDir.files, &newFile)
}
func (fs *fileSystem) cdToRoot() {
	fs.currentDirectory = fs.rootDirectory
}
func (fs *fileSystem) cdUp() {
	fs.currentDirectory = fs.currentDirectory.parentDir
}
func (fs *fileSystem) cdInto(dirName string) {
	for _, each_dir := range fs.currentDirectory.directories {
		if each_dir.name == dirName {
			fs.currentDirectory = each_dir
		}
	}
}

type directory struct {
	name        string
	directories []*directory
	files       []*file

	parentDir *directory
}

func (d *directory) totalSize() int {
	totalSize := 0

	for _, each_file := range d.files {
		totalSize += each_file.size
	}

	for _, each_dir := range d.directories {
		totalSize += each_dir.totalSize()
	}

	return totalSize
}
func (d *directory) prettyPrint() {
	fmt.Printf("\tdir %+v\n", *d)
	fmt.Printf("\tsize %d\n", d.totalSize())

	for _, each_file := range d.files {
		each_file.prettyPrint()
	}

	for _, each_dir := range d.directories {
		each_dir.prettyPrint()
	}
}

type file struct {
	name string
	size int
}

func (f *file) prettyPrint() {
	fmt.Printf("\t\tfile %+v\n", *f)
}

func Day7Part1(filepath string) int {
	text := file_reader.Read(filepath)

	fs := newFileSystem()
	for _, each_ln := range text {
		parse(each_ln, &fs)
	}

	fs.prettyPrint()
	return 0
}

func parseCommand(input string, fs *fileSystem) {
	args := strings.Split(input, " ")
	if args[1] == "cd" {
		move(args[2], fs)
	} else {
		fmt.Println("list")
	}
}

func parseOutput(input string, fs *fileSystem) {
	args := strings.Split(input, " ")
	if args[0] == "dir" {
		isDirectory(args[1], fs)
	} else {
		size, _ := strconv.Atoi(args[0])
		isFile(size, args[1], fs)
	}
}

func isDirectory(name string, fs *fileSystem) {
	fs.addDirectory(name)
	fmt.Printf("directory %s\n", name)
}

func isFile(size int, name string, fs *fileSystem) {
	fs.addFile(name, size)
	fmt.Printf("file %s with size %d\n", name, size)
}

func move(input string, fs *fileSystem) {
	if input == ".." {
		fs.cdUp()
		fmt.Println("Move out one level up (cd ..)")
	} else if input == "/" {
		fs.cdToRoot()
		fmt.Println("Move to the outermost directory")
	} else {
		fs.cdInto(input)
		fmt.Printf("Move in one level down (cd %s)\n", input)
	}
}

func parse(input string, fs *fileSystem) {
	if strings.Contains(input, "$") {
		parseCommand(input, fs)
	} else {
		parseOutput(input, fs)
	}
}
