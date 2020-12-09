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

type validator func(val string) bool

type keyValidator struct {
	key      string
	validate validator
}

var mandatory = []keyValidator{
	{key: "byr", validate: validatebyr},
	{key: "iyr", validate: validateiyr},
	{key: "eyr", validate: validateeyt},
	{key: "hgt", validate: validatehgt},
	{key: "hcl", validate: validatehcl},
	{key: "ecl", validate: validateecl},
	{key: "pid", validate: validatepid},
}

var optional = []string{"cid"}

type passport struct {
	kv     map[string]string
	valid  bool
	valid2 bool
}

func validateYear(val string, min, max int) bool {
	if len(val) != 4 {
		return false
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		return false
	}

	return v >= min && v <= max
}

func validatebyr(val string) bool {
	return validateYear(val, 1920, 2002)
}

func validateiyr(val string) bool {
	return validateYear(val, 2010, 2020)
}

func validateeyt(val string) bool {
	return validateYear(val, 2020, 2030)
}

func validatehgt(val string) bool {
	re := regexp.MustCompile(`^([0-9]*)(cm|in)$`)
	t := re.FindStringSubmatch(val)

	if t == nil {
		return false
	}

	if t[2] != "cm" && t[2] != "in" {
		return false
	}

	v, _ := strconv.Atoi(t[1])

	if t[2] == "cm" && (v < 150 || v > 193) {
		return false
	}

	if t[2] == "in" && (v < 59 || v > 76) {
		return false
	}

	return true
}

func validatehcl(val string) bool {
	re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return re.Match([]byte(val))
}

func validateecl(val string) bool {
	re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	return re.Match([]byte(val))
}

func validatepid(val string) bool {
	re := regexp.MustCompile(`^[0-9]{9}$`)
	return re.Match([]byte(val))
}

func newPassport() passport {
	return passport{
		kv: make(map[string]string),
	}
}

func (p *passport) validate() {
	for _, k := range mandatory {
		if _, ok := p.kv[k.key]; !ok {
			return
		}
	}
	p.valid = true
}

func (p *passport) validate2() {
	for _, k := range mandatory {
		val, ok := p.kv[k.key]
		if !ok || !k.validate(val) {
			return
		}
	}
	p.valid2 = true
}

func readInput(inputName string) []passport {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	ret := []passport{}
	ret = append(ret, newPassport())

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			t := strings.Split(s, " ")
			for i := 0; i < len(t); i++ {
				tt := strings.Split(t[i], ":")
				ret[len(ret)-1].kv[tt[0]] = tt[1]
			}
		} else {
			ret[len(ret)-1].validate()
			ret[len(ret)-1].validate2()
			ret = append(ret, newPassport())
		}
	}
	ret[len(ret)-1].validate()
	ret[len(ret)-1].validate2()

	err = scanner.Err()
	panicOnError(err)

	return ret
}

func part1(ps []passport) (int, int) {
	count, count2 := 0, 0
	for _, p := range ps {
		if p.valid {
			count++
		}
		if p.valid2 {
			count2++
		}
	}
	return count, count2
}

func main() {
	fmt.Println("Hello")
	ps := readInput("input.txt")
	c1, c2 := part1(ps)
	fmt.Printf("%#v\n", c1)
	fmt.Printf("%#v\n", c2)
}
