package main

import (
	"fmt"
	"strconv"
	"time"

	aoc "github.com/mongsimc/adventofcode2024/utils"
)

type customMap struct {
	index    int
	position int
	space    int
}

func main() {
	t1 := time.Now()

	lines := aoc.ReadFile("puzzle.input")

	amphipod, fileMap, spaceMap := processInput(lines[0])

	answer1 := processDiskMap(amphipod)
	answer2 := processDiskMap2(amphipod, fileMap, spaceMap)

	fmt.Printf("Answer1: %d\n", answer1) //6259790630969
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func processDiskMap(amphipod []int) (answer int) {
	//fmt.Printf("Before refragment: %v\n", amphipod)
	newAmphipod := refragment(amphipod)
	//fmt.Printf("After refragment: %v\n", newAmphipod)
	answer = checksum(newAmphipod)
	return
}

func processDiskMap2(amphipod []int, fileMap, spaceMap []customMap) (answer int) {
	//fmt.Printf("Before refragment: %v\n", amphipod)
	newAmphipod := refragment2(amphipod, fileMap, spaceMap)
	//fmt.Printf("After refragment: %v\n", newAmphipod)
	answer = checksum(newAmphipod)
	return
}

func checksum(amphipod []int) (sum int) {
	for i, v := range amphipod {
		if v > 0 {
			sum = sum + (i * v)
		}
	}
	return
}

func refragment(amphipod []int) (newAmphipod []int) {
	newAmphipod = make([]int, len(amphipod))
	copy(newAmphipod, amphipod)
	freeSpace, _ := nextFreeSpace(newAmphipod)
	for freeSpace > 0 {
		newAmphipod = moveToFreeSpace(newAmphipod, freeSpace)
		freeSpace, _ = nextFreeSpace(newAmphipod)
	}
	return
}

func refragment2(amphipod []int, fileMap, spaceMap []customMap) (newAmphipod []int) {
	newAmphipod = make([]int, len(amphipod))
	copy(newAmphipod, amphipod)
	for i := len(fileMap) - 1; i > 0; i-- {

		spaceNeeded := fileMap[i].space
		originalPos := fileMap[i].position
		for index, key := range spaceMap {
			spaceAvailable := key.space
			newPosition := key.position
			//fmt.Printf("Checking for FileID: %d on original Position: %d for new Position: %d\n", i, originalPos, newPosition)
			if spaceAvailable >= spaceNeeded && newPosition < originalPos {
				newAmphipod = moveToFreeSpace2(newAmphipod, spaceNeeded, i, originalPos, newPosition)
				newSpace := customMap{index: index, position: newPosition + spaceNeeded, space: spaceAvailable - spaceNeeded}
				spaceMap[index] = newSpace
				break
			}
		}

	}
	return
}

func moveToFreeSpace2(amphipod []int, spaceNeeded, value, originalPos, newPos int) (newAmphipod []int) {

	for i := 0; i < spaceNeeded; i++ {
		amphipod[newPos+i] = value
		amphipod[originalPos+i] = -1
	}

	newAmphipod = make([]int, len(amphipod))
	copy(newAmphipod, amphipod)
	return
}

func moveToFreeSpace(amphipod []int, freeSpace int) (newAmphipod []int) {
	amphipod[freeSpace] = amphipod[len(amphipod)-1]
	newAmphipod = make([]int, len(amphipod)-1)
	copy(newAmphipod, amphipod)
	return
}

func nextFreeSpace(amphipod []int) (int, int) {
	for i, v := range amphipod {
		if v == -1 {
			j := nextFileBlock(amphipod, i) - i
			return i, j
		}
	}
	return -1, -1
}

func nextFileBlock(amphipod []int, offset int) int {
	for i := offset; i < len(amphipod); i++ {
		v := amphipod[i]
		if v > 0 {
			return i
		}
	}
	return -1
}

func populate(amphipod []int, fileID int, fileBlock int, freeSpace int) (newAmphipod []int) {
	newAmphipod = make([]int, len(amphipod))
	copy(newAmphipod, amphipod)
	//fmt.Printf("Before populating: %v\n", newAmphipod)
	for i := 0; i < fileBlock; i++ {
		newAmphipod = append(newAmphipod, fileID)
	}
	for j := 0; j < freeSpace; j++ {
		newAmphipod = append(newAmphipod, -1)
	}
	//fmt.Printf("After populating: %v\n", newAmphipod)
	return newAmphipod
}

func processInput(input string) (amphipod []int, fileMap []customMap, spaceMap []customMap) {
	fileID := 0
	for i := 0; i < len(input)-1; i = i + 2 {
		fileBlock, _ := strconv.Atoi(input[i : i+1])
		freeSpace, _ := strconv.Atoi(input[i+1 : i+2])
		amphipod = populate(amphipod, fileID, fileBlock, freeSpace)
		file := customMap{index: fileID, position: len(amphipod) - freeSpace - fileBlock, space: fileBlock}
		space := customMap{index: fileID, position: len(amphipod) - freeSpace, space: freeSpace}
		fileMap = append(fileMap, file)
		spaceMap = append(spaceMap, space)

		fileID++
	}
	lastFileBlock, _ := strconv.Atoi(input[len(input)-1 : len(input)])
	amphipod = populate(amphipod, fileID, lastFileBlock, 0)
	file := customMap{index: fileID, position: len(amphipod) - lastFileBlock, space: lastFileBlock}
	fileMap = append(fileMap, file)
	return
}
