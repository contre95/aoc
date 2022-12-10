package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Motion struct {
	direction string
	module    int
}

func inputParser(filePath string) []Motion {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var input []Motion
	for fileScanner.Scan() {
		line := fileScanner.Text()
		v := strings.Split(line, " ")
		module, _ := strconv.Atoi(v[1])
		input = append(input, Motion{v[0], module})
	}
	return input
}

func main() {
	motions := inputParser(os.Args[1])
	ih := 0
	jh := 0
	it := 0
	jt := 0
	positions := map[string]bool{}
	ud := 1
	lr := 1
	for _, motion := range motions {
		//fmt.Println(motion.direction)
		for x := 0; x < motion.module; x++ {
			switch motion.direction {
			case "R":
				ih++
				lr = 1 // Save -1 or +1 to know last direction
				if math.Abs(float64(it)-float64(ih)) >= 2 {
					it++
					if math.Abs(float64(jt)-float64(jh)) >= 1 { // In case this is the first step after a direc change
						jt += ud // Increment/Decrement in one depending weather I was going up or down
					}
				}
			case "L":
				ih--
				lr = -1
				if math.Abs(float64(it)-float64(ih)) >= 2 {
					it--
					if math.Abs(float64(jt)-float64(jh)) >= 1 {
						jt += ud
					}
				}
			case "U":
				jh++
				ud = 1
				if math.Abs(float64(jt)-float64(jh)) >= 2 {
					jt++
					if math.Abs(float64(it)-float64(ih)) >= 1 {
						it += lr
					}
				}
			case "D":
				jh--
				ud = -1
				if math.Abs(float64(jt)-float64(jh)) >= 2 {
					jt--
					if math.Abs(float64(it)-float64(ih)) >= 1 {
						it += lr
					}
				}
			}
			positions[strconv.Itoa(it)+strconv.Itoa(jt)] = true // Save position
		}
	}
	fmt.Println(len(positions))
}
