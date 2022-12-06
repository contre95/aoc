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
			if string(line[i]) != "" {
				if _, ok := stacks[count]; ok {
					stacks[count] = strings.Replace(string(line[i]), " ", "", -1) + stacks[count]
				} else {
					stacks[count] = strings.Replace(string(line[i]), " ", "", -1)
				}
				count++
			}
		}
	}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		r := strings.NewReplacer("move", "", "from", "", "to", "")
		op := []int{}
		for _, v := range strings.Fields(r.Replace(line)) {
			i, _ := strconv.Atoi(string(v))
			op = append(op, i)
		}
		ops = append(ops, op)
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	return stacks, ops[1 : len(ops)-1]
}

func main() {
	stacks, ops := inputParser(os.Args[1])
	for _, op := range ops {
		move := op[0]
		from := op[1]
		to := op[2]
		//fmt.Printf("%v :\n Move %d from %d (%v) to %d (%v) \n", stacks, move, from, stacks[from], to, stacks[to])
		buff := stacks[from][len(stacks[from])-move:]
		for i := len(buff) - 1; i >= 0; i-- {
			stacks[to] += string(stacks[from][len(stacks[from])-move:][i])
		}
		stacks[from] = stacks[from][:len(stacks[from])-move]
	}
	for _, stack := range stacks {
		fmt.Printf("%c", stack[len(stack)-1])
	}

}
