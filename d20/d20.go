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

type tile struct {
	id   int
	rows [][]byte
}

var sides []int

var monster tile

const (
	north = iota
	east
	south
	west
)

type input struct {
	tiles map[int]tile
}

func (t tile) side(s int) []byte {
	if s == north {
		return t.rows[0]
	} else if s == south {
		return t.rows[len(t.rows)-1]
	} else {
		j := 0
		if s == east {
			j = len(t.rows[0]) - 1
		}

		ret := []byte{}

		for _, r := range t.rows {
			ret = append(ret, r[j])
		}

		return ret
	}
}

func sideMatch(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func (t *tile) flip() {
	nt := tile{id: t.id}
	for i := len(t.rows) - 1; i >= 0; i-- {
		nt.rows = append(nt.rows, t.rows[i])
	}
	t.rows = nt.rows
}

func (t tile) print() {
	fmt.Printf("--------------------- %v (%v x %v)\n", t.id, len(t.rows[0]), len(t.rows))
	for _, r := range t.rows {
		fmt.Printf("%v\n", string(r))
	}
	fmt.Printf("---------------------\n")
}

func (t *tile) rotate() {
	lenB := len(t.rows)

	nt := tile{id: t.id}
	for _ = range t.rows {
		nt.rows = append(nt.rows, make([]byte, lenB))
	}

	for j := 0; j < lenB/2; j++ {
		for i := j; i < len(t.rows[0])-j; i++ {
			nt.rows[j][i] = t.rows[lenB-1-i][j]
			nt.rows[i][lenB-1-j] = t.rows[j][i]
			nt.rows[lenB-1-j][lenB-1-i] = t.rows[i][lenB-1-j]
			nt.rows[lenB-1-i][j] = t.rows[lenB-1-j][lenB-1-i]
		}
	}

	t.rows = [][]byte{}
	for _, r := range nt.rows {
		t.rows = append(t.rows, r)
	}
}

func (t1 tile) match_side(t2 *tile, side int) bool {
	t1s := t1.side(side)

	for _, _ = range sides {
		t2.rotate()
		if sideMatch(t1s, t2.side((side+2)%4)) {
			return true
		}
	}

	return false

}

func (t1 tile) match_noflip(t2 *tile) (bool, int) {
	for _, s := range sides {
		if t1.match_side(t2, s) {
			return true, s
		}
	}
	return false, -1
}

func (t1 tile) match(t2 *tile) (bool, int) {
	if ok, s := t1.match_noflip(t2); ok {
		return ok, s
	}

	t2.flip()
	return t1.match_noflip(t2)
}

func (t tile) get(x, y int) byte {
	return t.rows[y][x]
}

func (t tile) countdash() int {
	count := 0
	for _, row := range t.rows {
		for _, c := range row {
			if c == '#' {
				count++
			}
		}
	}
	return count
}

func (t tile) correspond(i, j int, pattern tile) bool {
	for y, row := range pattern.rows {
		for x, c := range row {
			if c == '#' && t.get(i+x, j+y) != '#' {
				return false
			}
		}
	}
	return true
}

func (t tile) findMonster() int {
	count := 0
	for j := 0; j < len(t.rows)-len(monster.rows); j++ {
		for i := 0; i < len(t.rows[j])-len(monster.rows[0]); i++ {
			if t.correspond(i, j, monster) {
				count++
			}
		}
	}

	return count
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{tiles: make(map[int]tile)}

	newTile := true
	cTile := tile{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if newTile {
			newTile = false
			v, _ := strconv.Atoi(l[5 : len(l)-1])
			cTile = tile{id: v}
		} else if l == "" {
			i.tiles[cTile.id] = cTile
			newTile = true
			continue
		} else {
			cTile.rows = append(cTile.rows, []byte(l))
		}
	}
	i.tiles[cTile.id] = cTile

	err = scanner.Err()
	panicOnError(err)

	return i
}

type sol struct {
	ids [][]tile
}

func solve(in *input, topLeft tile) sol {
	sol := sol{}
	bsize := int(math.Round(math.Sqrt(float64(len(in.tiles)))))
	for i := 0; i < bsize; i++ {
		sol.ids = append(sol.ids, make([]tile, bsize))
	}

	sol.ids[0][0] = topLeft

	delete(in.tiles, topLeft.id)
	for i := 0; i < bsize; i++ {
		maxSol := 0
		for _, t := range in.tiles {
			if ok, s := sol.ids[i][0].match(&t); ok && s == south {
				sol.ids[i+1][0] = t
				delete(in.tiles, t.id)
				maxSol++
			}
		}
		if maxSol > 1 {
			panic("Haaaaaaaaaaaa")
		}
	}

	for j := 0; j < bsize; j++ {
		for i := 0; i < bsize-1; i++ {
			maxSol := 0
			for _, t := range in.tiles {
				if ok, s := sol.ids[j][i].match(&t); ok && s == east {
					sol.ids[j][i+1] = t
					delete(in.tiles, t.id)
					maxSol++
				}
			}
			if maxSol > 1 {
				panic("Haaaaaaaaaaaa")
			}
		}
	}

	return sol
}

func flatten(sol sol, in input) tile {
	tileSize := len(sol.ids[0][0].rows) - 2
	size := len(sol.ids) * tileSize
	ret := tile{rows: make([][]byte, size)}

	i := 0
	for _, row := range sol.ids {
		for _, col := range row {
			for _, r := range col.rows[1 : len(col.rows)-1] {
				ret.rows[i] = append(ret.rows[i], r[1:len(r)-1]...)
				i++
			}
			i -= tileSize
		}
		i += tileSize
	}

	return ret

}

func monsterCount(tileSol tile) int {
	for i := 0; i < 4; i++ {
		tileSol.rotate()
		count := tileSol.findMonster()
		if count != 0 {
			return count
		}
	}
	tileSol.flip()
	for i := 0; i < 4; i++ {
		tileSol.rotate()
		count := tileSol.findMonster()
		if count != 0 {
			return count
		}
	}
	panic("NNOOOOO")
}

func part1(i input) (int, int) {

	match := make(map[int][]tile)
	for _, t1 := range i.tiles {
		for _, t2 := range i.tiles {
			if t1.id == t2.id {
				continue
			}
			if ok, _ := t1.match(&t2); ok {
				match[t1.id] = append(match[t1.id], t2)
			}
		}
	}

	mult := 1
	topLeft := tile{}
	for k, v := range match {
		if len(v) == 2 {
			mult *= k
			_, s1 := i.tiles[k].match(&v[0])
			_, s2 := i.tiles[k].match(&v[1])
			if (s1 == east && s2 == south) || (s2 == east && s1 == south) {
				topLeft = i.tiles[k]
			}
		}
	}

	sol := solve(&i, topLeft)
	tileSole := flatten(sol, i)
	monsterCount := monsterCount(tileSole)

	part2 := tileSole.countdash() - monsterCount*monster.countdash()

	return mult, part2
}

func initSides() {
	sides = append(sides, north)
	sides = append(sides, east)
	sides = append(sides, south)
	sides = append(sides, west)

	m1 := []byte("                  # ")
	m2 := []byte("#    ##    ##    ###")
	m3 := []byte(" #  #  #  #  #  #   ")
	monster = tile{rows: [][]byte{m1, m2, m3}}
}

func main() {
	fmt.Println("Hello")
	initSides()
	// i := readInput("small_input.txt")
	i := readInput("input.txt")
	p1, p2 := part1(i)
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)
}
