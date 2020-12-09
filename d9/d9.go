package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type series struct {
	preamble int
	val      []int
}

func readInput(inputName string, preamble int) series {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	s := series{preamble: preamble}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		v, err := strconv.Atoi(l)
		panicOnError(err)
		s.val = append(s.val, v)
	}

	err = scanner.Err()
	panicOnError(err)

	return s
}

func (s series) valid(i int) bool {
	for j := i - s.preamble; j < i; j++ {
		for k := j + 1; k < i; k++ {
			if s.val[j]+s.val[k] == s.val[i] {
				return true
			}
		}
	}
	return false
}

func part1(s series) int {
	for i := s.preamble; i < len(s.val); i++ {
		if !s.valid(i) {
			return s.val[i]
		}
	}
	panic("haaaaaaaaaaaaaa")
}

func (s series) sumTo(i, val int) (int, bool) {
	var sum int
	for sum = 0; sum < val && i < len(s.val); i++ {
		sum += s.val[i]
	}
	if sum == val {
		return i, true
	}

	return -1, false
}

func (s series) minMax(i, j int) (int, int) {
	min, max := s.val[i], s.val[i]
	for k := i; k < j; k++ {
		v := s.val[k]
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}

func part2(s series) int {
	sum := part1(s)

	for i := 0; i < len(s.val); i++ {
		if j, ok := s.sumTo(i, sum); ok {
			min, max := s.minMax(i, j)
			return min + max
		}
	}

	return -1
}

func main() {
	fmt.Println("Hello")
	// s := readInput("small_input.txt", 5)
	s := readInput("input.txt", 25)
	fmt.Printf("%#v\n", part1(s))
	fmt.Printf("%#v\n", part2(s))
}
