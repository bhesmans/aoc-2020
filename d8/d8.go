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

type program struct {
	ins       []instruction
	acc, next int
}

type instruction interface {
	execute(p *program)
}

type nop struct {
	p int
}

type jmp struct {
	p int
}

type acc struct {
	p int
}

type newInstruction func(param string) instruction

var instructions = make(map[string]newInstruction)

func (n nop) execute(p *program) {
	p.next++
}

func (j jmp) execute(p *program) {
	p.next += j.p
}

func (j acc) execute(p *program) {
	p.acc += j.p
	p.next++
}

func newNop(param string) instruction {
	return nop{
		p: getInt(param),
	}
}

func getInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func newJmp(param string) instruction {
	return jmp{
		p: getInt(param),
	}
}

func newAcc(param string) instruction {
	return acc{
		p: getInt(param),
	}
}

func parseLine(s string) instruction {
	re := regexp.MustCompile(`^([a-z]*) ((\+|-){1}[0-9]*)$`)
	t := re.FindStringSubmatch(s)

	return instructions[t[1]](t[2])
}

func readInput(inputName string) program {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	p := program{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		ins := parseLine(s)
		p.ins = append(p.ins, ins)
	}

	err = scanner.Err()
	panicOnError(err)

	return p
}

func initInstructions() {
	instructions["nop"] = newNop
	instructions["jmp"] = newJmp
	instructions["acc"] = newAcc
}

func (p *program) step() {
	p.ins[p.next].execute(p)
}

func switchNop(p program, i int) (program, bool) {
	v := p.ins[i]
	if n, ok := v.(nop); ok {
		p.ins[i] = jmp{p: n.p}
		return p, true
	}

	return program{}, false
}

func switchJmp(p program, i int) (program, bool) {
	v := p.ins[i]
	if j, ok := v.(jmp); ok {
		p.ins[i] = nop{p: j.p}
		return p, true
	}

	return program{}, false
}

func part2(p program) program {
	for i, _ := range p.ins {
		pp, ok := switchNop(p, i)
		if ok {
			pp = part1(pp)
			if pp.next == len(p.ins) {
				return pp
			}
			pp, _ = switchJmp(p, i)
		}
		pp, ok = switchJmp(p, i)
		if ok {
			pp = part1(pp)
			if pp.next == len(p.ins) {
				return pp
			}
			pp, _ = switchNop(p, i)
		}
	}
	return program{}
}

func part1(p program) program {
	visited := make(map[int]bool)

	for v := false; p.next < len(p.ins) && !v; v = visited[p.next] {
		visited[p.next] = true
		p.step()
	}

	return p
}

func main() {
	fmt.Println("Hello")
	initInstructions()
	p := readInput("input.txt")
	fmt.Printf("%#v\n", part1(p).acc)
	fmt.Printf("%#v\n", part2(p).acc)
}
