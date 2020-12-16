package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type BoardingPass struct {
	Rows string
	Cols string
}

func fileToSlice(fp string) ([]BoardingPass, error) {
	var lines []BoardingPass
	f, err := os.Open(fp)
	defer f.Close()
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		bp := BoardingPass{Rows: string(line[:7]), Cols: string(line[7:])}
		pointer_to_bp := &bp
		lines = append(lines, *pointer_to_bp)
	}

	return lines, scanner.Err()
}

func getPosition(bp string, max int, min int) int {
	upper := max
	lower := min
	for _, c := range bp {
		if c == 'F' || c == 'L' {
			upper -= ((upper - lower) / 2) + 1
		}
		if c == 'B' || c == 'R' {
			lower += ((upper - lower) / 2) + 1
		}
	}
	return upper
}

func part1(ids []int) int {
	return ids[len(ids)-1]
}

func part2(ids []int) (int, error) {
	for i, _ := range ids {
		if ids[i]+1 != ids[i+1] {
			return ids[i] + 1, nil
		}
	}
	return 0, errors.New("asdasd")

}

func main() {
	bps, err := fileToSlice("input.txt")
	check(err)
	var ids []int
	check(err)
	for _, bp := range bps {
		row := getPosition(bp.Rows, 127, 0)
		col := getPosition(bp.Cols, 7, 0)
		ids = append(ids, row*8+col)
	}
	sort.Ints(ids)
	fmt.Println(part1(ids))
	fmt.Println(part2(ids))
	//Retrun Max Id from ids
}
