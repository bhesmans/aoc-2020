package main

import (
	"bufio"
	"fmt"
	"math"
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

type instruction interface {
	execute(p *prog)
	execute2(p *prog)
}

type prog struct {
	mem        map[int]int
	set, reset int
	ins        []instruction
	sr         []setReset
}

type setReset struct {
	set, reset int
}

type setMask struct {
	set, reset int
	sr         []setReset
}

type setMem struct {
	addr, val int
}

func (mask setMask) execute(p *prog) {
	p.set = mask.set
	p.reset = mask.reset
}

func (mask setMask) execute2(p *prog) {
	p.set = mask.set
	p.reset = mask.reset
	p.sr = mask.sr
}

func (mem setMem) execute(p *prog) {
	p.mem[mem.addr] = mem.val
	p.mem[mem.addr] |= p.set
	p.mem[mem.addr] &^= p.reset
}

func (mem setMem) execute2(p *prog) {
	addr := mem.addr
	addr |= p.set
	for _, sr := range p.sr {
		addr |= sr.set
		addr &^= sr.reset
		p.mem[addr] = mem.val
	}
}

func s2int(s string) int {
	v, err := strconv.Atoi(s)
	panicOnError(err)
	return int(v)
}

func bin2int(s string) int {
	v, err := strconv.ParseInt(s, 2, 0)
	panicOnError(err)
	return int(v)
}

func newSetMem(addr, val string) setMem {
	return setMem{
		addr: s2int(addr),
		val:  s2int(val),
	}
}

func getSetReset(mask string, i int) setReset {
	bin := strconv.FormatInt(int64(i), 2)
	set := 0
	reset := 0
	j := len(bin) - 1
	for i := len(mask) - 1; i >= 0; i-- {
		if mask[i] == 'X' {
			var b byte
			if j >= 0 {
				b = bin[j]
				j--
			} else {
				b = '0'
			}

			if b == '0' {
				reset |= (1 << (len(mask) - 1 - i))
			} else {
				set |= (1 << (len(mask) - 1 - i))
			}
		}
	}

	return setReset{set: set, reset: reset}
}

func newSetMask(mask string) setMask {
	set := strings.ReplaceAll(mask, "X", "0")
	reset := strings.ReplaceAll(mask, "1", "X")
	reset = strings.ReplaceAll(reset, "0", "1")
	reset = strings.ReplaceAll(reset, "X", "0")

	ret := setMask{
		set:   bin2int(set),
		reset: bin2int(reset),
	}

	max := int(math.Pow(2., float64(strings.Count(mask, "X"))))
	for i := 0; i < max; i++ {
		ret.sr = append(ret.sr, getSetReset(mask, i))
	}

	return ret
}

func readInput(inputName string) prog {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	p := prog{mem: make(map[int]int)}

	var ins instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		re := regexp.MustCompile(`^mask = ([01X]*)$`)
		t := re.FindStringSubmatch(l)
		if t != nil {
			ins = newSetMask(t[1])
		} else {
			re := regexp.MustCompile(`^mem\[([0-9]*)\] = ([0-9]*)$`)
			t = re.FindStringSubmatch(l)
			ins = newSetMem(t[1], t[2])
		}

		p.ins = append(p.ins, ins)
	}

	err = scanner.Err()
	panicOnError(err)

	return p
}

func (p *prog) execute(part2 bool) {
	for _, ins := range p.ins {
		if part2 {
			ins.execute2(p)
		} else {
			ins.execute(p)
		}
	}
}

func (p prog) sumMem() int {
	sum := 0
	for _, v := range p.mem {
		sum += v
	}
	return sum
}

func part1(p prog) int {
	p.execute(false)
	return p.sumMem()
}

func part2(p prog) int {
	p.execute(true)
	return p.sumMem()
}

func main() {
	fmt.Println("Hello")
	// p := readInput("small_input.txt")
	p := readInput("input.txt")
	fmt.Printf("%#v\n", part1(p))
	p.mem = make(map[int]int)
	fmt.Printf("%#v\n", part2(p))
}
