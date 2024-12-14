package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseInput(input string) [][]int {
	var l1, l2 []int
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, "   ")
		l1 = append(l1, Atoi(s2[0]))
		l2 = append(l2, Atoi(s2[1]))
	}
	return [][]int{l1, l2}
}

func p1(inp [][]int) (sum int) {
	slices.Sort(inp[0])
	slices.Sort(inp[1])

	for i := 0; i < len(inp[0]); i++ {
		sum += Abs(inp[0][i] - inp[1][i])
	}

	return
}

func p2(inp [][]int) (sum int) {
	l2 := make(map[int]int)
	for _, l := range inp[1] {
		l2[l]++
	}

	for _, l := range inp[0] {
		sum += l2[l] * l
	}

	return
}

func main() {
	input, _ := os.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
