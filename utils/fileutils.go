package utils

import (
	"bufio"
	"os"
)

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
