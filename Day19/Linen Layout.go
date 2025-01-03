package main

import (
	"fmt"
	"strings"
	"time"

	aoc "github.com/mongsimc/adventofcode2024/utils"
)

var maxTowelLength int

func main() {
	t1 := time.Now()

	lines := aoc.ReadFile("sample.input")

	towels, designs := processInput(lines)

	//fmt.Printf("Towels: %v\n", towels)
	//fmt.Printf("Designs: %v\n", designs)

	answer1 := checkDesigns(towels, designs)
	answer2 := SomeFunction2(lines)

	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func checkDesigns(towels map[string]int, designs []string) (answer int) {
	for _, v := range designs {
		
		fmt.Printf("Value: %v\n", v)
	}
	return
}

func processInput(input []string) (towels map[string]int, designs []string) {
	designs = make([]string, 0)
	towels = make(map[string]int)
	splits := strings.Split(input[0], ", ")
	for i := 2; i < len(input); i++ {
		designs = append(designs, input[i])
	}

	for _, v := range splits {
		if len(v) > maxTowelLength {
			maxTowelLength = len(v)
		}
		towels[v] = 1
		towels[reverse(v)] = 1
	}
	return
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func SomeFunction1(input []string) (answer int) {
	return
}

func SomeFunction2(input []string) (answer int) {
	return
}
