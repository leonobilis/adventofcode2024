package main

import (
	"fmt"
	"os"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)

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

func (s Set[T]) Copy() Set[T] {
	s2 := make(Set[T])
	for v := range s {
		s2.Add(v)
	}
	return s2
}

func (s Set[T]) Clear() {
	for v := range s {
		delete(s, v)
	}
}

type Point struct {
	x, y int
}

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

func Mod(a, b int) int {
	return (a%b + b) % b
}

func turnLeft(d Direction) Direction {
	return Direction(Mod(int(d-1), 4))
}

func turnRight(d Direction) Direction {
	return Direction(Mod(int(d+1), 4))
}

func parseInput(input string) (grid Set[Point], start, end Point) {
	grid = make(Set[Point])
	for y, s := range strings.Split(input, "\n") {
		for x, s2 := range s {
			if s2 != '#' {
				grid.Add(Point{x, y})
			}
			if s2 == 'S' {
				start = Point{x, y}
			} else if s2 == 'E' {
				end = Point{x, y}
			}
		}
	}
	return
}

type cacheKey struct {
	pos Point
	dir Direction
}

func traverse(grid Set[Point], pos, end Point, dir Direction, points int, cache map[cacheKey]int) int {
	if pos == end {
		return points
	}

	if v, ok := cache[cacheKey{pos, dir}]; ok {
		if points > v {
			return MaxInt
		}
	}
	cache[cacheKey{pos, dir}] = points

	var move Point

	switch dir {
	case Right:
		move = Point{pos.x + 1, pos.y}
	case Down:
		move = Point{pos.x, pos.y + 1}
	case Left:
		move = Point{pos.x - 1, pos.y}
	case Up:
		move = Point{pos.x, pos.y - 1}
	}

	a := MaxInt
	if grid.Contains(move) {
		a = traverse(grid, move, end, dir, points+1, cache)
	}

	b := traverse(grid, pos, end, turnLeft(dir), points+1000, cache)
	c := traverse(grid, pos, end, turnRight(dir), points+1000, cache)

	return min(a, b, c)

}

func p1(grid Set[Point], start, end Point) int {
	return traverse(grid, start, end, Right, 0, make(map[cacheKey]int))
}

func traverse2(grid Set[Point], pos, end Point, dir Direction, points, expectedScore int, cache map[cacheKey]int, path, global Set[Point]) {
	if points > expectedScore {
		return
	}
	path.Add(pos)
	if pos == end {
		if points == expectedScore {
			global.Union(path)
		}
		return
	}

	if v, ok := cache[cacheKey{pos, dir}]; ok {
		if points > v {
			return
		}
	}
	cache[cacheKey{pos, dir}] = points

	path.Add(pos)

	var move Point

	switch dir {
	case Right:
		move = Point{pos.x + 1, pos.y}
	case Down:
		move = Point{pos.x, pos.y + 1}
	case Left:
		move = Point{pos.x - 1, pos.y}
	case Up:
		move = Point{pos.x, pos.y - 1}
	}

	if grid.Contains(move) {
		traverse2(grid, move, end, dir, points+1, expectedScore, cache, path.Copy(), global)
	}

	traverse2(grid, pos, end, turnLeft(dir), points+1000, expectedScore, cache, path.Copy(), global)
	traverse2(grid, pos, end, turnRight(dir), points+1000, expectedScore, cache, path, global)
}

func p2(grid Set[Point], start, end Point, expectedScore int) (sum int) {
	global := make(Set[Point])
	traverse2(grid, start, end, Right, 0, expectedScore, make(map[cacheKey]int), make(Set[Point]), global)
	return len(global)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	grid, start, end := parseInput(string(input))
	score := p1(grid, start, end)
	fmt.Printf("Part 1: %v\n", score)
	fmt.Printf("Part 2: %v\n", p2(grid, start, end, score))
}
