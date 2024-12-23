package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Reading file")
	filename := "input.txt"
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var sum, sumComplete int
	var isDont bool
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
		reComplete := regexp.MustCompile(`do\(\)|don't\(\)|mul\([0-9]{1,3},[0-9]{1,3}\)`)

		split := re.FindAllString(line, -1)
		splitComplete := reComplete.FindAllString(line, -1)

		for i := range split {
			// fmt.Println(split[i])
			sum = sum + Multiply(split[i])
		}

		for i := range splitComplete {
			x := CheckStr(splitComplete[i])
			switch x {
			case 1:
				isDont = true
			case 2:
				isDont = false
			case 3:
				if isDont {
					continue
				}
				sumComplete = sumComplete + Multiply(splitComplete[i])
			}

		}
	}
	fmt.Printf("Answer: %d and %d", sum, sumComplete)
}

func Multiply(input string) int {

	reInt := regexp.MustCompile(`\d+`)
	matches := reInt.FindAllString(input, -1)
	x, _ := strconv.Atoi(matches[0])
	y, _ := strconv.Atoi(matches[1])

	return x * y
}

func CheckStr(input string) int {
	if input == "don't()" {
		return 1
	} else if input == "do()" {
		return 2
	} else {
		return 3
	}
}

func IsDo(input string) bool {
	if input == "do()" {
		return true
	}
	return false
}
