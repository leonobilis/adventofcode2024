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

type Equation struct {
	Test    int
	Numbers []int
}

func parseInput(input string) (out []Equation) {
	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, ": ")
		test := Atoi(l[0])
		numbers := make([]int, 0)
		for _, n := range strings.Split(l[1], " ") {
			numbers = append(numbers, Atoi(n))
		}
		out = append(out, Equation{Test: test, Numbers: numbers})
	}
	return
}

func addOp(val, test int, numbers []int) bool {
	if len(numbers) == 0 {
		return val == test
	}
	return addOp(val+numbers[0], test, numbers[1:]) || mulOp(val+numbers[0], test, numbers[1:])
}

func mulOp(val, test int, numbers []int) bool {
	if len(numbers) == 0 {
		return val == test
	}
	return addOp(val*numbers[0], test, numbers[1:]) || mulOp(val*numbers[0], test, numbers[1:])
}

func p1(inp []Equation) (sum int) {
	for _, e := range inp {
		if addOp(0, e.Test, e.Numbers) || mulOp(0, e.Test, e.Numbers) {
			sum += e.Test
		}
	}
	return
}

func addOp2(val, test int, numbers []int) bool {
	if len(numbers) == 0 {
		return val == test
	}
	return addOp2(val+numbers[0], test, numbers[1:]) || mulOp2(val+numbers[0], test, numbers[1:]) || conOp(val+numbers[0], test, numbers[1:])
}

func mulOp2(val, test int, numbers []int) bool {
	if len(numbers) == 0 {
		return val == test
	}
	return addOp2(val*numbers[0], test, numbers[1:]) || mulOp2(val*numbers[0], test, numbers[1:]) || conOp(val*numbers[0], test, numbers[1:])
}

func conOp(val, test int, numbers []int) bool {
	if len(numbers) == 0 {
		return val == test
	}
	concat := Atoi(strconv.Itoa(val) + strconv.Itoa(numbers[0]))
	return addOp2(concat, test, numbers[1:]) || mulOp2(concat, test, numbers[1:]) || conOp(concat, test, numbers[1:])
}

func p2(inp []Equation) (sum int) {
	for _, e := range inp {
		if addOp2(0, e.Test, e.Numbers) || mulOp2(0, e.Test, e.Numbers) || conOp(0, e.Test, e.Numbers) {
			sum += e.Test
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
