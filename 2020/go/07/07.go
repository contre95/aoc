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

func ruleParser(rule string) map[string][]Bag {
	parsedRule := map[string][]Bag{}
	sliced_rule := strings.Split(rule, " contain ")
	sliced_rule[0] = strings.Replace(sliced_rule[0], "bags", "", -1)
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
	:wq	check(err)
	scanner := bufio.NewScanner(file)
	rules := []map[string][]Bag{}
	for scanner.Scan() {
		rules = append(rules, ruleParser(scanner.Text()))
	}
	return rules
}

func main() {
	rules := rulesParser("example.txt")
	rulesJSON, err := json.Marshal(rules)
	 check(err)
	fmt.Println(string(rulesJSON))

	//fmt.Printf("%+v \n", rules)
	//fmt.Printf("%#v \n", rules)
}
