package main

import (
	"fmt"
	"time"

	aoc "github.com/mongsimc/adventofcode2024/utils"
)

func main() {
	t1 := time.Now()

	lines := aoc.ReadFile("puzzle.input")

	answer1 := SomeFunction1(lines)
	answer2 := SomeFunction2(lines)

	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func SomeFunction1(input []string) (answer int) {
	return
}

func SomeFunction2(input []string) (answer int) {
	return
}
