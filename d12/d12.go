package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	east = iota
	south
	west
	north
)

type ins struct {
	action byte
	value  int
}

type pair struct {
	x, y int
}

type nav struct {
	pair
	wp      pair
	dir     int
	ins     []ins
	current int
}

var dirDelta map[byte]pair

func readInput(inputName string) nav {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	nav := nav{wp: pair{x: 10, y: -1}}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		v, _ := strconv.Atoi(string(l[1:]))
		nav.ins = append(nav.ins, ins{action: l[0], value: v})
	}

	err = scanner.Err()
	panicOnError(err)

	return nav
}
func int2dir(i int) byte {
	switch i {
	case east:
		return 'E'
	case south:
		return 'S'
	case north:
		return 'N'
	case west:
		return 'W'
	}
	panic("haaaaaaaaaaaaa")
}

func (n *nav) navigate() {
	for _, ins := range n.ins {
		a := ins.action

		if a == 'F' {
			a = int2dir(n.dir)
		}

		if d, ok := dirDelta[a]; ok {
			n.x += (ins.value * d.x)
			n.y += (ins.value * d.y)
		} else if a == 'R' {
			n.dir += (ins.value / 90)
			n.dir %= 4
		} else {
			n.dir -= (ins.value / 90)
			n.dir %= 4
			if n.dir < 0 {
				n.dir += 4
			}
		}
	}
}

func (n *nav) navigate2() {
	for _, ins := range n.ins {
		a := ins.action

		if a == 'F' {
			n.x += (n.wp.x * ins.value)
			n.y += (n.wp.y * ins.value)
		} else if d, ok := dirDelta[a]; ok {
			n.wp.x += (ins.value * d.x)
			n.wp.y += (ins.value * d.y)
		} else {
			clock := 1.0
			if a == 'L' {
				clock *= -1.0
			}
			x, y := float64(n.wp.x), float64(n.wp.y)
			rad := clock * float64(ins.value/90) * math.Pi / 2.0
			n.wp.x = int(math.Round(x*math.Cos(rad) - y*math.Sin(rad)))
			n.wp.y = int(math.Round(x*math.Sin(rad) + y*math.Cos(rad)))
		}
	}
}

func initDirDelta() {
	dirDelta = make(map[byte]pair)
	dirDelta['E'] = pair{x: 1, y: 0}
	dirDelta['S'] = pair{x: 0, y: 1}
	dirDelta['W'] = pair{x: -1, y: 0}
	dirDelta['N'] = pair{x: 0, y: -1}
}

func abs(i int) int {
	if i < 0 {
		return -1
	}
	return i
}

func part1(n nav) int {
	n.navigate()
	return abs(n.x) + abs(n.y)
}

func part2(n nav) int {
	n.navigate2()
	return abs(n.x) + abs(n.y)
}

func main() {
	fmt.Println("Hello")
	initDirDelta()
	// n := readInput("small_input.txt")
	n := readInput("input.txt")
	fmt.Printf("%#v\n", part1(n))
	fmt.Printf("%#v\n", part2(n))
}
