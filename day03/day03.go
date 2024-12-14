package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

type Args struct {
	a, b    int
	enabled bool
}

func parseInput(input string) (out []Args) {
	re := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	instructions := re.FindAllStringSubmatch(input, -1)
	enabled := true
	for _, instr := range instructions {
		if instr[0] == "do()" {
			enabled = true
		} else if instr[0] == "don't()" {
			enabled = false
		} else {
			out = append(out, Args{Atoi(instr[2]), Atoi(instr[3]), enabled})
		}

	}
	return
}

func p1(inp []Args) (sum int) {
	for _, args := range inp {
		sum += args.a * args.b
	}
	return
}

func p2(inp []Args) (sum int) {
	for _, args := range inp {
		if args.enabled {
			sum += args.a * args.b
		}
	}
	return
}

func main() {
	input, _ := os.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
