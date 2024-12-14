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

func Mod(a, b int) int {
	return (a%b + b) % b
}

type Position struct {
	x, y int
}

type Direction int

const (
	Right = iota
	Down
	Left
	Up
)

func turnRight(d Direction) Direction {
	return Direction(Mod(int(d+1), 4))
}

type Grid struct {
	m          Set[Position]
	guard, max Position
}

func (g *Grid) IsEdge(position Position) bool {
	return position.x == 0 || position.y == 0 || position.x == g.max.x || position.y == g.max.y
}

type PosDir struct {
	p Position
	d Direction
}

type Guard struct {
	p       Position
	d       Direction
	visited Set[PosDir]
}

func NewGuard(p Position) *Guard {
	v := make(Set[PosDir])
	v.Add(PosDir{p, Up})
	return &Guard{p: p, d: Up, visited: v}
}

func (g *Guard) TurnRight() {
	g.d = turnRight(g.d)
}

func (g *Guard) Move(check func(Position) bool) (Position, bool) {
	var newPos Position
	switch g.d {
	case Right:
		newPos = Position{g.p.x + 1, g.p.y}
	case Down:
		newPos = Position{g.p.x, g.p.y + 1}
	case Left:
		newPos = Position{g.p.x - 1, g.p.y}
	case Up:
		newPos = Position{g.p.x, g.p.y - 1}
	}

	if check(newPos) {
		g.p = newPos
	} else {
		g.TurnRight()
	}

	pd := PosDir{g.p, g.d}
	if g.visited.Contains(pd) {
		return g.p, false
	}

	g.visited.Add(pd)
	return g.p, true
}

func (g *Guard) NumVisited() int {
	visited := make(Set[Position])
	for v := range g.visited {
		visited.Add(v.p)
	}
	return len(visited)
}

func parseInput(input string) *Grid {
	lines := strings.Split(input, "\n")
	output := &Grid{m: make(Set[Position]), max: Position{len(lines[0]) - 1, len(lines) - 1}}
	for y, s := range lines {
		for x, s2 := range s {
			if s2 == '#' {
				output.m.Add(Position{x, y})
			} else if s2 == '^' {
				output.guard = Position{x, y}
			}
		}
	}
	return output
}

func p1(inp *Grid) int {
	guard := NewGuard(inp.guard)

	for {
		if p, _ := guard.Move(func(p Position) bool {
			return !inp.m.Contains(p)
		}); inp.IsEdge(p) {
			return guard.NumVisited()
		}
	}
}

func p2(inp *Grid) (sum int) {
	for y := 0; y <= inp.max.y; y++ {
		for x := 0; x <= inp.max.x; x++ {
			if inp.m.Contains(Position{x, y}) {
				continue
			}
			guard := NewGuard(inp.guard)
			for {
				p, ok := guard.Move(func(p Position) bool {
					return p != Position{x, y} && !inp.m.Contains(p)
				})
				if !ok {
					sum++
					break
				} else if inp.IsEdge(p) {
					break
				}
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
