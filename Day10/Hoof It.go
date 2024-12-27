package main

import (
	"fmt"
	"time"

	aoc "github.com/mongsimc/adventofcode2024/utils"
)

var input = aoc.ReadFileAsCoordinate("puzzle.input")

type Node struct {
	data aoc.Coordinate
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Append(data aoc.Coordinate) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func (ll *LinkedList) Delete(value aoc.Coordinate) {
	if ll.head == nil {
		return
	}

	if ll.head.data == value {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (ll *LinkedList) Search(value aoc.Coordinate) bool {
	current := ll.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList) Length() int {
	size := 0
	current := ll.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (ll *LinkedList) Merge(newLL *LinkedList) {
	current := newLL.head
	for current != nil {
		ll.Append(current.data)
		current = current.next
	}
}

var ll [][]aoc.Coordinate
var distinctList [][]aoc.Coordinate

func main() {
	t1 := time.Now()
	ll = make([][]aoc.Coordinate, 0)
	distinctList = make([][]aoc.Coordinate, 0)

	answer1 := findTrailheadScore()
	answer2 := SomeFunction2()

	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func findTrailheadScore() (answer int) {

	//nodeMap := make(map[aoc.Coordinate][]aoc.Coordinate, 0)
	//paths := make([]aoc.Coordinate, 0)
	//trailheads := make([]aoc.Coordinate, 0)

	for k, v := range input {

		//nodeMap[k] = findAdjacent(k)
		//node := Node{data: k}

		// for x, y := range createNodes(node) {
		// 	nodeMap[x] = y
		// }
		if v == 0 {
			//trailheads = append(trailheads, k)
			path := make([]aoc.Coordinate, 0)
			path = append(path, k)
			createPath(k, path)
		}
	}

	for _, v := range ll {
		//fmt.Printf("Posible Paths: %v\n", v)
		if len(v) == 10 {

			checkDistinctList(v)
		}
	}

	// for _, v := range distinctList {
	// 	fmt.Printf("All Paths: %v\n", v)
	// }

	answer = len(distinctList)
	//answer = checkPathScore(distinctList)
	return
}

func checkDistinctList(list []aoc.Coordinate) {

	// if len(distinctList) == 0 {
	// 	distinctList = append(distinctList, list)
	// 	return
	// }

	for _, v := range distinctList {
		if v[0] == list[0] && v[len(v)-1] == list[len(v)-1] {
			return
		}
	}
	distinctList = append(distinctList, list)
}

func createPath(currentCoord aoc.Coordinate, path []aoc.Coordinate) {

	left, t1 := getAdjacent(currentCoord, 1)
	right, t2 := getAdjacent(currentCoord, 2)
	up, t3 := getAdjacent(currentCoord, 3)
	down, t4 := getAdjacent(currentCoord, 4)

	if t1 {
		//fmt.Printf("Current: %v and Left: %v\n", currentCoord, left)
		list1 := make([]aoc.Coordinate, len(path))
		copy(list1, path)
		//nextnode := Node{data: left}
		//currentnode := Node{data: currentCoord, next: &nextnode}
		//list1 := LinkedList{head: &currentnode}
		list1 = append(list1, left)
		createPath(left, list1)
		ll = append(ll, list1)
	}

	if t2 {
		//fmt.Printf("Current: %v and Right: %v\n", currentCoord, right)
		list2 := make([]aoc.Coordinate, len(path))
		copy(list2, path)
		//nextnode := Node{data: right}
		//currentnode := Node{data: currentCoord, next: &nextnode}
		//list1 := LinkedList{head: &currentnode}
		list2 = append(list2, right)
		createPath(right, list2)
		ll = append(ll, list2)
	}

	if t3 {
		//fmt.Printf("Current: %v and Up: %v\n", currentCoord, up)
		list3 := make([]aoc.Coordinate, len(path))
		copy(list3, path)
		//nextnode := Node{data: up}
		//currentnode := Node{data: currentCoord, next: &nextnode}
		//list1 := LinkedList{head: &currentnode}
		list3 = append(list3, up)
		createPath(up, list3)
		ll = append(ll, list3)
	}

	if t4 {
		//fmt.Printf("Current: %v and Down: %v\n", currentCoord, down)
		list4 := make([]aoc.Coordinate, len(path))
		copy(list4, path)
		//nextnode := Node{data: down}
		//currentnode := Node{data: currentCoord, next: &nextnode}
		//list1 := LinkedList{head: &currentnode}
		list4 = append(list4, down)
		createPath(down, list4)
		ll = append(ll, list4)
	}
}

func getAdjacent(current aoc.Coordinate, direction int) (aoc.Coordinate, bool) {

	var next aoc.Coordinate

	switch direction {
	case 1:
		next = aoc.Coordinate{X: current.X, Y: current.Y - 1}
	case 2:
		next = aoc.Coordinate{X: current.X, Y: current.Y + 1}
	case 3:
		next = aoc.Coordinate{X: current.X - 1, Y: current.Y}
	case 4:
		next = aoc.Coordinate{X: current.X + 1, Y: current.Y}
	}

	if input[current] == 9 {
		return current, false
	}

	if input[next] > 0 && input[next]-input[current] == 1 {
		return next, true
	}

	return current, false

}

func hasAdjacent(current aoc.Coordinate) bool {
	return false
}

func createPaths(ll LinkedList, current aoc.Coordinate, nodeMap map[aoc.Coordinate][]aoc.Coordinate, first bool) []LinkedList {
	paths := make([]LinkedList, 0)
	var new LinkedList
	for _, v := range nodeMap[current] {
		fmt.Printf("creating path for: %v with %v", current, v)
		if first {
			new = LinkedList{head: ll.head}
		} else {
			new = ll
		}
		node := Node{data: v}
		newLL := LinkedList{head: &node}
		new.Merge(&newLL)
		paths = append(paths, new)
		createPaths(new, v, nodeMap, false)
	}

	return paths
}

func createNodes(node Node) (nodes map[Node]aoc.Coordinate) {

	nodes = make(map[Node]aoc.Coordinate)
	newCoords := findAdjacent(node.data)

	for _, v := range newCoords {
		newNode := Node{data: v}
		newCurrent := Node{data: node.data, next: &newNode}
		nodes[newCurrent] = v
		//fmt.Printf("newCurrent: %v and next: %v\n", &newCurrent, &newNode)
	}
	return nodes
}

func findAdjacent(current aoc.Coordinate) []aoc.Coordinate {

	left := aoc.Coordinate{X: current.X, Y: current.Y - 1}
	right := aoc.Coordinate{X: current.X, Y: current.Y + 1}
	up := aoc.Coordinate{X: current.X - 1, Y: current.Y}
	down := aoc.Coordinate{X: current.X + 1, Y: current.Y}

	nextCoords := make([]aoc.Coordinate, 0)

	if input[current] == 9 {
		return nextCoords
	}
	if input[left] > 0 && input[left]-input[current] == 1 {
		nextCoords = append(nextCoords, left)
	}
	if input[right] > 0 && input[right]-input[current] == 1 {
		nextCoords = append(nextCoords, right)
	}
	if input[up] > 0 && input[up]-input[current] == 1 {
		nextCoords = append(nextCoords, up)
	}
	if input[down] > 0 && input[down]-input[current] == 1 {
		nextCoords = append(nextCoords, down)
	}
	return nextCoords
}

func checkNext2(current aoc.Coordinate, direction int) (path []aoc.Coordinate) {
	path = make([]aoc.Coordinate, 0)

	//TODO: need to check if exist in current path, if not it will traverse back

	switch direction {
	case 1: //left
		next := aoc.Coordinate{X: current.X, Y: current.Y - 1}
		fmt.Printf("NextCoordinate: %v\n", next)
		hasNext, nextOfNext := hasNextCoordinate(current)
		fmt.Printf("HasNextCoordinate Result: %t and %v\n", hasNext, nextOfNext)
		for hasNext {

			for _, v := range nextOfNext {
				hasNext, nextOfNext = hasNextCoordinate(v)
				path = append(path, nextOfNext...)
			}
		}

	case 2: //right
		next := aoc.Coordinate{X: current.X, Y: current.Y + 1}
		hasNext, nextOfNext := hasNextCoordinate(next)
		for hasNext {

			for _, v := range nextOfNext {
				hasNext, nextOfNext = hasNextCoordinate(v)
				path = append(path, nextOfNext...)
			}
		}
	case 3: //up
		next := aoc.Coordinate{X: current.X - 1, Y: current.Y}
		hasNext, nextOfNext := hasNextCoordinate(next)
		for hasNext {

			for _, v := range nextOfNext {
				hasNext, nextOfNext = hasNextCoordinate(v)
				path = append(path, nextOfNext...)
			}
		}
	case 4: //down
		next := aoc.Coordinate{X: current.X + 1, Y: current.Y}
		hasNext, nextOfNext := hasNextCoordinate(next)
		for hasNext {

			for _, v := range nextOfNext {
				hasNext, nextOfNext = hasNextCoordinate(v)
				path = append(path, nextOfNext...)
			}
		}
	}
	return path
}

func hasNextCoordinate(current aoc.Coordinate) (bool, []aoc.Coordinate) {
	//TODO: need to check if exist in current path

	nextCoordinates := make([]aoc.Coordinate, 0)

	left := aoc.Coordinate{X: current.X, Y: current.Y - 1}
	right := aoc.Coordinate{X: current.X, Y: current.Y + 1}
	up := aoc.Coordinate{X: current.X - 1, Y: current.Y}
	down := aoc.Coordinate{X: current.X + 1, Y: current.Y}

	if input[left] >= 0 {
		nextCoordinates = append(nextCoordinates, left)
	}
	if input[right] >= 0 {
		nextCoordinates = append(nextCoordinates, right)
	}
	if input[up] >= 0 {
		nextCoordinates = append(nextCoordinates, up)
	}
	if input[down] >= 0 {
		nextCoordinates = append(nextCoordinates, down)
	}

	if len(nextCoordinates) > 0 {
		return true, nextCoordinates
	}

	return false, nextCoordinates
}

func checkPathScore(paths []LinkedList) int {
	return 0
}

func SomeFunction2() (answer int) {
	return
}
