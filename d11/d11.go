package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type seats struct {
	rows  []string
	count int
	rule  int
}

const (
	rule1 = iota
	rule2
)

func readInput(inputName string) seats {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	s := seats{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		s.rows = append(s.rows, l)
	}

	err = scanner.Err()
	panicOnError(err)

	return s
}

func (s seats) outOfBonds(x, y int) bool {
	return x < 0 || y < 0 || x >= len(s.rows[0]) || y >= len(s.rows)
}

func (s seats) occupiedDir(x, y, dx, dy int) bool {
	for x, y := x+dx, y+dy; !s.outOfBonds(x, y); x, y = x+dx, y+dy {
		if s.rows[y][x] == 'L' {
			return false
		}
		if s.rows[y][x] == '#' {
			return true
		}
	}

	return false
}

func (s seats) occupied(x, y int) bool {
	if s.outOfBonds(x, y) {
		return false
	}

	return s.rows[y][x] == '#'
}

func (s seats) arround(x, y int) int {
	ret := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			occ := false
			if s.rule == rule1 {
				occ = s.occupied(x+i, y+j)
			} else {
				occ = s.occupiedDir(x, y, i, j)
			}
			if occ {
				ret++
			}
		}
	}
	return ret
}

func (s seats) stepFor(x, y int) byte {
	if s.rows[y][x] == '.' {
		return '.'
	}

	occ := s.occupied(x, y)
	arr := s.arround(x, y)

	max := 0
	if s.rule == rule1 {
		max = 4
	} else {
		max = 5
	}

	if occ && arr >= max {
		return 'L'
	}

	if !occ && arr == 0 {
		return '#'
	}

	return s.rows[y][x]
}

func (s seats) step() seats {
	ns := seats{rule: s.rule}
	for i := range s.rows {
		nr := ""
		for j := range s.rows[i] {
			status := s.stepFor(j, i)
			nr += string(status)
			if status == '#' {
				ns.count++
			}
		}
		ns.rows = append(ns.rows, nr)
	}
	return ns
}

func part1(s seats) int {
	//
	// Use string to compare states, kind of hacky
	//
	// prec := ""
	// cur := fmt.Sprintf("%#v", s)
	// for prec != cur {
	// 	prec = cur
	// 	s = s.step()
	// 	cur = fmt.Sprintf("%#v", s)
	// }
	// return s.count

	// Or use DeepEqual
	prec := seats{}
	cur := s
	for !reflect.DeepEqual(prec, cur) {
		prec = cur
		cur = cur.step()
	}

	return cur.count
}

func part2(s seats) int {
	s.rule = rule2
	return part1(s)
}

func main() {
	fmt.Println("Hello")
	// s := readInput("small_input.txt")
	s := readInput("input.txt")
	fmt.Printf("%#v\n", part1(s))
	fmt.Printf("%#v\n", part2(s))
}
