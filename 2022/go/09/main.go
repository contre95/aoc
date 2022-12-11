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
	a := map[string]int{"R": 1, "L": -1, "U": 1, "D": -1}
	motions := inputParser(os.Args[1])
	ih := 0
	jh := 0
	it := 0
	jt := 0
	positions := map[string]bool{}
	var ud, lr int
	fmt.Println(ud, lr)
	//fmt.Println(positions)
	for _, motion := range motions {
		for x := 0; x < motion.module; x++ {
			// Update H
			if motion.direction == "R" || motion.direction == "L" {
				ih += a[motion.direction]
				lr = a[motion.direction] // Save -1 or +1 to know last direction
			} else {
				jh += a[motion.direction]
				ud = a[motion.direction]
			}
			// Update T
			touching := math.Abs(float64(jt)-float64(jh)) <= 1 && math.Abs(float64(it)-float64(ih)) <= 1
			if !touching {
				if math.Abs(float64(jt)-float64(jh)) >= 2 {
					if math.Abs(float64(it)-float64(ih)) >= 1 {
						it += lr
					}
					jt += a[motion.direction]
				}
				if math.Abs(float64(it)-float64(ih)) >= 2 {
					if math.Abs(math.Abs(float64(jt)-float64(jh))) >= 1 {
						jt += ud // Increment/Decrement in one depending weather I was going up or down
					}
					it += a[motion.direction]
				}
			}
				positions[strconv.Itoa(it)+strconv.Itoa(jt)] = true
		}
	}
	if os.Args[1] == "sample" {
		for a := 0; a < 6; a++ {
			for b := 0; b < 5; b++ {
				if positions[strconv.Itoa(a)+strconv.Itoa(b)] {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
	}

	fmt.Println(len(positions))
}
