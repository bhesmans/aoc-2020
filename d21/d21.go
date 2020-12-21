package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type food struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

type input struct {
	foods             []food
	assignedAllergens map[string]bool
	contains          map[string]string
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{assignedAllergens: make(map[string]bool), contains: make(map[string]string)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		t := strings.Split(l, " (contains ")
		f := food{allergens: make(map[string]bool), ingredients: make(map[string]bool)}

		ingredients := strings.Split(t[0], " ")
		allergens := strings.Split(t[1][:len(t[1])-1], ", ")

		for _, al := range allergens {
			i.assignedAllergens[al] = false
			f.allergens[al] = true
		}

		for _, ing := range ingredients {
			f.ingredients[ing] = true
		}
		i.foods = append(i.foods, f)
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func (i input) allAssigned() bool {
	for _, v := range i.assignedAllergens {
		if !v {
			return false
		}
	}
	return true
}

func (i *input) initCandidate(al string) map[string]bool {
	candidate := make(map[string]bool)
	for _, f := range i.foods {
		if f.allergens[al] {
			for al, _ := range f.ingredients {
				candidate[al] = true
			}
		}
	}
	return candidate
}

func trim(s1, s2 map[string]bool) {
	for al, _ := range s1 {
		if _, ok := s2[al]; !ok {
			delete(s1, al)
		}
	}
}

func (i *input) removeIngredient(ing string) {
	for _, f := range i.foods {
		delete(f.ingredients, ing)
	}
}

func (i *input) tryAssign(al string) {
	candidate := i.initCandidate(al)

	for _, f := range i.foods {
		if f.allergens[al] {
			trim(candidate, f.ingredients)
		}
	}

	if len(candidate) == 1 {
		i.assignedAllergens[al] = true
		for ing, _ := range candidate {
			i.removeIngredient(ing)
			i.contains[al] = ing
		}
	}
}

func part1(i *input) int {
	for !i.allAssigned() {
		for al, ok := range i.assignedAllergens {
			if ok {
				continue
			}
			i.tryAssign(al)
		}
	}

	sum := 0
	for _, f := range i.foods {
		sum += len(f.ingredients)
	}

	return sum
}

func part2(i input) string {
	als := []string{}
	for al, _ := range i.contains {
		als = append(als, al)
	}
	sort.Strings(als)

	ret := ""
	for _, al := range als {
		ret += i.contains[al] + ","
	}

	ret = ret[:len(ret)-1]

	return ret
}

func main() {
	fmt.Println("Hello")
	// i := readInput("small_input.txt")
	i := readInput("input.txt")
	fmt.Printf("%#v\n", part1(&i))
	fmt.Printf("%v\n", part2(i))
}
