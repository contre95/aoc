package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inputParser(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var input []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		input = append(input, line)
	}
	return input
}

func HasDuplicates(s string) bool {
	for _, c := range s {
		if strings.Count(s, string(c)) > 1 {
			return true
		}
	}
	return false
}

func part(x int) int {
	input := inputParser(os.Args[1])
	for _, stream := range input {
		count := 0
		for i := 0; i < len(stream)-3; i++ {
			if HasDuplicates(stream[i : i+x]) {
				count++
			} else {
				return count + x
			}
		}
	}
	return 0
}

func main() {
	fmt.Printf("Part 1: %d \n", part(4))
	fmt.Printf("Part 2: %d \n", part(14))
}
