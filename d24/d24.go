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

var dir map[string]tile

type tile struct {
	x, y, z int
}

type floor struct {
	black map[tile]bool
	flips []string
}

func readInput(inputName string) floor {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	f := floor{black: make(map[tile]bool)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		f.flips = append(f.flips, l)
	}

	err = scanner.Err()
	panicOnError(err)

	return f
}

func trim(s *string, i int) {
	*s = (*s)[i:]
}

func nextDir(s *string) string {
	c := string((*s)[0])
	trim(s, 1)
	if c == "e" || c == "w" {
		return c
	} else {
		c2 := string((*s)[0])
		trim(s, 1)
		return c + c2
	}
}

func (t *tile) add(t2 tile) {
	t.x += t2.x
	t.y += t2.y
	t.z += t2.z
}

func (f *floor) _flip(t tile) {
	if !f.black[t] {
		f.black[t] = true
	} else {
		delete(f.black, t)
	}
}

func (f *floor) flip(s string) {
	current := tile{0, 0, 0}
	for len(s) != 0 {
		d := nextDir(&s)
		current.add(dir[d])
	}
	f._flip(current)
}

func (f *floor) flipThemAll() {
	for _, flip := range f.flips {
		f.flip(flip)
	}
}

func (f floor) arround(t tile) int {
	ret := 0
	for _, d := range dir {
		ct := t
		ct.add(d)
		if f.black[ct] {
			ret++
		}
	}
	return ret
}

func (f floor) step() floor {
	nf := floor{black: make(map[tile]bool)}
	for t, _ := range f.black {
		for _, d := range dir {
			ct := t
			ct.add(d)
			black := f.black[ct]
			arround := f.arround(ct)
			if black {
				nf.black[ct] = true
			}
			if black && (arround == 0 || arround > 2) {
				delete(nf.black, ct)
			}
			if !black && arround == 2 {
				nf.black[ct] = true
			}

		}
	}
	return nf
}

func part1(f floor) int {
	f.flipThemAll()
	return len(f.black)
}

func part2(f floor) int {
	f.flipThemAll()
	for i := 0; i < 100; i++ {
		f = f.step()
	}
	return len(f.black)
}

func initDir() {
	dir = make(map[string]tile)
	dir["ne"] = tile{1, 0, -1}
	dir["e"] = tile{1, -1, 0}
	dir["se"] = tile{0, -1, 1}
	dir["sw"] = tile{-1, 0, 1}
	dir["w"] = tile{-1, 1, 0}
	dir["nw"] = tile{0, 1, -1}
}

func main() {
	fmt.Println("Hello")
	initDir()
	// f := readInput("small_input.txt")
	f := readInput("input.txt")
	fmt.Printf("%v\n", part1(f))
	f = readInput("input.txt")
	fmt.Printf("%v\n", part2(f))
}
