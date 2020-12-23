package main

import (
	"fmt"
	"strconv"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type cup struct {
	id   int
	next int
}

type game struct {
	ccup int
	cups map[int]*cup
}

func b2int(b byte) int {
	v, _ := strconv.Atoi(string(b))
	return v
}

func newGame() game {
	return game{cups: make(map[int]*cup)}
}

func (g game) cup(i int) *cup {
	return g.cups[i]
}

func (g game) print() {
	fmt.Printf("(%v): ", g.ccup)

	var ccup *cup
	for ccup = g.cup(g.ccup); ccup.next != g.ccup; ccup = g.cup(ccup.next) {
		fmt.Printf("%v ", ccup.id)
	}
	fmt.Printf("%v\n", ccup.id)
}

func (g game) sol() int {
	ret := 0
	var ccup *cup
	for ccup = g.cup(1); ccup.next != 1; ccup = g.cup(ccup.next) {
		ret *= 10
		ret += g.cup(ccup.next).id
	}
	return ret
}

func (g game) inNextN(x, start, n int) bool {
	c := g.cup(g.cup(start).next)
	for i := 0; i < n; i++ {
		if c.id == x {
			return true
		}
		c = g.cup(c.next)
	}
	return false
}

func (g game) prevDst(i int) int {
	ret := i - 1
	if ret == 0 {
		ret = len(g.cups)
	}
	return ret
}

func (g game) destCup() int {
	dst := g.prevDst(g.ccup)
	for g.inNextN(dst, g.ccup, 3) {
		dst = g.prevDst(dst)
	}
	return dst
}

func (g game) next(start, n int) int {
	c := g.cup(start)
	for i := 0; i < n; i++ {
		c = g.cup(c.next)
	}
	return c.id

}

func (g *game) step() {
	dst := g.destCup()
	newCNext := g.next(g.ccup, 4)
	newDstNext := g.next(g.ccup, 1)
	newPickNext := g.next(dst, 1)

	ccup := g.cup(g.ccup)
	lastPickCup := g.cup(g.next(g.ccup, 3))
	dstCup := g.cup(dst)

	ccup.next = newCNext
	lastPickCup.next = newPickNext
	dstCup.next = newDstNext

	g.ccup = ccup.next
}

func readInput(input string) game {
	g := newGame()
	g.ccup = b2int(input[0])

	var lastCup, firstcup *cup
	for i := 0; i < len(input); i++ {
		ncup := &cup{id: b2int(input[i])}
		if lastCup != nil {
			lastCup.next = ncup.id
		} else {
			firstcup = ncup
		}
		g.cups[ncup.id] = ncup
		lastCup = ncup
	}
	lastCup.next = firstcup.id

	return g
}

func readInput2(input string) game {
	g := newGame()
	g.ccup = b2int(input[0])

	var lastCup, firstcup *cup
	for i := 0; i < len(input); i++ {
		ncup := &cup{id: b2int(input[i])}
		if lastCup != nil {
			lastCup.next = ncup.id
		} else {
			firstcup = ncup
		}
		g.cups[ncup.id] = ncup
		lastCup = ncup
	}

	for i := 10; i <= 1000000; i++ {
		ncup := &cup{id: i}
		lastCup.next = ncup.id
		g.cups[ncup.id] = ncup
		lastCup = ncup
	}
	lastCup.next = firstcup.id

	return g
}

func part1(g game) int {
	for i := 0; i < 100; i++ {
		g.step()
	}
	return g.sol()
}

func part2(g game) int {
	for i := 0; i < 10000000; i++ {
		g.step()
	}

	n1 := g.cup(g.cup(1).next).id
	n2 := g.cup(g.cup(g.cup(1).next).next).id
	return n1 * n2
}

func main() {
	fmt.Println("Hello")
	// g := readInput("389125467")
	g := readInput("685974213")
	fmt.Printf("%#v\n", part1(g))
	// g = readInput2("389125467")
	g = readInput2("685974213")
	fmt.Printf("%v\n", part2(g))
}
