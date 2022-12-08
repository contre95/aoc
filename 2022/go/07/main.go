package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"log"
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

type Tree map[string]*Dir

type Dir struct {
	Size int
	Tree Tree
}

func tprint(t Tree) {
	jsonStr, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(jsonStr))
	}
}

func main() {
	var pwd []string // Stack
	//var all map[string]bool
	dirs := map[string]int{}
	//root := Tree{}
	input := inputParser(os.Args[1])
	//var currentDir *Dir
	for _, cmd := range input {
		if cmd[:4] == "$ cd" {
			if cmd[5:] == ".." {
				pwd = pwd[:len(pwd)-1]
			} else if cmd[5:] == "/" {
				pwd = []string{}
			} else {
				pwd = append(pwd, cmd[5:])
			}
		} else if regexp.MustCompile(`\d`).MatchString(string(cmd[0])) { // If it starts with a number
			size, _ := strconv.Atoi(strings.Split(cmd, " ")[0])
			for i := 0; i < len(pwd)+1; i++ {
				path := "/" + strings.Join(pwd[:i], "/")
				dirs[path] += size
			}
		}
	}
	fmt.Println(dirs)
	count := 0
	for _, d := range dirs {
		if d < 100000 {
			count += d
		}
	}
	smallest := dirs["/"]
	spaceNeeded := 30000000 - (70000000 - dirs["/"])
	for _, d := range dirs {
		if d >= spaceNeeded && d <= smallest {
			smallest = d
		}

	}
	fmt.Println(count)
	fmt.Println(smallest)
}
