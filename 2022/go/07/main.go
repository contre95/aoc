package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

type Directory struct {
	size int
	dir  map[string]Directory
}

func main() {
	var pwd []string // Stack
	sizes := map[string]int{}
	dirs := map[string]Directory{}
	count := 0
	input := inputParser(os.Args[1])
	for _, cmd := range input {
		fmt.Println(sizes)
		switch {
		case string(cmd[0]) == "$":
			if cmd[2:4] == "cd" {
				if cmd[4:] != ".." {
					pwd = append(pwd, cmd[4:]) // Push
					sizes[pwd[len(pwd)-1]] = 0
					dirs[cmd[4:]] = Directory{0, nil}
				} else {
					pwd = pwd[:len(pwd)-1] // Pop
				}
			}
		case regexp.MustCompile(`\d`).MatchString(string(cmd[0])):
			size, _ := strconv.Atoi(strings.Split(cmd, " ")[0])
			sizes[pwd[len(pwd)-1]] += size
			if size < 100000 {
				count += size
			}
		}
	}
	fmt.Println(count)
}
