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
	a := map[string][]int{"R": {1, 0}, "L": {-1, 0}, "U": {0, 1}, "D": {0, -1}}
	motions := inputParser(os.Args[1])
	fmt.Println(len(motions))
	ih, jh, it, jt := 0, 0, 0, 0
	positions := map[string]bool{}
	for _, motion := range motions {
		for x := 0; x < motion.module; x++ {
			ih += a[motion.direction][0]
			jh += a[motion.direction][1]
			touching := math.Abs(float64(jt)-float64(jh)) <= 1 && math.Abs(float64(it)-float64(ih)) <= 1
			if !touching {
				if it != ih {
					it += (ih - it) / int(math.Abs(float64(ih)-float64(it)))
				}
				if jt != jh {
					jt += (jh - jt) / int(math.Abs(float64(jh)-float64(jt)))
				}
			}
			positions[strconv.Itoa(it)+"-"+strconv.Itoa(jt)] = true
		}
	}
	fmt.Println(len(positions))
}
