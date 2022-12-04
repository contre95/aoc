package main

import (
	"bufio"
	"fmt"
	"os"
)

func fileToA(filePath string) []string {
	var a []string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		a = append(a, fileScanner.Text())
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	return a
}

func priority(c rune) int {
	if int(c) <= 90 {
		return int(c) - 38
	}
	return int(c) - 96
}

func part() {
	input := fileToA(os.Args[1])
	prioritySum := 0
	for _, rucksack := range input {
		pocketOneTypes := map[rune]bool{}
		pocketTwoTypes := map[rune]bool{}
		for _, item := range rucksack[:len(rucksack)/2] {
			pocketOneTypes[item] = true
		}
		for _, item := range rucksack[len(rucksack)/2:] {
			pocketTwoTypes[item] = true
		}
		for item := range pocketTwoTypes {
			if _, ok := pocketOneTypes[item]; ok {
				prioritySum += priority(item)
			}
		}
	}
	fmt.Println(prioritySum)
}

func main() {
	part()
}
