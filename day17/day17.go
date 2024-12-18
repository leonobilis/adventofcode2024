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

func parseInput(input string) (a, b, c int, program []int) {
	s := strings.Split(input, "\n\n")
	s2 := strings.Split(s[0], "\n")
	a = Atoi(s2[0][12:])
	b = Atoi(s2[1][12:])
	c = Atoi(s2[2][12:])

	s3 := strings.Split(s[1][9:], ",")
	for _, s4 := range s3 {
		program = append(program, Atoi(s4))
	}
	return
}

func run(a, b, c int, program []int) (out []int) {

	val := func(o int) int {
		switch o {
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		default:
			return o

		}
	}

	for pointer := 0; pointer < len(program); {
		switch program[pointer] {
		case 0:
			a = a >> val(program[pointer+1])
		case 1:
			b = b ^ val(program[pointer+1])
		case 2:
			b = val(program[pointer+1]) % 8
		case 3:
			if a > 0 {
				pointer = val(program[pointer+1])
				continue
			}
		case 4:
			b = b ^ c
		case 5:
			out = append(out, val(program[pointer+1])%8)
		case 6:
			b = a >> val(program[pointer+1])
		case 7:
			c = a >> val(program[pointer+1])
		}
		pointer += 2
	}
	return
}

func p1(a, b, c int, program []int) string {
	out := run(a, b, c, program)
	strOut := make([]string, 0)
	for _, o := range out {
		strOut = append(strOut, strconv.Itoa(o))
	}
	return strings.Join(strOut, ",")
}

func p2(a, b, c int, program []int) int {
	toCheck := []int{0}
	for i := len(program) - 1; i >= 0; i-- {
		newCheck := make([]int, 0)
		for j := 0; j <= 8; j++ {
			for _, result := range toCheck {
				if slices.Equal(run(result*8+j, b, c, program), program[i:]) {
					newCheck = append(newCheck, result*8+j)
				}
			}
		}
		toCheck = newCheck
	}

	min := toCheck[0]
	for _, res := range toCheck {
		if res < min {
			min = res
		}
	}
	return min
}

func main() {
	input, _ := os.ReadFile("input.txt")
	a, b, c, program := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(a, b, c, program))
	fmt.Printf("Part 2: %v\n", p2(a, b, c, program))
}
