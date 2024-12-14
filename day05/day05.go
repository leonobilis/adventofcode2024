package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func Atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func parseInput(input string) (rules, rulesReverse map[int]Set[int], updates [][]int) {
	rules = make(map[int]Set[int])
	rulesReverse = make(map[int]Set[int])
	ru := strings.Split(input, "\n\n")
	for _, r := range strings.Split(ru[0], "\n") {
		l := strings.Split(r, "|")
		if _, ok := rules[Atoi(l[0])]; !ok {
			rules[Atoi(l[0])] = make(Set[int])
		}
		rules[Atoi(l[0])].Add(Atoi(l[1]))
		if _, ok := rulesReverse[Atoi(l[1])]; !ok {
			rulesReverse[Atoi(l[1])] = make(Set[int])
		}
		rulesReverse[Atoi(l[1])].Add(Atoi(l[0]))
	}
	for _, u := range strings.Split(ru[1], "\n") {
		update := make([]int, 0)
		for _, u2 := range strings.Split(u, ",") {
			update = append(update, Atoi(u2))
		}
		updates = append(updates, update)
	}
	return
}

func p1(rules map[int]Set[int], updates [][]int) (sum int) {
	for _, update := range updates {
		check := make([]int, 0)
		valid := true
	CheckLoop:
		for _, u := range update {
			if rule, ok := rules[u]; ok {
				for _, c := range check {
					if rule.Contains(c) {
						valid = false
						break CheckLoop
					}
				}
			}
			check = append(check, u)
		}
		if valid {
			sum += update[len(update)/2]
		}
	}
	return
}

func reorder(rulesReverse map[int]Set[int], update []int) (result []int) {
	lenUpdate := len(update)
	for len(result) < lenUpdate {
		for i, u := range update {
			valid := true
			if rule, ok := rulesReverse[u]; ok {
				for j, u2 := range update {
					if i != j && rule.Contains(u2) {
						valid = false
						break
					}
				}
			}
			if valid {
				result = append(result, u)
				update = append(update[:i], update[i+1:]...)
				break
			}
		}
	}
	return
}

func p2(rules, rulesReverse map[int]Set[int], updates [][]int) (sum int) {
	for _, update := range updates {
		check := make([]int, 0)
		valid := true
	CheckLoop:
		for _, u := range update {
			if rule, ok := rules[u]; ok {
				for _, c := range check {
					if rule.Contains(c) {
						valid = false
						break CheckLoop
					}
				}
			}
			check = append(check, u)
		}
		if !valid {
			newUpdate := reorder(rulesReverse, update)
			sum += newUpdate[len(newUpdate)/2]
		}
	}
	return
}

func main() {
	input, _ := os.ReadFile("input.txt")
	rules, rulesReverse, updates := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(rules, updates))
	fmt.Printf("Part 2: %v\n", p2(rules, rulesReverse, updates))
}
