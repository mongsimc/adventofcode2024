package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Define a struct for 3D coordinates
type Coord struct {
	X, Y int
}

func main() {

	fmt.Println("Reading file")
	filename := "puzzle.input"
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var x, y int
	m := make(map[Coord]string)
	for scanner.Scan() {
		line := scanner.Text()
		for y = 0; y < len(line); y++ {
			// Print each character (as a token)
			c := Coord{X: x, Y: y}
			m[c] = string(line[y])
		}
		x++
	}

	var lineList []string
	lineList = append(lineList, ConstructLine(m, x, y, 1)...)
	lineList = append(lineList, ConstructLine(m, x, y, 2)...)
	lineList = append(lineList, ConstructLine(m, x, y, 3)...)
	lineList = append(lineList, ConstructLine(m, x, y, 4)...)

	answer := FindXMAS(m, x, y)

	re := regexp.MustCompile(`XMAS`)
	reRev := regexp.MustCompile(`SAMX`)

	var found []string
	for i := 0; i < len(lineList); i++ {
		split := re.FindAllString(lineList[i], -1)
		splitRev := reRev.FindAllString(lineList[i], -1)
		found = append(found, split...)
		found = append(found, splitRev...)
		//fmt.Printf("Result in line %d %s: %v\n", i, lineList[i], split)
	}

	// fmt.Printf("Lines with total %d vertical and %d horizontal: %v\n", x, y, lineList)
	fmt.Printf("Answer1: %d\n", len(found))
	fmt.Printf("Answer2: %d\n", answer)
}

func FindXMAS(input map[Coord]string, xTotal, yTotal int) int {
	var answer int
	for i := 0; i < xTotal; i++ {
		for j := 0; j < yTotal; j++ {
			k := Coord{X: i, Y: j}
			if input[k] == "A" && i > 0 && i+1 < xTotal && j > 0 && j+1 < yTotal {
				//fmt.Printf("Found A at coord: %v\n", k)
				a := input[Coord{X: i - 1, Y: j - 1}]
				b := input[Coord{X: i + 1, Y: j + 1}]
				if (a == "M" && b == "S") || (a == "S" && b == "M") {
					c := input[Coord{X: i - 1, Y: j + 1}]
					d := input[Coord{X: i + 1, Y: j - 1}]
					if (c == "M" && d == "S") || (c == "S" && d == "M") {
						answer++
					}
				}
			}
		}
	}
	return answer
}

func ConstructLine(input map[Coord]string, xTotal, yTotal, c int) []string {
	var list []string
	copiedmap := make(map[Coord]string)
	for key, value := range input {
		copiedmap[key] = value
	}

	switch c {
	case 1: //construct horizontally
		for i := 0; i < xTotal; i++ {
			var line string
			for j := 0; j < yTotal; j++ {
				k := Coord{X: i, Y: j}
				line = line + input[k]
			}
			// fmt.Printf("Lines %s\n", line)
			list = append(list, line)
		}
	case 2: //construct vertically
		for i := 0; i < yTotal; i++ {
			var line string
			for j := 0; j < xTotal; j++ {
				k := Coord{X: j, Y: i}
				line = line + input[k]
			}
			// fmt.Printf("Lines %s\n", line)
			list = append(list, line)
		}
	case 3: //construct diagonally from left
		var line string
		for i := 0; i < xTotal; i++ {
			for j := 0; j < yTotal; j++ {
				xStr := i
				yStr := j
				line = ""

				for {
					k := Coord{X: xStr, Y: yStr}
					value, ok := copiedmap[k]
					if xStr > xTotal || yStr < 0 || !ok {
						//fmt.Printf("2. Break: xStr is %d and YStr is %d\n", xStr, yStr)
						break
					}

					line = line + value
					delete(copiedmap, k)
					//fmt.Printf("xStr is %d and YStr is %d so line is %s\n", xStr, yStr, line)
					xStr++
					yStr--
				}
				if len(line) > 0 {
					list = append(list, line)
					//fmt.Printf("Lines %s\n", line)
				}
			}
		}
	case 4: //construct diagonally from right
		var line string
		for i := 0; i < xTotal; i++ {
			for j := yTotal - 1; j >= 0; j-- {
				xStr := i
				yStr := j
				line = ""

				for {
					k := Coord{X: xStr, Y: yStr}
					value, ok := copiedmap[k]
					if xStr > xTotal || yStr < 0 || !ok {
						//fmt.Printf("2. Break: xStr is %d and YStr is %d and exist %v\n", xStr, yStr, ok)
						break
					}

					line = line + value
					delete(copiedmap, k)
					//fmt.Printf("xStr is %d and YStr is %d so line is %s\n", xStr, yStr, line)
					xStr++
					yStr++
				}
				if len(line) > 0 {
					list = append(list, line)
					//fmt.Printf("Lines %s\n", line)
				}
			}
		}
	}

	return list
}
