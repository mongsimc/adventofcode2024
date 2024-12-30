package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	aoc "github.com/mongsimc/adventofcode2024/utils"
)

var memo map[string]int

func main() {
	t1 := time.Now()

	lines := aoc.ReadFile("puzzle.input")
	input := strings.Split(lines[0], " ")
	// stones := make([]int, 0)

	// for _, v := range input {
	// 	i, _ := strconv.Atoi(v)
	// 	stones = append(stones, i)
	// }
	answer1 := blink(input, 0, 25)

	answer2 := 0
	memo = make(map[string]int)
	for i := 0; i < len(input); i++ {
		face, _ := strconv.Atoi(input[i])
		answer2 += countStones(face, 0, 75)
	}

	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func countStones(face int, depth int, max int) int {
	stones := 0
	if depth == max {
		return 1
	}

	key := strconv.Itoa(face) + "_" + strconv.Itoa(depth)
	if memo[key] != 0 {
		return memo[key]
	}
	if face == 0 {
		stones = countStones(1, depth+1, max)
		memo[key] = stones
		return stones
	}
	str := strconv.Itoa(face)
	if len(str)%2 == 0 {
		mid := len(str) / 2
		left, _ := strconv.Atoi(str[0:mid])
		right, _ := strconv.Atoi(str[mid:])
		stones = countStones(left, depth+1, max) +
			countStones(right, depth+1, max)
		memo[key] = stones
		return stones
	}
	stones = countStones(face*2024, depth+1, max)
	memo[key] = stones
	return stones
}

func blink(stones []string, count, total int) int {
	if count == total {
		return len(stones)
	}

	newStones := make([]string, 0)

	for _, v := range stones {
		//fmt.Printf("Value: %s\n", v)

		if v == "0" {
			newStones = append(newStones, "1")

		} else if len(v)%2 == 0 {
			s := len(v) / 2
			first := v[0:s]
			second := v[s:]
			fInt, _ := strconv.Atoi(first)
			sInt, _ := strconv.Atoi(second)
			fStr := strconv.Itoa(fInt)
			sStr := strconv.Itoa(sInt)
			newStones = append(newStones, fStr, sStr)
			//fmt.Printf("First: %s\n and Second: %s", fStr, sStr)
		} else {
			n, _ := strconv.Atoi(v)
			nStr := strconv.Itoa(n * 2024)
			newStones = append(newStones, nStr)
		}

	}
	count++
	return blink(newStones, count, total)
}
