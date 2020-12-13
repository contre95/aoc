package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToSlice(fp string) ([]string, error) {
	var lines []string
	f, err := os.Open(fp)
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	defer f.Close()
	return lines, scanner.Err()
}

func count_trees(path []string, du int, rl int) int {
	l1 := len(path[0])
	l2 := len(path)
	j := 0
	i := 0
	trees := 0
	for {
		if j+du >= l2 {
			return trees
		}
		i = i + rl
		if i >= l1 {
			i = i - l1
		}
		j = j + du
		if string(path[j][i]) == "#" {
			trees++
		}
	}
}

func main() {
	path, err := fileToSlice("input.txt")
	check(err)
	trees := count_trees(path, 1, 3)
	fmt.Printf("Number of trees in the path: %d \n", trees)
	slopes := [][]int{{1, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, s := range slopes {
		trees *= count_trees(path, s[1], s[0])
	}
	fmt.Printf("Number of trees in all paths: %d \n", trees)
}
