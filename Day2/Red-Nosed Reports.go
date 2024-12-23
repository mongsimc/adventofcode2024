package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reading file")
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Fields(line)
		result, newList1, newList2, newList3 := CheckReport(s)
		if result {
			count++
		} else {
			result1, _, _, _ := CheckReport(newList1)
			result2, _, _, _ := CheckReport(newList2)
			result3, _, _, _ := CheckReport(newList3)
			if result1 || result2 || result3 {
				count++
			} else {
				fmt.Printf("Report: %v\n", s)
			}
		}
	}
	fmt.Printf("Answer: %d\n", count)
	file.Close()
}

func CheckReport(reportList []string) (bool, []string, []string, []string) {
	// fmt.Printf("Report: %v\n", reportList)
	var isAsc, isDesc bool
	var newList1, newList2, newList3 []string
	for i := 0; i < len(reportList)-1; i += 1 {
		copiedList := make([]string, len(reportList))
		copiedList2 := make([]string, len(reportList))
		copiedList3 := make([]string, len(reportList))
		copy(copiedList, reportList)
		copy(copiedList2, reportList)
		copy(copiedList3, reportList)
		curr, err := strconv.Atoi(reportList[i])
		next, err := strconv.Atoi(reportList[i+1])
		if err != nil {
			panic(err)
		}
		diff := curr - next
		diffAbs := int(math.Abs(float64(diff)))
		// fmt.Printf("Diff: %d\n", diff)
		if diffAbs < 1 || diffAbs > 3 {
			newList1 = DeleteElement(copiedList, i-1)
			newList2 = DeleteElement(copiedList2, i)
			newList3 = DeleteElement(copiedList3, i+1)
			// fmt.Printf("Violate rule 1: %v, %v, %v\n", newList1, newList2, newList3)
			return false, newList1, newList2, newList3
		}
		//diff > 0 is desc
		if !isAsc && !isDesc {
			if diff > 0 {
				// fmt.Println("Value is Desc")
				isDesc = true
			} else if diff < 0 {
				// fmt.Println("Value is Asc")
				isAsc = true
			}
		} else if (diff > 0 && isAsc) || (diff < 0 && isDesc) {
			newList1 = DeleteElement(copiedList, i-1)
			newList2 = DeleteElement(copiedList2, i)
			newList3 = DeleteElement(copiedList3, i+1)
			// fmt.Printf("Violate rule 2: %v, %v, %v\n", newList1, newList2, newList3)
			return false, newList1, newList2, newList3
		}
	}
	return true, newList1, newList2, newList3
}

func DeleteElement(slice []string, index int) []string {

	if index < 0 || index >= len(slice) {
		return slice // Invalid index, return the original slice
	}
	// Remove the element at the specified index by slicing
	if (index + 1) == len(slice) {
		return slice[0:index]
	}

	return append(slice[:index], slice[index+1:]...)
}
