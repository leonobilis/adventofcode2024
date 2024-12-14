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

type Robot struct {
	pos, vel Point
}

func (r Robot) posAfter(time int) Point {
	x := ((r.pos.x+time*r.vel.x)%101 + 101) % 101
	y := ((r.pos.y+time*r.vel.y)%103 + 103) % 103
	return Point{x, y}

}

func parseInput(input string) (output []Robot) {
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, " v=")
		p := strings.Split(s2[0], ",")
		v := strings.Split(s2[1], ",")
		output = append(output, Robot{Point{Atoi(p[0][2:]), Atoi(p[1])}, Point{Atoi(v[0]), Atoi(v[1])}})
	}
	return
}

func p1(input []Robot) (sum int) {
	counter1, counter2, counter3, counter4 := 0, 0, 0, 0
	for _, r := range input {
		pos := r.posAfter(100)
		if pos.x < 50 {
			if pos.y < 51 {
				counter1++
			} else if pos.y > 51 {
				counter2++

			}
		} else if pos.x > 50 {
			if pos.y < 51 {
				counter3++
			} else if pos.y > 51 {
				counter4++

			}
		}
	}
	return counter1 * counter2 * counter3 * counter4
}

func printGrid(grid Set[Point]) {
	for y := 0; y < 103; y++ {
		for x := 0; x < 101; x++ {
			if grid.Contains(Point{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func p2(input []Robot) (sum int) {
	for i := 1; true; i += 1 {
		grid := make(Set[Point])

		for _, r := range input {
			grid.Add(r.posAfter(i))
		}

		for y := 0; y < 93; y++ {
			for x := 0; x < 101; x++ {
				if grid.Contains(Point{x, y}) {
					j := 0
					for ; j < 10; j++ {
						if !grid.Contains(Point{x, y + j}) {
							break
						}
					}
					if j == 10 {
						printGrid(grid)
						return i
					}
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
