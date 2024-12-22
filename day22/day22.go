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

func parseInput(input string) (output []int) {
	for _, s := range strings.Split(input, "\n") {
		output = append(output, Atoi(s))
	}
	return
}

func mix(a, b int) int {
	return a ^ b
}

func prune(a int) int {
	return a % 16777216
}

func nextSecret(a int) int {
	a = prune(mix(a*64, a))
	a = prune(mix(a/32, a))
	a = prune(mix(a*2048, a))
	return a
}

func p1(inp []int) (sum int) {
	for _, a := range inp {
		for j := 0; j < 2000; j++ {
			a = nextSecret(a)
		}
		sum += a
	}
	return
}

func p2(inp []int) (priceMax int) {
	changes := make([]map[[4]int]int, len(inp))
	for i, a := range inp {
		changes[i] = make(map[[4]int]int)
		changesList := make([]int, 0)
		for j := 0; j < 2000; j++ {
			result := nextSecret(a)
			changesList = append(changesList, result%10-a%10)
			if j > 3 {
				key := [4]int{changesList[j-3], changesList[j-2], changesList[j-1], changesList[j]}
				if _, ok := changes[i][key]; !ok {
					changes[i][key] = result % 10
				}
			}
			a = result
		}
	}

	allChanges := make(map[[4]int]int)
	for _, ch := range changes {
		for key, val := range ch {
			allChanges[key] = allChanges[key] + val
		}
	}

	for _, val := range allChanges {
		if val > priceMax {
			priceMax = val
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
