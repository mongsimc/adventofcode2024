package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	aoc "github.com/mongsimc/adventofcode/utils"
)

func main() {
	t1 := time.Now()

	lines := aoc.ReadFile("puzzle.input")

	answer1 := bruteForce(lines, []string{" + ", " * "})
	answer2 := bruteForce(lines, []string{" + ", " * ", " || "})

	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func bruteForce(inputs []string, operands []string) (answer int) {

	for _, input := range inputs {
		var matching bool
		in := strings.Split(input, ":")
		left, _ := strconv.Atoi(in[0])
		right := strings.Split(strings.TrimSpace(in[1]), " ")

		numberList := make([]int, 0)
		for _, v := range right {
			n, _ := strconv.Atoi(v)
			numberList = append(numberList, n)
		}

		equations := generateCombinations(numberList, operands)
		//fmt.Printf("Equations: %v\n", equations)

		for _, e := range equations {
			result := evaluateExpressionLeftToRight(e)

			if result == left {
				matching = true

			}
		}

		if matching {
			answer = answer + left
		}
	}

	return answer
}

// Function to generate all combinations using + or *
func generateCombinations(nums []int, operands []string) []string {
	var result []string
	// Start with the first number
	generateRecursively(nums, "", 0, operands, &result)
	return result
}

// Helper function to generate combinations recursively
func generateRecursively(nums []int, current string, index int, operands []string, result *[]string) {
	// Base case: if we have processed all numbers, add the current expression to the result
	if index == len(nums)-1 {
		*result = append(*result, current+strconv.Itoa(nums[index]))
		return
	}

	for _, v := range operands {
		generateRecursively(nums, current+strconv.Itoa(nums[index])+v, index+1, operands, result)
	}
	// Recursively add '+' and '*' and move to the next number
	// generateRecursively(nums, current+strconv.Itoa(nums[index])+" + ", index+1, result)
	// generateRecursively(nums, current+strconv.Itoa(nums[index])+" * ", index+1, result)
	// generateRecursively(nums, current+strconv.Itoa(nums[index])+" || ", index+1, result)
}

// Function to evaluate the expression left to right
func evaluateExpressionLeftToRight(expr string) int {
	//fmt.Printf("Expression: %s\n", expr)
	tokens := strings.Split(expr, " ")

	// Initialize the result with the first number
	result, _ := strconv.Atoi(string(tokens[0]))

	// Iterate over the expression from left to right
	for i := 1; i < len(tokens); i += 2 {
		operator := string(tokens[i])                   // The operator (+ or *)
		operand, _ := strconv.Atoi(string(tokens[i+1])) // The next number

		// Apply the operator to the result and operand
		if operator == "+" {
			result += operand
		} else if operator == "*" {
			result *= operand
		} else if operator == "||" {
			concat := strconv.Itoa(result) + strconv.Itoa(operand)
			result, _ = strconv.Atoi(concat)
		}
	}

	return result
}
