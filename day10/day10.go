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

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	for v := range s2 {
		s.Add(v)
	}
	return s
}

type Point struct {
	x, y int
}

func parseInput(input string) (output [][]int32) {
	for _, s := range strings.Split(input, "\n") {
		var temp []int32
		for _, s2 := range s {
			temp = append(temp, s2-'0')
		}
		output = append(output, temp)
	}
	return
}

func traverse(inp [][]int32, val int32, x, y, maxX, maxY int) Set[Point] {
	result := make(Set[Point])

	if val == 9 {
		result.Add(Point{x, y})
		return result
	}

	if x > 0 && inp[y][x-1] == val+1 {
		result.Union(traverse(inp, val+1, x-1, y, maxX, maxY))
	}
	if x < maxX && inp[y][x+1] == val+1 {
		result.Union(traverse(inp, val+1, x+1, y, maxX, maxY))
	}
	if y > 0 && inp[y-1][x] == val+1 {
		result.Union(traverse(inp, val+1, x, y-1, maxX, maxY))
	}
	if y < maxY && inp[y+1][x] == val+1 {
		result.Union(traverse(inp, val+1, x, y+1, maxX, maxY))
	}
	return result
}

func p1(inp [][]int32) (sum int) {
	for y, row := range inp {
		for x, cell := range row {
			if cell == 0 {
				sum += len(traverse(inp, cell, x, y, len(row)-1, len(inp)-1))
			}
		}
	}
	return
}

func traverse2(inp [][]int32, val int32, x, y, maxX, maxY int) (sum int) {
	if val == 9 {
		return 1
	}

	if x > 0 && inp[y][x-1] == val+1 {
		sum += traverse2(inp, val+1, x-1, y, maxX, maxY)
	}
	if x < maxX && inp[y][x+1] == val+1 {
		sum += traverse2(inp, val+1, x+1, y, maxX, maxY)
	}
	if y > 0 && inp[y-1][x] == val+1 {
		sum += traverse2(inp, val+1, x, y-1, maxX, maxY)
	}
	if y < maxY && inp[y+1][x] == val+1 {
		sum += traverse2(inp, val+1, x, y+1, maxX, maxY)
	}
	return
}

func p2(inp [][]int32) (sum int) {
	for y, row := range inp {
		for x, cell := range row {
			if cell == 0 {
				sum += traverse2(inp, cell, x, y, len(row)-1, len(inp)-1)
			}
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
