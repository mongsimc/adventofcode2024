package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	aoc "github.com/mongsimc/adventofcode2024/utils"
)

var maxX = 101
var maxY = 103
var test = 11000

// var maxX = 11
// var maxY = 7

type Robot struct {
	pX int
	pY int
	vX int
	vY int
}

func main() {
	t1 := time.Now()

	lines := aoc.ReadFile("puzzle.input")

	re := regexp.MustCompile(`p=(\d{1,3}),(\d{1,3}) v=([-]{0,1}\d{1,3}),([-]{0,1}\d{1,3})`)

	robots := make([]Robot, 0)

	for _, v := range lines {
		split := re.FindStringSubmatch(v)
		px, _ := strconv.Atoi(split[1])
		py, _ := strconv.Atoi(split[2])
		vx, _ := strconv.Atoi(split[3])
		vy, _ := strconv.Atoi(split[4])
		robot := Robot{pX: px, pY: py, vX: vx, vY: vy}
		robots = append(robots, robot)
	}

	//robots = reposition(robots, 17000)
	robots = reposition(robots, test)

	//answer1 := calculateSafetyFactor(robots)
	//answer2 := SomeFunction2(lines)

	//fmt.Printf("Answer1: %d\n", answer1)
	//fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func calculateSafetyFactor(input []Robot) (answer int) {
	var q1, q2, q3, q4 int
	for _, v := range input {
		if v.pX < ((maxX-1)/2) && v.pY < ((maxY-1)/2) {
			q1++
		} else if v.pX < ((maxX-1)/2) && v.pY > ((maxY-1)/2) {
			q2++
		} else if v.pX > ((maxX-1)/2) && v.pY < ((maxY-1)/2) {
			q3++
		} else if v.pX > ((maxX-1)/2) && v.pY > ((maxY-1)/2) {
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

func reposition(input []Robot, time int) []Robot {

	if time == 0 {
		return input
	}

	newInput := make([]Robot, 0)

	for _, v := range input {
		newpx := v.pX + v.vX
		newpy := v.pY + v.vY

		if v.vX >= 0 && newpx >= maxX {
			newpx = newpx - maxX
		} else if v.vX < 0 && newpx < 0 {
			newpx = maxX + newpx
		}

		if v.vY >= 0 && newpy >= maxY {
			newpy = newpy - maxY
		} else if v.vY < 0 && newpy < 0 {
			newpy = maxY + newpy
		}

		newRobot := Robot{pX: newpx, pY: newpy, vX: v.vX, vY: v.vY}
		newInput = append(newInput, newRobot)
	}

	time--

	grid := sortRobots(newInput)
	day := findChristmasTree(time, grid)
	//fmt.Printf("Day %d\n", (test - day))
	if day > 0 {
		fmt.Printf("Day %d:\n", test-day)
		printRobots(grid)
	}
	//printRobots(grid)
	//fmt.Printf("NewRobot Sorted: %v\n", grid)

	return reposition(newInput, time)
}

func findChristmasTree(day int, grid map[int][]int) int {
	christmasTree := make(map[int]int)
	for i := 0; i < maxY; i++ {
		cols := grid[i]
		for j := 0; j < maxX; j++ {
			if contains5Consecutives(cols, j) {
				christmasTree[day] = christmasTree[day] + 1
				//fmt.Printf("Day %d on Row %d\n", day, i)
				continue
			}
		}
	}

	for v, k := range christmasTree {
		//fmt.Printf("Day %d: %d\n", v, k)
		if k > 5 {
			return v
		}
	}
	return -1
}

func printRobots(grid map[int][]int) {
	for i := 0; i < maxY; i++ {
		cols := grid[i]
		for j := 0; j < maxX; j++ {
			if contains(cols, j) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func contains5Consecutives(cols []int, value int) bool {
	if contains(cols, value) && contains(cols, value+1) && contains(cols, value+2) && contains(cols, value+3) && contains(cols, value+4) {
		return true
	}
	return false
}

func contains(cols []int, value int) bool {
	for _, v := range cols {
		if v == value {
			return true
		}
	}
	return false
}

func sortRobots(input []Robot) map[int][]int {
	grid := make(map[int][]int)
	for _, v := range input {
		grid[v.pY] = append(grid[v.pY], v.pX)
	}
	return grid
}

func SomeFunction2(input []string) (answer int) {
	return
}
