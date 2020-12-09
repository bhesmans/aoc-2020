package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(inputName string) []int {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	ret := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		panicOnError(err)
		ret = append(ret, val)
	}

	err = scanner.Err()
	panicOnError(err)

	return ret
}

func part2(input []int) (int, int, int) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					return i, j, k
				}
			}
		}
	}
	panic("eara")
}

func part1(input []int) (int, int) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				return i, j
			}
		}
	}
	panic("eara")
}

func main() {
	fmt.Println("Hello")
	input := readInput("input.txt")
	i, j := part1(input)
	fmt.Printf("(%v, %v) [%v, %v] check(%v) mult(%v)\n", i, j, input[i], input[j], input[i]+input[j], input[i]*input[j])
	i, j, k := part2(input)
	fmt.Printf("(%v, %v, %v) [%v, %v, %v] check(%v) mult(%v)\n", i, j, k, input[i], input[j], input[k], input[i]+input[j]+input[k], input[i]*input[j]*input[k])
}
