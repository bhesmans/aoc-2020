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

var reduce_part2 = true

type term interface {
	reduce() number
}

type operation interface {
	execute(a, b term) number
}

type input struct {
	exps []expression
}

type expression struct {
	vals []term
	ops  []operation
}

func (e expression) reduce() number {
	if reduce_part2 {
		return e.reduce2()
	} else {
		return e.reduce1()
	}
}

func (e expression) reduce1() number {
	i := 1
	res := e.vals[0]
	for _, op := range e.ops {
		res = op.execute(res, e.vals[i])
		i++
	}
	return res.(number)
}

func (e expression) reduce2() number {
	i := 1
	nvals := []term{e.vals[0]}
	nops := []operation{}
	for _, op := range e.ops {
		if _, ok := op.(add); ok {
			nvals[len(nvals)-1] = op.execute(nvals[len(nvals)-1], e.vals[i])
		} else {
			nops = append(nops, op)
			nvals = append(nvals, e.vals[i])
		}
		i++
	}

	i = 1
	res := nvals[0]
	for _, op := range nops {
		res = op.execute(res, nvals[i])
		i++
	}

	return res.(number)
}

type number struct {
	n int
}

func (n number) reduce() number {
	return n
}

type add struct{}
type mul struct{}

func (_ add) execute(a, b term) number {
	return number{n: a.reduce().n + b.reduce().n}
}

func (_ mul) execute(a, b term) number {
	return number{n: a.reduce().n * b.reduce().n}
}

func trim(s *string, i int) {
	*s = (*s)[i:]
}

func newExpression(s *string) expression {
	exp := expression{}
	for len(*s) != 0 {
		c := (*s)[0]
		if c == '(' {
			trim(s, 1)
			exp.vals = append(exp.vals, newExpression(s))
		} else if c == ')' {
			trim(s, 1)
			return exp
		} else if c == ' ' {
			trim(s, 1)
		} else if c == '+' {
			trim(s, 1)
			exp.ops = append(exp.ops, add{})
		} else if c == '*' {
			trim(s, 1)
			exp.ops = append(exp.ops, mul{})
		} else {
			re := regexp.MustCompile(`^([0-9]*)(.*)$`)
			t := re.FindStringSubmatch(*s)
			v, _ := strconv.Atoi(t[1])
			exp.vals = append(exp.vals, number{n: v})
			trim(s, len(t[1]))
		}
	}
	return exp
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		i.exps = append(i.exps, newExpression(&l))
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func sum(i input) int {
	sum := 0
	for _, e := range i.exps {
		sum += e.reduce().n
	}
	return sum
}

func part1(i input) int {
	reduce_part2 = false
	return sum(i)
}

func part2(i input) int {
	reduce_part2 = true
	return sum(i)
}

func main() {
	fmt.Println("Hello")
	// i := readInput("small_input.txt")
	i := readInput("input.txt")
	fmt.Printf("%#v\n", part1(i))
	fmt.Printf("%#v\n", part2(i))
}
