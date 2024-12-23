package main

import (
	"fmt"
	"os"
	"slices"
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

func (s Set[T]) Clone() Set[T] {
	result := make(Set[T])
	for v := range s {
		result.Add(v)
	}
	return result
}

func parseInput(input string) (output map[string]Set[string]) {
	output = make(map[string]Set[string])
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, "-")
		if _, ok := output[s2[0]]; !ok {
			output[s2[0]] = make(Set[string])
		}
		output[s2[0]].Add(s2[1])
		if _, ok := output[s2[1]]; !ok {
			output[s2[1]] = make(Set[string])
		}
		output[s2[1]].Add(s2[0])
	}
	return
}

func p1(inp map[string]Set[string]) int {
	interConnected := make(Set[[3]string])
	for k, v := range inp {
		for k2 := range v {
			for k3 := range inp[k2] {
				if inp[k3].Contains(k) && (k[0] == 't' || k2[0] == 't' || k3[0] == 't') {
					key := [3]string{k, k2, k3}
					slices.Sort(key[:])
					interConnected.Add(key)
				}
			}

		}
	}
	return len(interConnected)
}

func traverse(inp map[string]Set[string], current, visited Set[string], next, end string) Set[string] {
	if next == end {
		current.Add(end)
		return current
	}

	if visited.Contains(next) {
		return nil
	}
	visited.Add(next)

	check, ok := inp[next]
	if !ok {
		return nil
	}

	if !check.Contains(end) {
		return nil
	}

	for k := range current {
		if !check.Contains(k) {
			return nil
		}
	}

	maxLen := 0
	var result Set[string]
	current.Add(next)

	for k := range check {
		if current.Contains(k) {
			continue
		}
		res := traverse(inp, current.Clone(), visited, k, end)
		if len(res) > maxLen {
			maxLen = len(res)
			result = res
		}
	}

	return result
}

func p2(inp map[string]Set[string]) string {
	var result Set[string]
	maxLen := 0

	for k, v := range inp {
		delete(inp, k)
		for k2 := range v {
			res := traverse(inp, make(Set[string]), make(Set[string]), k2, k)
			if len(res) > maxLen {
				maxLen = len(res)
				result = res
			}
		}
	}

	list := make([]string, 0, len(result))
	for k := range result {
		list = append(list, k)
	}
	slices.Sort(list)
	return strings.Join(list, ",")
}

func main() {
	input, _ := os.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
