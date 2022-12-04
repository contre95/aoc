package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func main() {
	p := os.Args[1]
	input := fileToA(p)
	calls := []int{}
	elfCals := 0
	for _, v := range input {
		if v == "" {
			calls = append(calls, elfCals)
			elfCals = 0
		} else {
			cals, _ := strconv.Atoi(v)
			elfCals += cals
		}
	}
	calls = append(calls, elfCals)
	sort.Ints(calls)
	fmt.Println(calls[len(calls)-1])
	fmt.Println(calls[len(calls)-3] + calls[len(calls)-2] + calls[len(calls)-1])
}
