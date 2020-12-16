package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type ticket struct {
	vals []int
}

type minMax struct {
	min, max int
}

type rule struct {
	mm []minMax
}

type notes struct {
	myTicket ticket
	tickets  []ticket
	rules    map[string]rule
}

func s2int(s string) int {
	v, err := strconv.Atoi(s)
	panicOnError(err)
	return int(v)
}

func (r rule) valid(i int) bool {
	for _, mm := range r.mm {
		if i >= mm.min && i <= mm.max {
			return true
		}
	}

	return false
}

func (n notes) valid(i int) bool {
	for _, r := range n.rules {
		if r.valid(i) {
			return true
		}
	}

	return false
}

func (n notes) validTicket(t ticket) (bool, int) {
	for _, val := range t.vals {
		if !n.valid(val) {
			return false, val
		}
	}
	return true, 0
}

func part1(n notes) int {
	sum := 0
	for _, t := range n.tickets {
		_, v := n.validTicket(t)
		sum += v
	}
	return sum
}

func newRule(s string) (rule, string) {
	r := rule{}

	re := regexp.MustCompile(`^([a-z ]*): (.*)$`)
	t := re.FindStringSubmatch(s)
	name := t[1]

	s = s[len(name)+2:]

	re = regexp.MustCompile(`^(([0-9]*)-([0-9]*)( or )?)(.*)$`)
	t = re.FindStringSubmatch(s)
	for t != nil {
		r.mm = append(r.mm, minMax{s2int(t[2]), s2int(t[3])})
		s = s[len(t[1]):]
		t = re.FindStringSubmatch(s)
	}

	return r, name

}

func newTicket(s string) ticket {
	ti := ticket{}
	ta := strings.Split(s, ",")
	for _, v := range ta {
		ti.vals = append(ti.vals, s2int(v))
	}
	return ti
}

func readInput(inputName string) notes {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	n := notes{rules: make(map[string]rule)}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	l := scanner.Text()

	re := regexp.MustCompile(`^(your ticket:)$`)
	t := re.FindStringSubmatch(l)
	for t == nil {
		if l != "" {
			r, name := newRule(l)
			n.rules[name] = r
		}
		scanner.Scan()
		l = scanner.Text()
		t = re.FindStringSubmatch(l)
	}

	scanner.Scan()
	l = scanner.Text()
	n.myTicket = newTicket(l)

	scanner.Scan()
	l = scanner.Text()
	scanner.Scan()
	l = scanner.Text()

	for scanner.Scan() {
		l = scanner.Text()
		n.tickets = append(n.tickets, newTicket(l))
	}

	err = scanner.Err()
	panicOnError(err)

	return n
}

func (n *notes) trimInvalid() {
	valid := []ticket{}
	for _, t := range n.tickets {
		if ok, _ := n.validTicket(t); ok {
			valid = append(valid, t)
		}
	}
	n.tickets = valid
}

func (r rule) validForAllTickets(n notes, i int) bool {
	for _, t := range n.tickets {
		if !r.valid(t.vals[i]) {
			return false
		}
	}
	return true

}

func (n *notes) validRuleFor(i int) (bool, string) {
	ret := ""
	for k, r := range n.rules {
		if r.validForAllTickets(*n, i) {
			if ret != "" {
				// More that one, give up for this round
				return false, ""
			}
			ret = k
		}
	}

	if ret == "" {
		panic("Haaaaaaaaaaaaaa")
	}

	delete(n.rules, ret)
	return true, ret
}

func part2(n notes) int {
	n.trimInvalid()

	i2name := make(map[int]string)

	for i := 0; len(i2name) != len(n.myTicket.vals); i = (i + 1) % len(n.myTicket.vals) {
		if _, ok := i2name[i]; ok {
			continue
		}

		if uniq, rs := n.validRuleFor(i); uniq {
			i2name[i] = rs
		}
	}

	mul := 1
	for k, v := range i2name {
		if strings.HasPrefix(v, "departure") {
			mul *= n.myTicket.vals[k]
		}
	}

	return mul
}

func main() {
	fmt.Println("Hello")
	// n := readInput("small_input.txt")
	// n := readInput("small_input2.txt")
	n := readInput("input.txt")
	fmt.Printf("%#v\n", part1(n))
	fmt.Printf("%#v\n", part2(n))
}
