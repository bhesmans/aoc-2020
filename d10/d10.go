package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type adapter struct {
	out int
}

type bag []adapter

type diffs map[int]int

func readInput(inputName string) bag {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	var b bag

	b = append(b, adapter{out: 0}) // the wall

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		v, err := strconv.Atoi(l)
		panicOnError(err)
		b = append(b, adapter{out: v})
	}

	sort.Slice(b, func(i, j int) bool { return b[i].out < b[j].out })

	b = append(b, adapter{out: b[len(b)-1].out + 3}) // the device

	err = scanner.Err()
	panicOnError(err)

	return b
}

func (b bag) getDiffs() diffs {
	ret := diffs{}

	for i := 0; i+1 < len(b); i++ {
		ret[b[i+1].out-b[i].out]++
	}

	return ret
}

func part1(b bag) int {
	diffs := b.getDiffs()
	return diffs[3] * diffs[1]
}

func (b bag) combi(from int) int {
	if from == len(b)-1 {
		return 1
	}

	s := 0
	for i := from + 1; i < from+4 && i < len(b); i++ {
		if b[i].out-b[from].out <= 3 {
			s += b.combi(i)
		}
	}

	return s
}

func (b bag) combi2(from int, memo map[int]int) int {
	if from == len(b)-1 {
		return 1
	}

	s := 0
	for i := from + 1; i < from+4 && i < len(b); i++ {
		if b[i].out-b[from].out <= 3 {
			s += memo[i]
		}
	}

	return s
}

// dumb default implem not good enough with larger input :/
func dummy_part2(b bag) int {
	return b.combi(0)
}

// Second implem, start from the end and memorize the path lengths on the way back
func part2(b bag) int {
	memo := make(map[int]int)
	for i := len(b) - 1; i >= 0; i-- {
		memo[i] = b.combi2(i, memo)
	}

	return memo[0]
}

func main() {
	fmt.Println("Hello")
	// b := readInput("small_input.txt")
	b := readInput("input.txt")
	fmt.Printf("%#v\n", part1(b))
	fmt.Printf("%#v\n", part2(b))
}
