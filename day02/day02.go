package main

import (
	"fmt"
	"os"
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

func parseInput(input string) (output [][]int) {
	for _, s := range strings.Split(input, "\n") {
		var temp []int
		for _, s2 := range strings.Split(s, " ") {
			temp = append(temp, Atoi(s2))
		}
		output = append(output, temp)
	}
	return
}

func p1(inp [][]int) (sum int) {
loop:
	for y := 0; y < len(inp); y++ {
		mod := 1
		if inp[y][0]-inp[y][1] < 0 {
			mod = -1
		}
		for x := 0; x < len(inp[y])-1; x++ {
			diff := mod * (inp[y][x] - inp[y][x+1])
			if diff < 1 || diff > 3 {
				continue loop
			}
		}
		sum++
	}
	return
}

func p2(inp [][]int) (sum int) {
loop:
	for y := 0; y < len(inp); y++ {
		mod := 1
		if inp[y][0]-inp[y][1] < 0 {
			mod = -1
		}
		bad := false
		for x := 0; x < len(inp[y])-1; x++ {
			diff := mod * (inp[y][x] - inp[y][x+1])
			if diff < 1 || diff > 3 {
				if bad {
					continue loop
				}
				bad = true
			}
		}
		sum++
	}
	return
}

func main() {
	input, _ := os.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
