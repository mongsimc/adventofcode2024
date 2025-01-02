package utils

import (
	"bufio"
	"os"
	"strconv"
)

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) ToString() string {
	return "(" + strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y) + ")"
}

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadFileAsCoordinate(filename string) map[Coordinate]int {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	x := 0
	coords := make(map[Coordinate]int)

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			value, _ := strconv.Atoi(line[i : i+1])
			coord := Coordinate{X: x, Y: i}
			coords[coord] = value
		}
		x++
	}

	return coords
}

func ReadFileAsCoordinateWithString(filename string) (map[Coordinate]string, []Coordinate, int, int) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	x := 0
	coordMap := make(map[Coordinate]string)
	coords := make([]Coordinate, 0)
	maxRow := 0
	maxCol := 0

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			value := line[i : i+1]
			coord := Coordinate{X: x, Y: i}
			coordMap[coord] = value
			coords = append(coords, coord)
			maxRow = x
			maxCol = i
		}
		x++
	}

	return coordMap, coords, maxRow, maxCol
}
