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

func parseInputToSlice(path string) ([][]string, error) {
	//[ ['abc'], ['a','b','c'] ]
	groups := [][]string{}
	group_ans := []string{}
	f, err := os.Open(path)
	defer f.Close()
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			groups = append(groups, group_ans)
			group_ans = []string{}
		} else {
			person_ans := scanner.Text()
			group_ans = append(group_ans, person_ans)
		}
	}
	return groups, nil
}

func sum(array [26]int, amount_people int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	//Part II
	//for _, v := range array {
	//if v == amount_people {
	//result += 1
	//}
	//}
	return result
}

func getPositiveAnswers(group_answers []string) int {
	var different_answers [26]int
	for _, answer := range group_answers {
		for _, c := range answer {
			different_answers[int(c)-int('a')] = 1
			//different_answers2[int(c)-int('a')] += 1 // Part II
		}
	}
	return sum(different_answers, len(group_answers))
}

func main() {
	groups_answers, err := parseInputToSlice("input.txt")
	//groups_answers, err := parseInputToSlice("example.txt")
	check(err)
	count := 0
	for _, group_answers := range groups_answers {
		count += getPositiveAnswers(group_answers)
	}
	fmt.Printf("Parte 1 : %d \n", count)

}
