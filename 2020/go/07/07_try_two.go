package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type BagOfBags map[string]map[string]int

func ruleParser(rule string) (string, map[string]int) {
	sliced_rule := strings.Split(rule, " contain ")
	bag := strings.Replace(sliced_rule[0], " bags", "", -1)
	bags := map[string]int{}
	re := regexp.MustCompile(`(\d) (\w* \w*)`)
	split_result := re.FindAllStringSubmatch(sliced_rule[1], -1)
	for _, b := range split_result {
		number, err := strconv.Atoi(b[1])
		check(err)
		bags[b[2]] = number
	}
	return bag, bags
}

func parseFile(path string) []string {
	file, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(file)
	rules := []string{}
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}
	return rules
}

func rulesParser(rules []string) BagOfBags {
	bag_of_bags := BagOfBags{}
	for _, rule := range rules {
		bag, bags := ruleParser(rule)
		bag_of_bags[bag] = bags
	}
	return bag_of_bags
}

func (b BagOfBags) show() {
	nicer, err := json.Marshal(b)
	check(err)
	fmt.Println(string(nicer))

}

//func (b BagOfBags) contains(bc string) (bool) {

//}

func inQueue(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func part1(b BagOfBags, bc string) int {
	bag_queue := []string{bc}
	bag_history := []string{}
	count := 0
	for len(bag_queue) > 0 {
		searched_bag := bag_queue[0]
		bag_queue = bag_queue[1:]
		for bag, bags := range b {
			if _, found := bags[searched_bag]; found {
				if !inQueue(bag, bag_history) && !inQueue(bag, bag_queue) {
					bag_queue = append(bag_queue, bag)
					bag_history = append(bag_history, bag)
					count++
				}
			}
		}
	}
	return count
}

//func part2(b BagOfBags, bc string) int {
//bag_queue := []string{bc}
//count := 0
//for len(bag_queue) > 0 {
//searched_bag := bag_queue[0]
//bag_queue = bag_queue[1:]
//for bag, bags := range b {
//if bag == searched_bag {
//for color, amount := range bags {
//bag_queue = append(bag_queue, color)
//count += amount

//}
//}
//}
//}
//}

func main() {
	input_rules := rulesParser(parseFile("input.txt"))
	//input_rules := rulesParser(parseFile("example.txt"))
	//input_rules := rulesParser(parseFile("example2.txt"))
	//input_rules.show()

	//example_rules.show()
	//result := part1(example_rules, "shiny gold")
	p1 := part1(input_rules, "shiny gold")
	fmt.Printf("Part 1: %d \n", p1)

}
