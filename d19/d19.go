package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type rule interface {
	match(s string, i int, rm rules) (bool, int)
	explore(rm rules) []string
}

type rules map[int]rule

type input struct {
	rules rules
	msgs  []string
}

type leafRule struct {
	c byte
}

type refRule struct {
	r [][]int
}

func (lr leafRule) match(s string, i int, _ rules) (bool, int) {
	if i < len(s) && s[i] == lr.c {
		return true, 1
	} else {
		return false, 0
	}
}

func match(s string, r []int, i int, rm rules) (bool, int) {
	match := 0
	for _, rid := range r {
		ok, m := rm[rid].match(s, match+i, rm)
		if !ok {
			return false, 0
		}
		match += m
	}
	return true, match
}

func (rr refRule) match(s string, i int, rm rules) (bool, int) {
	for _, v := range rr.r {
		ok, m := match(s, v, i, rm)
		if ok {
			return ok, m
		}
	}
	return false, 0
}

func (lr leafRule) explore(rm rules) []string {
	return []string{string(lr.c)}
}

func explore(r []int, rm rules) []string {
	ret := []string{""}
	for _, rid := range r {
		nret := []string{}
		for _, s := range rm[rid].explore(rm) {
			for _, ss := range ret {
				nret = append(nret, ss+s)
			}
		}
		ret = nret
	}
	return ret
}

func (rr refRule) explore(rm rules) []string {
	ret := []string{}
	for _, v := range rr.r {
		ret = append(ret, explore(v, rm)...)
	}
	return ret
}

func newRule(s string) (int, rule) {
	t := strings.Split(s, ": ")
	id, _ := strconv.Atoi(t[0])

	if strings.Contains(t[1], `"`) {
		return id, leafRule{c: t[1][1]}
	}

	refR := refRule{}
	t = strings.Split(t[1], " | ")
	for _, dis := range t {
		conj := []int{}
		tt := strings.Split(dis, " ")
		for _, rr := range tt {
			v, _ := strconv.Atoi(rr)
			conj = append(conj, v)
		}
		refR.r = append(refR.r, conj)
	}

	return id, refR

}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{rules: make(map[int]rule)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			break
		}
		id, rule := newRule(l)
		i.rules[id] = rule
	}

	for scanner.Scan() {
		l := scanner.Text()
		i.msgs = append(i.msgs, l)
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func (i input) match(s string) bool {
	ok, m := i.rules[0].match(s, 0, i.rules)
	if ok {
		return m == len(s)
	}
	return false
}

func part1(i input) int {
	sum := 0
	for _, msg := range i.msgs {
		if i.match(msg) {
			sum++
		}
	}
	return sum
}

func (i input) matchFor(id int) (map[string]bool, int) {
	ret := make(map[string]bool)
	matches := i.rules[id].explore(i.rules)
	for _, s := range matches {
		ret[s] = true
	}
	return ret, len(matches[0])
}

func match_part2(s string, m42, m31 map[string]bool, l42, l31 int) bool {
	c11 := 0
	for len(s) >= l31 && m31[s[len(s)-l31:]] {
		c11++
		s = s[:len(s)-l31]
	}

	// at least once rule 11
	if c11 == 0 {
		return false
	}

	for c11 > 0 {
		if len(s) >= l42 && m42[s[len(s)-l42:]] {
			s = s[:len(s)-l42]
		} else {
			return false
		}
		c11--
	}

	c42 := 0

	for len(s) >= l42 && m42[s[len(s)-l42:]] {
		c42++
		s = s[:len(s)-l42]
	}

	// at least once rule 42
	if c42 == 0 {
		return false
	}

	return len(s) == 0 && c42 != 0
}

func part2(i input) int {
	// root rule is 0: 8 11 and modified rule are
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	// --> we can "hardcode" the scheme (done in match_part2)
	// and precalculate 42 and 31 that don't contains loops

	i.rules[8] = refRule{r: [][]int{{42}, {42, 8}}}
	i.rules[11] = refRule{r: [][]int{{42, 31}, {42, 11, 31}}}
	m42, l42 := i.matchFor(42)
	m31, l31 := i.matchFor(31)

	sum := 0
	for _, msg := range i.msgs {
		if match_part2(msg, m42, m31, l42, l31) {
			sum++
		}
	}

	return sum
}

func main() {
	fmt.Println("Hello")
	// i := readInput("small_input.txt")
	// i := readInput("small_input2.txt")
	i := readInput("input.txt")
	fmt.Printf("%#v\n", part1(i))
	fmt.Printf("%#v\n", part2(i))
}
