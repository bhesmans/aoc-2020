package main

import (
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	start []int
	mem   map[int]int
}

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func s2int(s string) int {
	v, err := strconv.Atoi(s)
	panicOnError(err)
	return int(v)
}

func readInput(input string) game {
	p := game{}

	t := strings.Split(input, ",")
	for _, v := range t {
		p.start = append(p.start, s2int(v))
	}

	return p
}

func (g game) play(max int) int {
	i := 0

	g.mem = make(map[int]int)

	for _, v := range g.start {
		g.mem[v] = i
		i++
	}

	prec := g.start[len(g.start)-1]

	var say int
	for ; i < max; i++ {
		v, ok := g.mem[prec]
		if ok {
			say = i - v - 1
		} else {
			say = 0
		}
		g.mem[prec] = i - 1
		prec = say
	}

	return say
}

func part1(g game) int {
	return g.play(2020)
}

func part2(g game) int {
	return g.play(30000000)
}

func main() {
	fmt.Println("Hello")
	// g := readInput("0,3,6")
	g := readInput("13,16,0,12,15,1")
	fmt.Printf("%#v\n", part1(g))
	fmt.Printf("%#v\n", part2(g))
}
