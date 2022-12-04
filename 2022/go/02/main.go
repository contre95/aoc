package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fileToA(filePath string) [][]string {
	var a [][]string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		a = append(a, strings.Split(fileScanner.Text(), " "))
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	return a
}

type Results map[string]int

func part(a, b, c Results) {
	points := map[string]Results{"A": a, "B": b, "C": c}
	score := 0
	for _, m := range fileToA(os.Args[1]) {
		score += points[m[0]][m[1]]
	}
	fmt.Println(score)

}
func main() {
	a1 := Results{"X": 4, "Y": 8, "Z": 3}
	b1 := Results{"X": 1, "Y": 5, "Z": 9}
	c1 := Results{"X": 7, "Y": 2, "Z": 6}

	a2 := Results{"X": 3, "Y": 4, "Z": 8}
	b2 := Results{"X": 1, "Y": 5, "Z": 9}
	c2 := Results{"X": 2, "Y": 6, "Z": 7}
	part(a1, b1, c1)
	part(a2, b2, c2)
}
