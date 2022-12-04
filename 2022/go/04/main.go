package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fileToA(filePath string) [][]int {
	var a [][]int
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), ",")
		p1a, _ := strconv.Atoi(strings.Split(line[0], "-")[0])
		p2a, _ := strconv.Atoi(strings.Split(line[1], "-")[0])
		p1b, _ := strconv.Atoi(strings.Split(line[0], "-")[1])
		p2b, _ := strconv.Atoi(strings.Split(line[1], "-")[1])
		a = append(a, []int{p1a, p1b, p2a, p2b})
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	return a
}

func main() {
	input := fileToA(os.Args[1])
	count := 0
	for _, v := range input {
		if (v[2] >= v[0] && v[3] <= v[1]) || (v[2] <= v[0] && v[3] >= v[1]) {
			count += 1
		}
	}
	fmt.Println(count)
}
