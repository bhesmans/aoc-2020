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

type group struct {
	q      map[byte]int
	size   int
	allYes int
}

func newGroup() group {
	return group{
		q: make(map[byte]int),
	}
}

func (g *group) validate() {
	for _, v := range g.q {
		if v == g.size {
			g.allYes++
		}
	}
}

func readInput(inputName string) []group {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	ret := []group{}
	ret = append(ret, newGroup())

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			for i := 0; i < len(s); i++ {
				ret[len(ret)-1].q[s[i]]++
			}
			ret[len(ret)-1].size++
		} else {
			ret[len(ret)-1].validate()
			ret = append(ret, newGroup())
		}
	}
	ret[len(ret)-1].validate()

	err = scanner.Err()
	panicOnError(err)

	return ret
}

func part1(gg []group) int {
	count := 0
	for _, g := range gg {
		count += len(g.q)
	}
	return count
}

func part2(gg []group) int {
	count := 0
	for _, g := range gg {
		count += g.allYes
	}
	return count
}

func main() {
	fmt.Println("Hello")
	gg := readInput("input.txt")
	c := part1(gg)
	fmt.Printf("%#v\n", c)
	c = part2(gg)
	fmt.Printf("%#v\n", c)
}
