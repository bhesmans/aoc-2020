package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxDim = 4

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type point struct {
	vals [maxDim]int // need to be fixed, slice can not be used as key for map
	dim  int         // Consequence: we need to precise the considered dimension for the point
}

type conway struct {
	mem   map[point]byte
	count int
	dim   int
}

func readInput(inputName string) conway {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	c := conway{mem: make(map[point]byte)}

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		l := scanner.Text()
		for x := 0; x < len(l); x++ {
			if l[x] == '#' {
				p := point{vals: [maxDim]int{x, y, 0, 0}}
				c.mem[p] = l[x]
			}
		}
		y++
	}

	err = scanner.Err()
	panicOnError(err)

	return c
}

func (c conway) active(p point) bool {
	p.dim = 0
	if v, ok := c.mem[p]; ok {
		return v == '#'
	}
	return false
}

func (p1 point) add(p2 point) point {
	p := point{}
	for i := 0; i < maxDim; i++ {
		p.vals[i] = p1.vals[i] + p2.vals[i]
	}
	return p
}

func (c conway) _around(p, d point) int {
	if d.dim == c.dim {
		zero := point{dim: d.dim}
		if d == zero {
			return 0
		}
		if c.active(p.add(d)) {
			return 1
		} else {
			return 0
		}
	}

	sum := 0
	for delta := -1; delta < 2; delta++ {
		cp := d
		cp.vals[cp.dim] = delta
		cp.dim += 1
		sum += c._around(p, cp)
	}

	return sum

}

func (c conway) arround(p point) int {
	return c._around(p, point{})
}

func minf(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (c conway) minMax() (point, point) {
	min, max := point{}, point{}
	for p, _ := range c.mem {
		for i := 0; i < c.dim; i++ {
			min.vals[i] = minf(min.vals[i], p.vals[i])
			max.vals[i] = maxf(max.vals[i], p.vals[i])
		}
	}
	return min, max
}

func (c conway) stepFor(p point) byte {
	active := c.active(p)
	arr := c.arround(p)

	if active {
		if arr == 2 || arr == 3 {
			return '#'
		} else {
			return '.'
		}
	} else {
		if arr == 3 {
			return '#'
		} else {
			return '.'
		}
	}
}

func (c *conway) _step(nc *conway, min, max, p point) {
	if p.dim == c.dim {
		status := c.stepFor(p)
		p.dim = 0
		if status == '#' {
			nc.mem[p] = status
			nc.count++
		}
		return
	}

	for d := min.vals[p.dim] - 1; d <= max.vals[p.dim]+1; d++ {
		cp := p
		cp.vals[p.dim] = d
		cp.dim += 1
		c._step(nc, min, max, cp)
	}
}

func (c conway) step() conway {
	nc := conway{mem: make(map[point]byte), dim: c.dim}
	min, max := c.minMax()

	c._step(&nc, min, max, point{})
	return nc
}

func exec(c conway, step int) int {
	for i := 0; i < step; i++ {
		c = c.step()
	}
	return c.count
}

func part1(c conway) int {
	c.dim = 3
	return exec(c, 6)
}

func part2(c conway) int {
	c.dim = 4
	return exec(c, 6)
}

func main() {
	fmt.Println("Hello")
	// c := readInput("small_input.txt")
	c := readInput("input.txt")
	fmt.Printf("%#v\n", part1(c))
	fmt.Printf("%#v\n", part2(c))
}
