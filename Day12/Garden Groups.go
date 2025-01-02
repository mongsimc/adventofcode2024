package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	aoc "github.com/mongsimc/adventofcode2024/utils"
)

var input, coords = aoc.ReadFileAsCoordinateWithString("sample2.input")
var garden map[aoc.Coordinate]int
var areaMap map[int]int
var perimeterMap map[int]int
var discountedPerimeterMap map[int]int
var topPerimeterMap map[string]int
var bottomPerimeterMap map[string]int
var leftPerimeterMap map[string]int
var rightPerimeterMap map[string]int
var seq = 1

func main() {
	t1 := time.Now()

	garden = make(map[aoc.Coordinate]int)
	areaMap = make(map[int]int)
	perimeterMap = make(map[int]int)
	discountedPerimeterMap = make(map[int]int)
	topPerimeterMap = make(map[string]int)
	bottomPerimeterMap = make(map[string]int)
	leftPerimeterMap = make(map[string]int)
	rightPerimeterMap = make(map[string]int)

	for _, v := range coords {
		plot := garden[v]

		if plot == 0 {
			plot = seq
			seq++
		}

		groupPlant(v, plot)
	}

	calculateDiscountedPerimeter()

	// fmt.Printf("Garden: %v\n", garden)
	// fmt.Printf("AreaMap: %v\n", areaMap)
	fmt.Printf("PerimeterMap: %v\n", perimeterMap)
	fmt.Printf("DiscountedPerimeterMap: %v\n", discountedPerimeterMap)

	answer1 := calculatePrice()
	answer2 := calculateDiscountedPrice()

	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func calculateDiscountedPerimeter() {
	for k, v := range topPerimeterMap {
		keys := strings.Split(k, "_")
		plot, _ := strconv.Atoi(keys[0])
		discountedPerimeterMap[plot] = discountedPerimeterMap[plot] + v
	}
	for k, v := range bottomPerimeterMap {
		keys := strings.Split(k, "_")
		plot, _ := strconv.Atoi(keys[0])
		discountedPerimeterMap[plot] = discountedPerimeterMap[plot] + v
	}
	for k, v := range leftPerimeterMap {
		keys := strings.Split(k, "_")
		plot, _ := strconv.Atoi(keys[0])
		discountedPerimeterMap[plot] = discountedPerimeterMap[plot] + v
	}
	for k, v := range rightPerimeterMap {
		keys := strings.Split(k, "_")
		plot, _ := strconv.Atoi(keys[0])
		discountedPerimeterMap[plot] = discountedPerimeterMap[plot] + v
	}
}

func calculatePrice() (price int) {
	for i := 1; i < seq; i++ {
		price = price + (areaMap[i] * perimeterMap[i])
		// fmt.Printf("Price for plot %d: %d\n", i, price)
	}
	return
}

func calculateDiscountedPrice() (price int) {
	for i := 1; i < seq; i++ {
		price = price + (areaMap[i] * discountedPerimeterMap[i])
		// fmt.Printf("Price for plot %d: %d\n", i, price)
	}
	return
}

func groupPlant(current aoc.Coordinate, plot int) {

	if garden[current] != 0 {
		return
	}

	garden[current] = plot

	area := areaMap[plot]
	area = area + 1
	areaMap[plot] = area

	hasLeft := false
	hasTop := false
	hasRight := false
	hasBottom := false

	p := 4

	left := aoc.Coordinate{X: current.X, Y: current.Y - 1}
	if input[left] == input[current] {
		p--
		hasLeft = true
		groupPlant(left, plot)
	}
	right := aoc.Coordinate{X: current.X, Y: current.Y + 1}
	if input[right] == input[current] {
		p--
		hasRight = true
		groupPlant(right, plot)
	}
	top := aoc.Coordinate{X: current.X - 1, Y: current.Y}
	if input[top] == input[current] {
		p--
		hasTop = true
		groupPlant(top, plot)
	}
	bottom := aoc.Coordinate{X: current.X + 1, Y: current.Y}
	if input[bottom] == input[current] {
		p--
		hasBottom = true
		groupPlant(bottom, plot)
	}

	perimeterMap[plot] = perimeterMap[plot] + p

	calculateContinuousPerimeter(current, plot, hasLeft, hasRight, hasTop, hasBottom)

}

func calculateContinuousPerimeter(current aoc.Coordinate, plot int, left, right, top, bottom bool) (perimeter int) {
	horizontalKey := strconv.Itoa(plot) + "_" + strconv.Itoa(current.Y)
	verticalKey := strconv.Itoa(plot) + "_" + strconv.Itoa(current.X)

	topPerimeter := topPerimeterMap[horizontalKey]
	bottomPerimeter := bottomPerimeterMap[horizontalKey]
	leftPerimeter := leftPerimeterMap[verticalKey]
	rightPerimeter := rightPerimeterMap[verticalKey]

	if !top && !left {
		topPerimeterMap[horizontalKey] = topPerimeter + 1
	}

	if !left && !top {
		leftPerimeterMap[verticalKey] = leftPerimeter + 1
	}

	if !bottom && !left {
		bottomPerimeterMap[horizontalKey] = bottomPerimeter + 1
	}

	if !right && !top {
		rightPerimeterMap[verticalKey] = rightPerimeter + 1
	}

	fmt.Printf("Coord: %v and TopPerimeter: %v\n", current, topPerimeterMap)
	fmt.Printf("Coord: %v and BottomPerimeter: %v\n", current, bottomPerimeterMap)
	fmt.Printf("Coord: %v and LeftPerimeter: %v\n", current, leftPerimeterMap)
	fmt.Printf("Coord: %v and RightPerimeter: %v\n", current, rightPerimeterMap)
	return
}
