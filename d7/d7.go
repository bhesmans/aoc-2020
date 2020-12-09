package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type bag struct {
	color     string
	ncontains map[string]int
}

func newBag(s string) bag {
	re := regexp.MustCompile(`^(.*)( bags contain )(.*)$`)
	t := re.FindStringSubmatch(s)

	nb := bag{
		color:     t[1],
		ncontains: make(map[string]int),
	}
	s = s[len(t[1])+len(" bags contain "):]

	re = regexp.MustCompile(`^(no other bags.)$`)
	t = re.FindStringSubmatch(s)
	if t != nil {
		return nb
	}

	re = regexp.MustCompile(`^([0-9]*)([a-z ]*)( bags?(, |.))(.*)$`)
	t = re.FindStringSubmatch(s)
	for t != nil {
		c, _ := strconv.Atoi(t[1])
		nb.ncontains[t[2][1:]] = c

		s = s[len(t[1])+len(t[2])+len(t[3]):]
		t = re.FindStringSubmatch(s)
	}

	return nb
}

func (b bag) contains(bags map[string]bag, c string) bool {
	for sbag, _ := range b.ncontains {
		if sbag == c || bags[sbag].contains(bags, c) {
			return true
		}
	}
	return false
}

func whocontains(bags map[string]bag, c string) []bag {
	ret := []bag{}
	for _, b := range bags {
		if b.contains(bags, c) {
			ret = append(ret, b)
		}
	}

	return ret
}

func (b bag) sum(bags map[string]bag) int {
	ret := 1
	for sbag, v := range b.ncontains {
		ret += (v * bags[sbag].sum(bags))
	}
	return ret
}

func readInput(inputName string) map[string]bag {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	ret := make(map[string]bag)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		b := newBag(s)
		ret[b.color] = b
	}

	err = scanner.Err()
	panicOnError(err)

	return ret
}

func main() {
	fmt.Println("Hello")
	bb := readInput("input.txt")
	fmt.Printf("%#v\n", len(whocontains(bb, "shiny gold")))
	fmt.Printf("%#v\n", bb["shiny gold"].sum(bb)-1)
}
