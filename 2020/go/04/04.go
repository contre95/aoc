package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func valid_byr(passport map[string]string) bool {
	if len(passport["byr"]) >= 1 {
		n, err := strconv.Atoi(string(passport["byr"]))
		check(err)
		return n >= 1920 && n <= 2002
	}
	return false
}
func valid_iyr(passport map[string]string) bool {
	if len(passport["iyr"]) >= 1 {
		n, err := strconv.Atoi(string(passport["iyr"]))
		check(err)
		return n >= 2010 && n <= 2020
	}
	return false
}
func valid_eyr(passport map[string]string) bool {
	if len(passport["eyr"]) >= 1 {
		n, err := strconv.Atoi(string(passport["eyr"]))
		check(err)
		return n >= 2020 && n <= 2030
	}
	return false
}
func valid_hcl(passport map[string]string) bool {
	s := []rune(string(passport["hcl"]))
	if string(s[0]) != "#" || len(s) != 7 {
		return false
	}
	return true
}

func valid_hgt(passport map[string]string) bool {
	s := []rune(string(passport["hgt"]))
	if len(s) == 0 {
		return false
	}
	unit := string(s[len(s)-2:])
	n, err := strconv.Atoi(string(s[0 : len(s)-2]))
	check(err)
	if unit == "cm" {
		return n >= 150 && n <= 193
	}
	if unit == "in" {
		return n >= 59 && n <= 76
	}
	return false
}

func valid_ecl(passport map[string]string) bool {
	valid_colors := map[string]bool{"brn": true, "amb": true, "gry": true, "grn": true, "hzl": true, "oth": true, "blu": true}
	if _, ok := valid_colors[string(passport["ecl"])]; ok {
		return true
	}
	return false
}

func valid_pid(passport map[string]string) bool {
	return len(passport["pid"]) == 9
}
func have_valid_fields(passport map[string]string) bool {
	for _, k := range []string{"ecl", "eyr", "hcl", "byr", "iyr", "hgt", "pid"} {
		if _, ok := passport[k]; !ok {
			return false
		}
	}
	return true
}

func have_valid_values(passport map[string]string) bool {
	if !valid_byr(passport) || !valid_eyr(passport) || !valid_iyr(passport) || !valid_pid(passport) || !valid_hcl(passport) || !valid_ecl(passport) || !valid_hgt(passport) {
		return false
	}
	return true
}

func main() {
	passports, err := fileToMap("input.txt")
	check(err)
	valid_passports1 := 0
	valid_passports2 := 0
	for _, p := range passports {
		//fmt.Printf(" %v -> %v \n", p, is_valid_passport(p))
		if have_valid_fields(p) {
			valid_passports1++
			if have_valid_values(p) {
				valid_passports2++
			}
		}
	}
	fmt.Printf("Part1: %d \n", valid_passports1)
	fmt.Printf("Part2: %d \n", valid_passports2)

}
