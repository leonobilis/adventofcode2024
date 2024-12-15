package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

const (
	Right = '>'
	Down  = 'v'
	Left  = '<'
	Up    = '^'
)

func copyMap[K comparable, V any](original map[K]V) map[K]V {
	copy := make(map[K]V)
	for key, value := range original {
		copy[key] = value
	}
	return copy
}

func parseInput(input string) (grid map[Position]rune, moves []rune, robot Position) {
	s := strings.Split(input, "\n\n")
	grid = make(map[Position]rune)
	for y, s2 := range strings.Split(s[0], "\n") {
		for x, s3 := range s2 {
			if s3 == '@' {
				robot = Position{x, y}
			} else if s3 != '.' {
				grid[Position{x, y}] = s3
			}

		}
	}
	for _, s2 := range strings.Split(s[1], "\n") {
		moves = append(moves, []rune(s2)...)
	}
	return
}

func p1(grid map[Position]rune, moves []rune, robot Position) (sum int) {
	grid = copyMap(grid)

	for _, m := range moves {
		switch m {
		case Right:
			p := Position{robot.x + 1, robot.y}
			for ; grid[p] != '#'; p.x++ {
				_, ok := grid[p]
				if !ok {
					if p != (Position{robot.x + 1, robot.y}) {
						delete(grid, Position{robot.x + 1, robot.y})
						grid[p] = 'O'
					}
					robot = Position{robot.x + 1, robot.y}
					break
				}

			}
		case Down:
			for p := (Position{robot.x, robot.y + 1}); grid[p] != '#'; p.y++ {
				_, ok := grid[p]
				if !ok {
					if p != (Position{robot.x, robot.y + 1}) {
						delete(grid, Position{robot.x, robot.y + 1})
						grid[p] = 'O'
					}
					robot = Position{robot.x, robot.y + 1}
					break
				}

			}
		case Left:
			for p := (Position{robot.x - 1, robot.y}); grid[p] != '#'; p.x-- {
				_, ok := grid[p]
				if !ok {
					if p != (Position{robot.x - 1, robot.y}) {
						delete(grid, Position{robot.x - 1, robot.y})
						grid[p] = 'O'
					}
					robot = Position{robot.x - 1, robot.y}
					break
				}

			}
		case Up:
			for p := (Position{robot.x, robot.y - 1}); grid[p] != '#'; p.y-- {
				_, ok := grid[p]
				if !ok {
					if p != (Position{robot.x, robot.y - 1}) {
						delete(grid, Position{robot.x, robot.y - 1})
						grid[p] = 'O'
					}
					robot = Position{robot.x, robot.y - 1}
					break
				}

			}
		}

	}
	for p, v := range grid {
		if v == 'O' {
			sum += p.x + 100*p.y
		}
	}

	return
}

