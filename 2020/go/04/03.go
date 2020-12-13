package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToMap(fp string) ([]map[string]string, error) {
	var s []map[string]string
	f, err := os.Open(fp)
	check(err)
	scanner := bufio.NewScanner(f)
	m := make(map[string]string)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), " ")
		if len(l) == 1 && l[0] == "" {
			//fmt.Println(m)
			s = append(s, m)
			m = make(map[string]string)
			continue
		}
		for _, v := range l {
			a := strings.Split(v, ":")
			m[a[0]] = a[1]
		}
	}
	s = append(s, m)
	defer f.Close()
	return s, scanner.Err()
}

func is_valid_passport(passport map[string]string) bool {
	for _, k := range []string{"ecl", "eyr", "hcl", "byr", "iyr", "hgt", "pid"} {
		if _, ok := passport[k]; !ok {
			return false
		}
	}
	return true
}

func main() {
	passports, err := fileToMap("input.txt")
	check(err)
	valid_passports := 0
	for _, p := range passports {
		//fmt.Printf(" %v -> %v \n", p, is_valid_passport(p))
		if is_valid_passport(p) {
			valid_passports++
		}
	}
	fmt.Println(valid_passports)

}
