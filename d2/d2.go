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

type rule struct {
	min, max int
	r        byte
}

type entry struct {
	rule   rule
	pwd    string
	valid  bool
	valid2 bool
}

func buildEntry(s string) entry {
	r := regexp.MustCompile(`([[:alnum:]]*)-([[:alnum:]]*) (.): (.*)`)
	t := r.FindStringSubmatch(s)
	e := entry{}

	v, _ := strconv.Atoi(t[1])
	e.rule.min = v
	v, _ = strconv.Atoi(t[2])
	e.rule.max = v
	e.rule.r = t[3][0]
	e.pwd = t[4]

	e.validate()

	return e
}

func (e *entry) validate() {
	count := strings.Count(e.pwd, string(e.rule.r))
	e.valid = count >= e.rule.min && count <= e.rule.max

	e.valid2 = (e.pwd[e.rule.min-1] == e.rule.r && e.pwd[e.rule.max-1] != e.rule.r) ||
		(e.pwd[e.rule.min-1] != e.rule.r && e.pwd[e.rule.max-1] == e.rule.r)
}

func readInput(inputName string) []entry {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	ret := []entry{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		ret = append(ret, buildEntry(s))
	}

	err = scanner.Err()
	panicOnError(err)

	return ret
}

func part1(entries []entry) {
	valid := 0
	valid2 := 0
	for _, e := range entries {
		if e.valid {
			valid++
		}
		if e.valid2 {
			valid2++
		}
	}
	fmt.Println(valid)
	fmt.Println(valid2)
}

func main() {
	entries := readInput("input.txt")
	part1(entries)
}
