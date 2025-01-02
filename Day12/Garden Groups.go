package main

import (
	"fmt"
	"time"

	aoc "github.com/mongsimc/adventofcode2024/utils"
)

var input, coords, maxRow, maxCol = aoc.ReadFileAsCoordinateWithString("puzzle.input")
var garden map[aoc.Coordinate]int
var areaMap map[int]int
var perimeterMap map[int]int
var topPerimeterMap map[aoc.Coordinate]int
var bottomPerimeterMap map[aoc.Coordinate]int
var leftPerimeterMap map[aoc.Coordinate]int
var rightPerimeterMap map[aoc.Coordinate]int
var discountedPerimeterMap map[int]int
var seq = 1

func main() {
	t1 := time.Now()

	garden = make(map[aoc.Coordinate]int)
	areaMap = make(map[int]int)
	perimeterMap = make(map[int]int)
	topPerimeterMap = make(map[aoc.Coordinate]int)
	bottomPerimeterMap = make(map[aoc.Coordinate]int)
	leftPerimeterMap = make(map[aoc.Coordinate]int)
	rightPerimeterMap = make(map[aoc.Coordinate]int)
	discountedPerimeterMap = make(map[int]int)

	for _, v := range coords {
		plot := garden[v]

		if plot == 0 {
			plot = seq
			seq++
		}

		groupPlant(v, plot)

		topPerimeterMap[getTopPerimeter(v, false)] = plot
		bottomPerimeterMap[getBottomPerimeter(v, false)] = plot
		leftPerimeterMap[getLeftPerimeter(v, false)] = plot
		rightPerimeterMap[getRightPerimeter(v, false)] = plot

	}

	// fmt.Printf("Garden: %v\n", garden)
	// fmt.Printf("AreaMap: %v\n", areaMap)
	// fmt.Printf("PerimeterMap: %v\n", perimeterMap)
	// fmt.Printf("TopPerimeterMap: %v\n", topPerimeterMap)
	// fmt.Printf("BottomPerimeterMap: %v\n", bottomPerimeterMap)
	// fmt.Printf("LeftPerimeterMap: %v\n", leftPerimeterMap)
	// fmt.Printf("RightPerimeterMap: %v\n", rightPerimeterMap)

	for _, v := range topPerimeterMap {
		discountedPerimeterMap[v] = discountedPerimeterMap[v] + 1
	}

	for _, v := range bottomPerimeterMap {
		discountedPerimeterMap[v] = discountedPerimeterMap[v] + 1
	}

	for _, v := range leftPerimeterMap {
		discountedPerimeterMap[v] = discountedPerimeterMap[v] + 1
	}

	for _, v := range rightPerimeterMap {
		discountedPerimeterMap[v] = discountedPerimeterMap[v] + 1
	}

	//fmt.Printf("DiscountedPerimeterMap: %v\n", discountedPerimeterMap)

	answer1 := calculatePrice()
	answer2 := calculateDiscountedPrice()

	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
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

	p := 4

	left := aoc.Coordinate{X: current.X, Y: current.Y - 1}
	if input[left] == input[current] {
		p--
		groupPlant(left, plot)
	}
	right := aoc.Coordinate{X: current.X, Y: current.Y + 1}
	if input[right] == input[current] {
		p--
		groupPlant(right, plot)
	}
	top := aoc.Coordinate{X: current.X - 1, Y: current.Y}
	if input[top] == input[current] {
		p--
		groupPlant(top, plot)
	}
	bottom := aoc.Coordinate{X: current.X + 1, Y: current.Y}
	if input[bottom] == input[current] {
		p--
		groupPlant(bottom, plot)
	}

	perimeterMap[plot] = perimeterMap[plot] + p
}

func getTopPerimeter(coord aoc.Coordinate, topMost bool) (topPerimeter aoc.Coordinate) {

	if coord.X == 0 || topMost {
		leftCoord := aoc.Coordinate{X: coord.X, Y: coord.Y - 1}
		diagCoord := aoc.Coordinate{X: coord.X - 1, Y: coord.Y - 1}

		if garden[diagCoord] == garden[coord] {
			return coord
		}

		if garden[leftCoord] != garden[coord] {
			return coord
		}

		return getTopPerimeter(leftCoord, false)
	}

	topCoord := aoc.Coordinate{X: coord.X - 1, Y: coord.Y}
	if garden[topCoord] != garden[coord] {
		return getTopPerimeter(coord, true)
	}

	return getTopPerimeter(topCoord, false)
}

func getBottomPerimeter(coord aoc.Coordinate, bottomMost bool) (topPerimeter aoc.Coordinate) {

	if coord.X == maxRow || bottomMost {
		leftCoord := aoc.Coordinate{X: coord.X, Y: coord.Y - 1}
		diagCoord := aoc.Coordinate{X: coord.X + 1, Y: coord.Y - 1}

		if garden[diagCoord] == garden[coord] {
			return coord
		}

		if garden[leftCoord] != garden[coord] {
			return coord
		}

		return getBottomPerimeter(leftCoord, false)
	}

	bottomCoord := aoc.Coordinate{X: coord.X + 1, Y: coord.Y}
	if garden[bottomCoord] != garden[coord] {
		return getBottomPerimeter(coord, true)
	}

	return getBottomPerimeter(bottomCoord, false)
}

func getLeftPerimeter(coord aoc.Coordinate, leftMost bool) (topPerimeter aoc.Coordinate) {

	if coord.Y == 0 || leftMost {
		topCoord := aoc.Coordinate{X: coord.X - 1, Y: coord.Y}
		diagCoord := aoc.Coordinate{X: coord.X - 1, Y: coord.Y - 1}

		if garden[diagCoord] == garden[coord] {
			return coord
		}

		if garden[topCoord] != garden[coord] {
			return coord
		}

		return getLeftPerimeter(topCoord, false)
	}

	leftCoord := aoc.Coordinate{X: coord.X, Y: coord.Y - 1}
	if garden[leftCoord] != garden[coord] {
		return getLeftPerimeter(coord, true)
	}

	return getLeftPerimeter(leftCoord, false)
}

func getRightPerimeter(coord aoc.Coordinate, rightMost bool) (topPerimeter aoc.Coordinate) {

	if coord.Y == maxCol || rightMost {
		topCoord := aoc.Coordinate{X: coord.X - 1, Y: coord.Y}
		diagCoord := aoc.Coordinate{X: coord.X - 1, Y: coord.Y + 1}

		if garden[diagCoord] == garden[coord] {
			return coord
		}

		if garden[topCoord] != garden[coord] {
			return coord
		}

		return getRightPerimeter(topCoord, false)
	}

	rightCoord := aoc.Coordinate{X: coord.X, Y: coord.Y + 1}
	if garden[rightCoord] != garden[coord] {
		return getRightPerimeter(coord, true)
	}

	return getRightPerimeter(rightCoord, false)
}
