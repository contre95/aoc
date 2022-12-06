package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			if _, ok := stacks[count]; ok {
				stacks[count] = string(line[i]) + stacks[count]
			} else {
				stacks[count] = string(line[i])
			}
			count++
		}
	}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		r := strings.NewReplacer("move", "", "from", "", "to", "", " ", "")
		op := []int{}
		for _, v := range r.Replace(line) {
			i, _ := strconv.Atoi(string(v))
			op = append(op, i)
		}
		ops = append(ops, op)
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	return stacks, ops[1 : len(ops)-2]
}

func main() {
	stacks, ops := inputParser(os.Args[1])
	fmt.Println(stacks, ops)
}
