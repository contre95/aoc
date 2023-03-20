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

func getForest(input []string) [][]int {
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

func Product(numbers []int) int {
	prod := 1
	for _, number := range numbers {
		prod *= number
	}
	return prod
}

func main() {
	input := inputParser(os.Args[1])
	forest := getForest(input)
	count := len(forest)*4 - 4
	maxScenicScore := 0
	i := 1
	n := len(forest)
	for i < n-1 {
		j := 1
		for j < n-1 {
			isHidden := []bool{false, false, false, false}
			scenicScore := []int{0, 0, 0, 0}
			//fmt.Printf("%d ", forest[i][j])
			for x := 0; x < i; x++ {
				if forest[x][j] >= forest[i][j] {
					isHidden[0] = true
					scenicScore[0] = 1
				} else {
					scenicScore[0] += 1
				}
			}
			for x := i + 1; x < n; x++ {
				if forest[x][j] >= forest[i][j] {
					isHidden[1] = true
					scenicScore[1] += 1
					break
				} else {
					scenicScore[1] += 1
				}
			}
			for y := 0; y < j; y++ {
				if forest[i][y] >= forest[i][j] {
					isHidden[2] = true
					scenicScore[2] = 1
				} else {
					scenicScore[2] += 1
				}
			}
			for y := j + 1; y < n; y++ {
				if forest[i][y] >= forest[i][j] {
					isHidden[3] = true
					scenicScore[3] += 1
					break
				} else {
					scenicScore[3] += 1
				}
			}

			if !isHidden[0] || !isHidden[1] || !isHidden[2] || !isHidden[3] {
				count++
			}
			if Product(scenicScore) >= maxScenicScore {
				maxScenicScore = Product(scenicScore)
			}
			j++
		}
		//fmt.Println()
		i++
	}
	fmt.Println(count)
	fmt.Println(maxScenicScore)
}
