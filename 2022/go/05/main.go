package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputParser(filePath string) (map[int]string, [][]int) {
	stacks := map[int]string{}
	var ops [][]int
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if string(line[1]) == "1" {
			break
		}
		count := 1
		for i := 1; i < len(line); i += 4 {
			if _, ok := stacks[i-1]; ok {
				stacks[count] = string(line[i]) + stacks[i-1]
			} else {
				stacks[count] = string(line[i])
			}
			count++
		}
		fmt.Println(stacks)
		// We are still in the first part of the input
		//r := strings.NewReplacer("[", "", "]", "")
		//fmt.Println(r.Replace(line))
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	return stacks, ops
}

func main() {
	inputParser(os.Args[1])
}
