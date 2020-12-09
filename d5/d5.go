package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type seat struct {
	srow, scol   string
	row, col, id int
}

func newSeat(s string) seat {
	srow := s[0:7]
	scol := s[7:10]

	srow = strings.ReplaceAll(srow, "F", "0")
	srow = strings.ReplaceAll(srow, "B", "1")
	scol = strings.ReplaceAll(scol, "L", "0")
	scol = strings.ReplaceAll(scol, "R", "1")

	row, _ := strconv.ParseInt(srow, 2, 0)
	col, _ := strconv.ParseInt(scol, 2, 0)

	return seat{
		srow: srow,
		scol: scol,
		row:  int(row),
		col:  int(col),
	}
}

func readInput(inputName string) map[int]seat {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	ret := make(map[int]seat)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		ns := newSeat(s)
		ret[int(ns.row*8+ns.col)] = ns
	}

	err = scanner.Err()
	panicOnError(err)

	return ret
}

func minmax(ss map[int]seat) (int, int) {
	min, max := 99999999, 0
	for id, _ := range ss {
		if id > max {
			max = id
		}
		if id < min {
			min = id
		}
	}
	return min, max
}

func part1(ss map[int]seat) int {
	_, max := minmax(ss)

	return max
}

func part2(ss map[int]seat) int {
	min, max := minmax(ss)

	for i := min; i < max; i++ {
		if _, ok := ss[i]; !ok {
			return i
		}
	}

	return -1
}

func main() {
	ss := readInput("input.txt")
	fmt.Printf("%#v\n", part1(ss))
	fmt.Printf("%#v\n", part2(ss))
}
