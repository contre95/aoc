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
	n := len(forest)
	for i < n-1 {
		j := 1
		for j < n-1 {
			isHidden := []bool{false, false, false, false}
			fmt.Printf("%d ", forest[i][j])
			for x := 0; x < i; x++ {
				if forest[x][j] >= forest[i][j] {
					isHidden[0] = true
				}
			}
			for x := i + 1; x < n; x++ {
				if forest[x][j] >= forest[i][j] {
					isHidden[1] = true
				}
			}
			for y := 0; y < j; y++ {
				if forest[i][y] >= forest[i][j] {
					isHidden[2] = true
				}
			}
			for y := j + 1; y < n; y++ {
				if forest[i][y] >= forest[i][j] {
					isHidden[3] = true
				}
			}
			if !isHidden[0] || !isHidden[1] || !isHidden[2] || !isHidden[3] {
				count++
			}
			j++
		}
		fmt.Println()
		i++
	}
	fmt.Println(count)
}
