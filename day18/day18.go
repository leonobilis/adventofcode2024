package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yourbasic/graph"
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

const GridLen = 71

func parseInput(input string) (output []Point) {
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, ",")
		output = append(output, Point{Atoi(s2[0]), Atoi(s2[1])})
	}
	return
}

func GetGraph(input []Point) *graph.Mutable {
	corrupted := make(Set[Point])
	for i := 0; i < 1024; i++ {
		corrupted.Add(input[i])
	}

	g := graph.New(GridLen * GridLen)
	for y := 0; y < GridLen; y++ {
		for x := 0; x < GridLen; x++ {
			if !corrupted.Contains(Point{x, y}) {
				if (x < GridLen-1) && !corrupted.Contains(Point{x + 1, y}) {
					g.AddCost(y*GridLen+x, y*GridLen+x+1, 1)
					g.AddCost(y*GridLen+x+1, y*GridLen+x, 1)
				}
				if (y < GridLen-1) && !corrupted.Contains(Point{x, y + 1}) {
					g.AddCost(y*GridLen+x, (y+1)*GridLen+x, 1)
					g.AddCost((y+1)*GridLen+x, y*GridLen+x, 1)
				}
			}
		}
	}
	return g
}

func p1(input []Point) int64 {
	_, dist := graph.ShortestPath(GetGraph(input), 0, GridLen*GridLen-1)
	return dist
}

func p2(input []Point) string {
	g := GetGraph(input)

	for i := 1024; i < len(input); i++ {
		p := input[i]
		if p.x < GridLen-1 {
			g.Delete(p.y*GridLen+p.x, p.y*GridLen+p.x+1)
			g.Delete(p.y*GridLen+p.x+1, p.y*GridLen+p.x)
		}
		if p.x > 0 {
			g.Delete(p.y*GridLen+p.x, p.y*GridLen+p.x-1)
			g.Delete(p.y*GridLen+p.x-1, p.y*GridLen+p.x)
		}
		if p.y < GridLen-1 {
			g.Delete((p.y+1)*GridLen+p.x, p.y*GridLen+p.x)
			g.Delete(p.y*GridLen+p.x, (p.y+1)*GridLen+p.x)

		}
		if p.y > 0 {
			g.Delete(p.y*GridLen+p.x, (p.y-1)*GridLen+p.x)
			g.Delete((p.y-1)*GridLen+p.x, p.y*GridLen+p.x)
		}
		_, dist := graph.ShortestPath(g, 0, GridLen*GridLen-1)
		if dist == -1 {
			return fmt.Sprintf("%d,%d\n", p.x, p.y)
		}

	}
	return "error"
}

func main() {
	input, _ := os.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
