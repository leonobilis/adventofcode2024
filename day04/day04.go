package main

import (
	"fmt"
	"os"
	"strings"
)

func p1(input []string) (sum int) {
	maxY, maxX := len(input)-1, len(input[0])-1

	check := func(a, b, c byte) bool {
		return a == 'M' && b == 'A' && c == 'S'
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if input[y][x] == 'X' {
				if x+3 <= maxX {
					if check(input[y][x+1], input[y][x+2], input[y][x+3]) {
						sum++
					}
					if y+3 <= maxY {
						if check(input[y+1][x+1], input[y+2][x+2], input[y+3][x+3]) {
							sum++
						}
					}
					if y-3 >= 0 {
						if check(input[y-1][x+1], input[y-2][x+2], input[y-3][x+3]) {
							sum++
						}
					}
				}
				if y+3 <= maxY {
					if check(input[y+1][x], input[y+2][x], input[y+3][x]) {
						sum++
					}
				}
				if y-3 >= 0 {
					if check(input[y-1][x], input[y-2][x], input[y-3][x]) {
						sum++
					}
				}

				if x-3 >= 0 {
					if check(input[y][x-1], input[y][x-2], input[y][x-3]) {
						sum++
					}
					if y+3 <= maxY {
						if check(input[y+1][x-1], input[y+2][x-2], input[y+3][x-3]) {
							sum++
						}
					}
					if y-3 >= 0 {
						if check(input[y-1][x-1], input[y-2][x-2], input[y-3][x-3]) {
							sum++
						}
					}
				}

			}
		}
	}
	return
}

func p2(input []string) (sum int) {
	maxY, maxX := len(input)-1, len(input[0])-1

	check := func(a, b byte) bool {
		return a == 'M' && b == 'S' || a == 'S' && b == 'M'
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if input[y][x] == 'A' {
				if x+1 <= maxX && x-1 >= 0 && y+1 <= maxY && y-1 >= 0 {
					if check(input[y+1][x+1], input[y-1][x-1]) && check(input[y-1][x+1], input[y+1][x-1]) {
						sum++
					}
				}
			}
		}
	}
	return
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Printf("Part 1: %v\n", p1(strings.Split(string(input), "\n")))
	fmt.Printf("Part 2: %v\n", p2(strings.Split(string(input), "\n")))
}
