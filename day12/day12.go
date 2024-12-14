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

type Point struct {
	x, y int
}

const (
	Right = iota
	Down
	Left
	Up
)

type Side struct {
	x, y, dir int
}

func parseInput(input string) (output []string) {
	for _, s := range strings.Split(input, "\n") {
		output = append(output, s)
	}
	return
}

func traverse(inp []string, x, y int, letter byte, visited Set[Point], sides Set[Side]) (area int) {
	if visited.Contains(Point{x, y}) {
		return
	}
	area = 1
	visited.Add(Point{x, y})
	if x > 0 && inp[y][x-1] == letter {
		area += traverse(inp, x-1, y, letter, visited, sides)
	} else {
		sides.Add(Side{x - 1, y, Left})
	}
	if x < len(inp[y])-1 && inp[y][x+1] == letter {
		area += traverse(inp, x+1, y, letter, visited, sides)
	} else {
		sides.Add(Side{x + 1, y, Right})
	}
	if y > 0 && inp[y-1][x] == letter {
		area += traverse(inp, x, y-1, letter, visited, sides)
	} else {
		sides.Add(Side{x, y - 1, Up})
	}
	if y < len(inp)-1 && inp[y+1][x] == letter {
		area += traverse(inp, x, y+1, letter, visited, sides)
	} else {
		sides.Add(Side{x, y + 1, Down})
	}
	return
}

func reduceSides(sides Set[Side]) int {
	for s := range sides {
		if s.dir == Up || s.dir == Down {
			for x := s.x - 1; sides.Contains(Side{x, s.y, s.dir}); x-- {
				delete(sides, Side{x, s.y, s.dir})
			}
			for x := s.x + 1; sides.Contains(Side{x, s.y, s.dir}); x++ {
				delete(sides, Side{x, s.y, s.dir})
			}
		} else {
			for y := s.y - 1; sides.Contains(Side{s.x, y, s.dir}); y-- {
				delete(sides, Side{s.x, y, s.dir})
			}
			for y := s.y + 1; sides.Contains(Side{s.x, y, s.dir}); y++ {
				delete(sides, Side{s.x, y, s.dir})
			}
		}
	}
	return len(sides)
}

func p1(inp []string) (sum int) {
	visited := make(Set[Point])

	for y, s := range inp {
		for x, s2 := range s {
			if !visited.Contains(Point{x, y}) {
				sides := make(Set[Side])
				a := traverse(inp, x, y, byte(s2), visited, sides)
				sum += a * len(sides)
			}
		}
	}
	return
}

func p2(inp []string) (sum int) {
	visited := make(Set[Point])

	for y, s := range inp {
		for x, s2 := range s {
			if !visited.Contains(Point{x, y}) {
				sides := make(Set[Side])
				a := traverse(inp, x, y, byte(s2), visited, sides)
				sum += a * reduceSides(sides)
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
