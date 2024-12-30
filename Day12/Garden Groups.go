package main

import (
	"fmt"
	"time"

	aoc "github.com/mongsimc/adventofcode2024/utils"
)

var input = aoc.ReadFileAsCoordinateWithString("sample.input")

func main() {
	t1 := time.Now()

	answer1 := SomeFunction1(input)
	answer2 := SomeFunction2(input)

	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func getAdjacent(current aoc.Coordinate, direction int) string {

	var next aoc.Coordinate

	switch direction {
	case 1: //left
		next = aoc.Coordinate{X: current.X, Y: current.Y - 1}
	case 2: //right
		next = aoc.Coordinate{X: current.X, Y: current.Y + 1}
	case 3: //up
		next = aoc.Coordinate{X: current.X - 1, Y: current.Y}
	case 4: //down
		next = aoc.Coordinate{X: current.X + 1, Y: current.Y}
	case 5: //top-left
		next = aoc.Coordinate{X: current.X - 1, Y: current.Y - 1}
	case 6: //top-right
		next = aoc.Coordinate{X: current.X - 1, Y: current.Y + 1}
	case 7: //bottom-left
		next = aoc.Coordinate{X: current.X + 1, Y: current.Y - 1}
	case 8: //bottom-right
		next = aoc.Coordinate{X: current.X + 1, Y: current.Y + 1}
	}

	return input[next]

}

func SomeFunction1(input map[aoc.Coordinate]string) (answer int) {
	return
}

func SomeFunction2(input map[aoc.Coordinate]string) (answer int) {
	return
}
