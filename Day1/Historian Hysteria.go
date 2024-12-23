package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reading file")
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var left []int
	var right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Fields(line)
		l, err_l := strconv.Atoi(s[0])
		r, err_r := strconv.Atoi(s[1])
		if err_l != nil {
			panic(err_l)
		}
		if err_r != nil {
			panic(err_r)
		}
		left = Sort(left, l)
		right = Sort(right, r)
	}
	distance := FindDistance(left, right)
	fmt.Printf("Distance: %d\n", distance)

	similarity := FindSimilarity(left, right)
	fmt.Printf("Similarity: %d\n", similarity)
	file.Close()
}

func Sort(intList []int, c int) []int {
	var newList []int
	for i := 0; i < len(intList); i += 1 {
		// fmt.Printf("Comparing %d with %d\n", c, intList[i])
		if c <= intList[i] {
			newList = slices.Insert(intList, i, c)
			// fmt.Printf("Inserting %d at index %d\n", c, i)
			// fmt.Printf("New List: %v\n", newList)
			return newList
		}
	}
	// fmt.Printf("Appending %d at the end\n", c)
	newList = append(intList, c)
	// fmt.Printf("New List: %v\n", newList)
	return newList
}

func FindDistance(left, right []int) int {
	var sum int
	for i := 0; i < len(left); i += 1 {
		sum = sum + int(math.Abs(float64(left[i]-right[i])))
	}
	return sum
}

func FindSimilarity(left, right []int) int {
	var sum int
	for i := 0; i < len(left); i += 1 {
		var occurence int
		for j := 0; j < len(right); j += 1 {

			if left[i] == right[j] {
				occurence++
			}
		}
		sum = sum + (left[i] * occurence)
	}
	return sum
}
