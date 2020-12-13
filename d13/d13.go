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

type line struct {
	start   int
	pos, id int
}

type bus struct {
	edt   int
	lines []line
}

func readInput(inputName string) bus {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	b := bus{}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	l := scanner.Text()
	v, _ := strconv.Atoi(l)
	b.edt = int(v)

	scanner.Scan()
	l = scanner.Text()
	t := strings.Split(l, ",")
	for i := 0; i < len(t); i++ {
		if t[i] == "x" {
			continue
		}
		v, _ := strconv.Atoi(t[i])
		b.lines = append(b.lines, line{pos: i, id: int(v)})
	}

	err = scanner.Err()
	panicOnError(err)

	return b
}

func (b bus) busAt(at int) int {
	for _, line := range b.lines {
		if at%line.id == 0 {
			return line.id
		}
	}

	return -1
}

func part1(b bus) int {
	at, line := b.edt, -1

	for line < 0 {
		at++
		line = b.busAt(at)
	}

	return (at - b.edt) * line
}

func mergeLines(l1, l2 line) line {
	x0, y0, diff := l1.start, l2.start, l2.pos-l1.pos
	visited := make(map[int]int)

	cycle := false
	for !cycle {
		visited[y0-x0] = x0
		if x0 < y0 {
			x0 += l1.id
		} else {
			y0 += l2.id
		}
		_, cycle = visited[y0-x0]
	}

	firstOcc, _ := visited[y0-x0]
	cycleLength := x0 - firstOcc

	return line{start: visited[diff], id: cycleLength, pos: l1.pos}
}

func reduce(b bus) bus {
	nb := bus{}
	for i := 0; i < len(b.lines)-1; i++ {
		nb.lines = append(nb.lines, mergeLines(b.lines[i], b.lines[i+1]))
	}

	return nb
}

func part22(b bus) int {
	for len(b.lines) != 1 {
		b = reduce(b)
	}
	return b.lines[0].start
}

func main() {
	fmt.Println("Hello")
	// b := readInput("small_input.txt")
	b := readInput("input.txt")
	fmt.Printf("%#v\n", part1(b))
	fmt.Printf("%#v\n", part22(b))
}