func p2(grid map[Position]rune, moves []rune, robot Position) (sum int) {
	newGrid := make(map[Position]rune)
	for p, v := range grid {
		if v == '#' {
			newGrid[Position{p.x * 2, p.y}] = '#'
			newGrid[Position{p.x*2 + 1, p.y}] = '#'
		} else if v == 'O' {
			newGrid[Position{p.x * 2, p.y}] = '['
			newGrid[Position{p.x*2 + 1, p.y}] = ']'
		}
	}
	grid = newGrid
	robot.x *= 2

	for _, m := range moves {
		switch m {
		case Right:
			for p, boxes := (Position{robot.x + 1, robot.y}), make([]Position, 0); grid[p] != '#'; p.x += 2 {
				g, ok := grid[p]
				if ok {
					if g == '[' {
						boxes = append(boxes, p)
					}
				} else {
					for i := len(boxes) - 1; i >= 0; i-- {
						b := boxes[i]
						delete(grid, b)
						delete(grid, Position{b.x + 1, b.y})
						grid[Position{b.x + 1, b.y}] = '['
						grid[Position{b.x + 2, b.y}] = ']'
					}
					robot = Position{robot.x + 1, robot.y}
					break
				}

			}
		case Left:
			for p, boxes := (Position{robot.x - 1, robot.y}), make([]Position, 0); grid[p] != '#'; p.x -= 2 {
				g, ok := grid[p]
				if ok {
					if g == ']' {
						boxes = append(boxes, p)
					}
				} else {
					for i := len(boxes) - 1; i >= 0; i-- {
						b := boxes[i]
						delete(grid, b)
						delete(grid, Position{b.x - 1, b.y})
						grid[Position{b.x - 1, b.y}] = ']'
						grid[Position{b.x - 2, b.y}] = '['
					}
					robot = Position{robot.x - 1, robot.y}
					break
				}
			}
		case Down:
			boxes := make([][]Position, 0)
			boxesUpdate := make([]Position, 0)

			p := Position{robot.x, robot.y + 1}
			g, ok := grid[p]
			if ok {
				if g == '[' {
					boxesUpdate = append(boxesUpdate, p)
				} else if g == ']' {
					boxesUpdate = append(boxesUpdate, Position{p.x - 1, p.y})
				} else {
					continue
				}
				boxes = append(boxes, boxesUpdate)
			} else {
				robot = p
				continue
			}

		downLoop:
			for {
				newBoxesUpdate := make([]Position, 0)
				for _, b := range boxesUpdate {
					p1, p2 := Position{b.x, b.y + 1}, Position{b.x + 1, b.y + 1}
					g1, g2 := grid[p1], grid[p2]
					if g1 == '#' || g2 == '#' {
						break downLoop
					}
					if g1 == '[' {
						newBoxesUpdate = append(newBoxesUpdate, p1)
					}
					if g1 == ']' {
						newBoxesUpdate = append(newBoxesUpdate, Position{p1.x - 1, p1.y})
					}
					if g2 == '[' {
						newBoxesUpdate = append(newBoxesUpdate, p2)
					}
				}

				if len(newBoxesUpdate) == 0 {
					robot = Position{robot.x, robot.y + 1}
					for i := len(boxes) - 1; i >= 0; i-- {
						b := boxes[i]
						for _, b2 := range b {
							delete(grid, b2)
							delete(grid, Position{b2.x + 1, b2.y})
							grid[Position{b2.x, b2.y + 1}] = '['
							grid[Position{b2.x + 1, b2.y + 1}] = ']'
						}
					}
					break downLoop
				} else {
					boxes = append(boxes, newBoxesUpdate)
					boxesUpdate = newBoxesUpdate
				}
			}
		case Up:
			boxes := make([][]Position, 0)
			boxesUpdate := make([]Position, 0)

			p := Position{robot.x, robot.y - 1}
			g, ok := grid[p]
			if ok {
				if g == '[' {
					boxesUpdate = append(boxesUpdate, p)
				} else if g == ']' {
					boxesUpdate = append(boxesUpdate, Position{p.x - 1, p.y})
				} else {
					continue
				}
				boxes = append(boxes, boxesUpdate)
			} else {
				robot = p
				continue
			}

		upLoop:
			for {
				newBoxesUpdate := make([]Position, 0)
				for _, b := range boxesUpdate {
					p1, p2 := Position{b.x, b.y - 1}, Position{b.x + 1, b.y - 1}
					g1, g2 := grid[p1], grid[p2]
					if g1 == '#' || g2 == '#' {
						break upLoop
					}
					if g1 == '[' {
						newBoxesUpdate = append(newBoxesUpdate, p1)
					}
					if g1 == ']' {
						newBoxesUpdate = append(newBoxesUpdate, Position{p1.x - 1, p1.y})
					}
					if g2 == '[' {
						newBoxesUpdate = append(newBoxesUpdate, p2)
					}
				}

				if len(newBoxesUpdate) == 0 {
					robot = Position{robot.x, robot.y - 1}
					for i := len(boxes) - 1; i >= 0; i-- {
						b := boxes[i]
						for _, b2 := range b {
							delete(grid, b2)
							delete(grid, Position{b2.x + 1, b2.y})
							grid[Position{b2.x, b2.y - 1}] = '['
							grid[Position{b2.x + 1, b2.y - 1}] = ']'
						}
					}
					break upLoop
				} else {
					boxes = append(boxes, newBoxesUpdate)
					boxesUpdate = newBoxesUpdate
				}
			}

		}
	}

	for p, v := range grid {
		if v == '[' {
			sum += p.x + 100*p.y
		}
	}

	return
}

func main() {
	input, _ := os.ReadFile("input.txt")
	grid, moves, robot := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(grid, moves, robot))
	fmt.Printf("Part 2: %v\n", p2(grid, moves, robot))
}
