package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// type PageRule struct {
// 	before int
// 	after  int
// }

type PageOrder struct {
	lefts  []int
	rights []int
}

func main() {


	filename := "input.txt"
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var update []string
	pageMap := make(map[int]PageOrder)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`(\d+)\|(\d+)`)
		reComplete := regexp.MustCompile(`(\d+,)+\d+`)

		matches := re.FindStringSubmatch(line)
		if matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			// p := PageRule{before: x, after: y}
			// pageRules = append(pageRules, p)

			pageRulesX := pageMap[x]
			rightRules := AddRule(pageRulesX.rights, y)
			newXRules := PageOrder{lefts: pageRulesX.lefts, rights: rightRules}
			pageMap[x] = newXRules

			pageRulesY := pageMap[y]
			leftRules := AddRule(pageRulesY.lefts, x)
			newYRules := PageOrder{lefts: leftRules, rights: pageRulesY.rights}
			pageMap[y] = newYRules

		}

		update = append(update, reComplete.FindAllString(line, -1)...)

	}

	//fmt.Printf("Answer1: %v\n", pageRules)
	//fmt.Printf("Answer2: %v\n", update)
	//fmt.Printf("Answer3: %v\n,", pageMap)

	//list := SortPages(pageRules)
	//list.Print()

	var sum, fixedSum int
	newList := make([][]string, 0)
	for _, value := range update {
		pages := strings.Split(value, ",")
		midpoint, wrongList := CheckPageRule(pageMap, pages)
		if midpoint < 0 {
			newList = append(newList, wrongList)
		} else {
			sum = sum + midpoint
		}
	}

	for i := 0; i < len(newList); i++ {
		midpoint, wrongList := CheckPageRule(pageMap, newList[i])
		if midpoint < 0 {
			newList = append(newList, wrongList)
		} else {
			fixedSum = fixedSum + midpoint
		}
	}

	fmt.Printf("Answer1: %d\n", sum)
	fmt.Printf("Answer2: %d\n", fixedSum)
}

func AddRule(slice []int, value int) []int {
	for _, v := range slice {
		if v == value {
			return slice
		}
	}
	return append(slice, value)
}

func Exist(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func CheckPageRule(rules map[int]PageOrder, pages []string) (int, []string) {

	var found bool
	for i := 0; i < len(pages)-1; i++ {
		current, _ := strconv.Atoi(pages[i])
		next, _ := strconv.Atoi(pages[i+1])

		rights := rules[current].rights
		if !Exist(rights, next) {
			found = true
			//SwapThem
			pages[i] = strconv.Itoa(next)
			pages[i+1] = strconv.Itoa(current)
		}
	}

	if found {
		return -1, pages
	}

	mid, _ := strconv.Atoi(pages[len(pages)/2])
	return mid, pages
}
