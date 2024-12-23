package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	answer1 := SomeFunction1()
	answer2 := SomeFunction2()

	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)

	t2 := time.Now()

	fmt.Printf("Time Taken: %v\n", t2.Sub(t1))
}

func SomeFunction1() (answer int) {
	return
}

func SomeFunction2() (answer int) {
	return
}
