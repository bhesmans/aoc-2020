package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 20201227
const subject = 7

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type input struct {
	pkeys []int
	lsize []int
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		v, _ := strconv.Atoi(l)
		i.pkeys = append(i.pkeys, v)
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func transfrom(s, l int) int {
	t := 1
	for i := 0; i < l; i++ {
		t = (t * s) % mod
	}
	return t
}

func (i *input) _getLoopSize(pk int) {
	lsize := 0
	t := 1
	for t != pk {
		lsize++
		t = (t * subject) % mod
	}
	i.lsize = append(i.lsize, lsize)
}

func (i *input) getLoopSize() {
	for _, pk := range i.pkeys {
		i._getLoopSize(pk)
	}
}

func part1(i input) int {
	i.getLoopSize()
	return transfrom(i.pkeys[0], i.lsize[1])
}

func part2(i input) int {
	return 0
}

func main() {
	fmt.Println("Hello")
	// i := readInput("small_input.txt")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
