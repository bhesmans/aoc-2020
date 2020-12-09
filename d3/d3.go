package main

import (
	"bufio"
	"fmt"
	"os"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type forest struct {
	lines []string
}

func (f forest) get(x, y int) byte {
	x = x % len(f.lines[0])
	y = y % len(f.lines)
	return f.lines[y][x]
}

func (f forest) countTrees(x0, y0, dx, dy int) int {
	trees := 0
	for y, x := y0, x0; y < len(f.lines); y, x = y+dy, x+dx {
		if f.get(x, y) == '#' {
			trees += 1
		}
	}
	return trees

}

func readInput(inputName string) forest {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	f := forest{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		f.lines = append(f.lines, s)
	}

	err = scanner.Err()
	panicOnError(err)

	return f
}

func main() {
	fmt.Println("Hello")
	f := readInput("input.txt")
	fmt.Println(f.countTrees(0, 0, 3, 1))
	fmt.Println(1 *
		f.countTrees(0, 0, 1, 1) *
		f.countTrees(0, 0, 3, 1) *
		f.countTrees(0, 0, 5, 1) *
		f.countTrees(0, 0, 7, 1) *
		f.countTrees(0, 0, 1, 2))

}
