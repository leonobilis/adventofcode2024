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

func process(inp string, iter5 int) (length int) {
	process5 := func(s string) map[string]int {
		count := make(map[string]int)
		stones := []string{s}
		for i := 0; i < 5; i++ {
			newStones := make([]string, 0)
			for _, s := range stones {
				if s == "0" {
					newStones = append(newStones, "1")
				} else if len(s)%2 == 0 {
					newStones = append(newStones, s[:len(s)/2])
					newStones = append(newStones, strconv.Itoa(Atoi(s[len(s)/2:])))
				} else {
					newStones = append(newStones, strconv.Itoa(Atoi(s)*2024))
				}
				stones = newStones
			}
		}

		for _, s := range stones {
			count[s]++
		}
		return count
	}

	stones := make(map[string]int)
	cache := make(map[string]map[string]int)

	for _, s := range strings.Split(inp, " ") {
		stones[s]++
	}

	for i := 0; i < iter5; i++ {
		newStones := make(map[string]int)
		for stone, count := range stones {
			var stonesToAdd map[string]int
			var ok bool
			if stonesToAdd, ok = cache[stone]; !ok {
				stonesToAdd = process5(stone)
				cache[stone] = stonesToAdd

			}
			for stone2, count2 := range stonesToAdd {
				newStones[stone2] += count2 * count
			}
		}
		stones = newStones
	}

	for _, count := range stones {
		length += count
	}

	fmt.Printf("Cache size: %v\n", len(cache))

	return length
}

func p1(inp string) int {
	return process(inp, 5)
}

func p2(inp string) int {
	return process(inp, 15)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Printf("Part 1: %v\n", p1(string(input)))
	fmt.Printf("Part 2: %v\n", p2(string(input)))
}
