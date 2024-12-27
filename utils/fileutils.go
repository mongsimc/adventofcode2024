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
