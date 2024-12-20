package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yourbasic/graph"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	x, y int
}

func parseInput(input string) (output []string, start, end int) {
	output = strings.Split(input, "\n")
	lenX := len(output[0])

	for y, s := range output {
		for x, c := range s {
			if c == 'S' {
				start = y*lenX + x
				output[y] = output[y][:x] + "." + output[y][x+1:]
			} else if c == 'E' {
				end = y*lenX + x
				output[y] = output[y][:x] + "." + output[y][x+1:]
			}
			if start > 0 && end > 0 {
				break
			}
		}
	}

	return
}

func getGraph(input []string, lenX, lenY int) *graph.Mutable {
	g := graph.New(lenX * lenY)
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if input[y][x] == '.' {
				if x < lenX-1 && input[y][x+1] == '.' {
					g.AddCost(y*lenX+x, y*lenX+x+1, 1)
					g.AddCost(y*lenX+x+1, y*lenX+x, 1)
				}
				if y < lenY-1 && input[y+1][x] == '.' {
					g.AddCost(y*lenX+x, (y+1)*lenX+x, 1)
					g.AddCost((y+1)*lenX+x, y*lenX+x, 1)
				}
			}
		}
	}
	return g
}

const condition = 100

func check(input []string, start, end int, cheats int) (sum int) {
	lenX, lenY := len(input[0]), len(input)
	g := getGraph(input, lenX, lenY)

	points := make(map[Point]int)

	path, _ := graph.ShortestPath(g, start, end)
	for i, point := range path {
		points[Point{point % lenX, point / lenX}] = i
	}

	for p1, d1 := range points {
		delete(points, p1)
		for p2, d2 := range points {
			mDist := Abs(p1.x-p2.x) + Abs(p1.y-p2.y)
			if mDist >= 2 && mDist <= cheats && Abs(d1-d2)-mDist >= condition {
				sum++
			}
		}
	}

	return
}

func p1(input []string, start, end int) (sum int) {
	return check(input, start, end, 2)
}

func p2(input []string, start, end int) (sum int) {
	return check(input, start, end, 20)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	inp, start, end := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp, start, end))
	fmt.Printf("Part 2: %v\n", p2(inp, start, end))
}
