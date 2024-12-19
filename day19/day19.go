package main

import (
	"fmt"
	"os"
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

func parseInput(input string) (Set[string], int, []string) {
	s := strings.Split(input, "\n\n")

	patterns := make(Set[string])
	maxLen := 0
	for _, s2 := range strings.Split(s[0], ", ") {
		if len(s2) > maxLen {
			maxLen = len(s2)
		}
		patterns.Add(s2)
	}

	return patterns, maxLen, strings.Split(s[1], "\n")
}

func check2(patterns, antipatterns Set[string], cache map[string]int, maxLen int, design string) (count int) {
	if antipatterns.Contains(design) {
		return 0
	}

	if c, ok := cache[design]; ok {
		return c
	}

	j := maxLen
	if j > len(design) {
		j = len(design)
	}

	for ; j > 0; j-- {
		if patterns.Contains(design[:j]) {
			if j == len(design) {
				count++
				continue
			}

			count += check2(patterns, antipatterns, cache, maxLen, design[j:])
		}
	}

	if count > 0 {
		cache[design] = count
	} else {
		antipatterns.Add(design)
	}

	return
}

func p1(patterns Set[string], maxLen int, designs []string) (sum int) {
	antipatterns := make(Set[string])
	cache := make(map[string]int)

	for _, design := range designs {
		if check2(patterns, antipatterns, cache, maxLen, design) > 0 {
			sum++
		}
	}
	return
}

func p2(patterns Set[string], maxLen int, designs []string) (sum int) {
	antipatterns := make(Set[string])
	cache := make(map[string]int)

	for _, design := range designs {
		result := check2(patterns, antipatterns, cache, maxLen, design)
		sum += result
	}
	return
}

func main() {
	input, _ := os.ReadFile("input.txt")
	patterns, maxLen, designs := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(patterns, maxLen, designs))
	fmt.Printf("Part 2: %v\n", p2(patterns, maxLen, designs))
}
