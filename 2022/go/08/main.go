package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func part() [][]int {
	input := inputParser(os.Args[1])
	forest := [][]int{}
	for _, row := range input {
		treeLine := []int{}
		for _, col := range row {
			height, _ := strconv.Atoi(string(col))
			treeLine = append(treeLine, height)
		}
		forest = append(forest, treeLine)
	}
	return forest
}

func main() {
	forest := part()
	count := len(forest)*4 - 4
	i := 1
	j := 1
	n := len(forest)
	for i > n && j > n {
		count++
		i++
		j++
	}
}
