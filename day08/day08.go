package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

type Grid struct {
	m   map[int32][]Position
	max Position
}

func (g *Grid) Add(key int32, position Position) {
	g.m[key] = append(g.m[key], position)
}

func (g *Grid) IsAvalid(position Position) bool {
	return position.x >= 0 && position.y >= 0 && position.x <= g.max.x && position.y <= g.max.y
}

func parseInput(input string) *Grid {
	lines := strings.Split(input, "\n")
	output := &Grid{m: make(map[int32][]Position), max: Position{len(lines[0]) - 1, len(lines) - 1}}
	for y, s := range lines {
		for x, s2 := range s {
			if s2 >= 'A' && s2 <= 'Z' || s2 >= 'a' && s2 <= 'z' || s2 >= '0' && s2 <= '9' {
				output.Add(s2, Position{x, y})
			}
		}
	}
	return output
}

func p1(inp *Grid) (sum int) {
	antinodes := make(Set[Position])
	printMap := make(map[Position]string)

	for a, freq := range inp.m {
		for _, a1 := range freq {
			printMap[a1] = string(a)
			for _, a2 := range freq {
				if a1 != a2 {
					diffX := a1.x - a2.x
					diffY := a1.y - a2.y
					if Abs(diffX)+Abs(diffY) > 1 {
						antinode1 := Position{a1.x + diffX, a1.y + diffY}
						if inp.IsAvalid(antinode1) {
							antinodes.Add(antinode1)
						}
						antinode2 := Position{a2.x - diffX, a2.y - diffY}
						if inp.IsAvalid(antinode2) {
							antinodes.Add(antinode2)
						}
					}
				}
			}
		}
	}

	return len(antinodes)
}

func p2(inp *Grid) (sum int) {
	antinodes := make(Set[Position])

	printMap := make(map[Position]string)

	for a, freq := range inp.m {
		for _, a1 := range freq {
			antinodes.Add(a1)
			printMap[a1] = string(a)
			for _, a2 := range freq {
				if a1 != a2 {
					diffX := a1.x - a2.x
					diffY := a1.y - a2.y
					p := Position{a1.x + diffX, a1.y + diffY}
					for inp.IsAvalid(p) {
						antinodes.Add(p)
						p = Position{p.x + diffX, p.y + diffY}
					}
					p = Position{a2.x - diffX, a2.y - diffY}
					for inp.IsAvalid(p) {
						antinodes.Add(p)
						p = Position{p.x - diffX, p.y - diffY}
					}
				}
			}
		}
	}

	return len(antinodes)
}
func main() {
	input, _ := os.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
