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

type Point struct {
	x, y int
}

func parseInput(input string) (output []string) {
	return strings.Split(input, "\n")
}

var (
	numericKeypad = map[rune]Point{
		'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
		'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
		'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
		/*		  */ '0': {1, 3}, 'A': {2, 3},
	}

	directionalKeypad = map[rune]Point{
		/*		  */ '^': {1, 0}, 'A': {2, 0},
		'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
	}

	numericGap     = Point{0, 3}
	directionalGap = Point{0, 0}
)

func dirKeypadSeq(upDown, leftRight int) (seq1, seq2 []rune) {
	if upDown > 0 {
		for i := 0; i < upDown; i++ {
			seq2 = append(seq2, '^')
		}
	} else if upDown < 0 {
		for i := 0; i > upDown; i-- {
			seq2 = append(seq2, 'v')
		}
	}

	if leftRight > 0 {
		for i := 0; i < leftRight; i++ {
			seq1 = append(seq1, '<')
		}
	} else if leftRight < 0 {
		for i := 0; i > leftRight; i-- {
			seq1 = append(seq1, '>')
		}
	}
	return
}

type cacheKey struct {
	pos, newPos Point
	depth       int
}

func checkSeq(seq []rune, depth int, cache map[cacheKey]int) (result int) {
	dirPos1 := directionalKeypad['A']
	for _, s := range seq {
		newDirPos := directionalKeypad[s]
		result += check(dirPos1, newDirPos, directionalGap, depth-1, cache)
		dirPos1 = newDirPos
	}
	return
}

func check(pos, newPos, gap Point, depth int, cache map[cacheKey]int) int {
	ck := cacheKey{pos, newPos, depth}
	if v, ok := cache[ck]; ok {
		return v
	}

	upDown := pos.y - newPos.y
	leftRight := pos.x - newPos.x

	seq1, seq2 := dirKeypadSeq(upDown, leftRight)

	if depth == 0 {
		return len(seq1) + len(seq2) + 1
	}

	seq1valid := pos.y != gap.y || newPos.x != gap.x
	seq2valid := pos.x != gap.x || newPos.y != gap.y

	result := int(^uint(0) >> 1)

	if seq1valid {
		result = checkSeq(append(seq1, append(seq2, 'A')...), depth, cache)
	}

	if seq2valid {
		c := checkSeq(append(seq2, append(seq1, 'A')...), depth, cache)
		if c < result {
			result = c
		}
	}

	cache[ck] = result
	return result
}

func solve(inp []string, depth int) (sum int) {
	cache := make(map[cacheKey]int)

	for _, code := range inp {
		numPos := numericKeypad['A']
		result := 0
		for _, c := range code {
			newNumPos := numericKeypad[c]
			result += check(numPos, newNumPos, numericGap, depth, cache)
			numPos = newNumPos

		}
		sum += result * Atoi(code[:3])
	}

	return
}

func p1(inp []string) int {
	return solve(inp, 2)
}

func p2(inp []string) int {
	return solve(inp, 25)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
