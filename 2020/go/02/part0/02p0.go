package part0

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type Line struct {
	Constraints []int
	Letter      string
	Password    string
}

func FileToMap(fp string) ([]Line, error) {
	var lines []Line
	f, err := os.Open(fp)
	Check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		s := strings.Split(l, " ")
		s1 := strings.Split(s[0], "-")
		min, err := strconv.Atoi(s1[0])
		Check(err)
		max, err := strconv.Atoi(s1[1])
		Check(err)
		a := Line{[]int{min, max}, strings.TrimSuffix(s[1], ":"), s[2]}
		lines = append(lines, a)
	}
	defer f.Close()
	return lines, nil
}
