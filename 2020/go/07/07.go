package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	BagAmount int
	BagColor  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func containsBag(color string, bags []Bag) bool {
	for _, b := range bags {
		//fmt.Printf("%s is %s  \n ", b.BagColor[0:], a[0:])
		//fmt.Printf("%d is %d  \n ", len(b.BagColor[0:]), len(a[0:]))
		if b.BagColor == color {
			return true
		}
	}
	return false
}

func ruleParser(rule string) map[string][]Bag {
	parsedRule := map[string][]Bag{}
	sliced_rule := strings.Split(rule, " contain ")
	sliced_rule[0] = strings.Replace(sliced_rule[0], " bags", "", -1)
	bags := []Bag{}
	for _, bag := range strings.Split(sliced_rule[1], ", ") {
		number_bag := strings.Split(bag, " ")
		if number_bag[0] != "no" {
			amount, err := strconv.Atoi(number_bag[0])
			check(err)
			b := Bag{
				BagAmount: amount,
				BagColor:  number_bag[1] + " " + number_bag[2],
			}
			bags = append(bags, b)
		}
	}
	parsedRule[sliced_rule[0]] = bags
	return parsedRule
}


func rulesParser(path string) []map[string][]Bag {
	file, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(file)
	rules := []map[string][]Bag{}
	for scanner.Scan() {
		rules = append(rules, ruleParser(scanner.Text()))
	}
	return rules
}

func part1(rules []map[string][]Bag, color string) int {
	queue := []string{}
	queue = append(queue, color)
	c := 0
	slice := []string{}
	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:]
		for _, m := range rules {
			for k, v := range m {
				//fmt.Printf(" %v in %s -> %v ? \n", b, k, v)
				if containsBag(b, v) {
					//fmt.Println("YES")
					if !inQueue(k, queue) && !inQueue(k, slice) {
						queue = append(queue, k)
						slice = append(slice, k)
						c++
					}
				}
			}
		}
	}
	return c
}

func main() {
	rules := rulesParser("example.txt")
	//rules := rulesParser("input.txt")
	rulesJSON, err := json.Marshal(rules)
	check(err)
	fmt.Printf("%+v \n", string(rulesJSON))

	//p1 := part1(rules, "shiny gold")
	if _, ok :=
	//fmt.Printf("\n", p1)
	//fmt.Printf("%#v \n", rules)
}
