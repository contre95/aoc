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
	files := map[string]int{}
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
				files[path] += size
			}
		}
	}
	fmt.Println(files)
	count := 0
	for _, s := range files {
		if s < 100000 {
			count += s
		}
	}
	fmt.Println(count)
}
