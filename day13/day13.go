package main

import (
	"fmt"
	"os"
	"regexp"
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

type Machine struct {
	buttonA, buttonB Point
	prize            Point
}

func parseInput(input string) []Machine {
	re := regexp.MustCompile(`[\d]+`)
	entries := strings.Split(input, "\n\n")
	machines := make([]Machine, 0, len(entries))
	for _, s := range entries {
		s2 := strings.Split(s, "\n")
		ints := re.FindAllString(s2[0], 2)
		buttonA := Point{Atoi(ints[0]), Atoi(ints[1])}
		ints = re.FindAllString(s2[1], 2)
		buttonB := Point{Atoi(ints[0]), Atoi(ints[1])}
		ints = re.FindAllString(s2[2], 2)
		prize := Point{Atoi(ints[0]), Atoi(ints[1])}
		machines = append(machines, Machine{buttonA, buttonB, prize})
	}

	return machines
}

func (m Machine) Play() int {
	a := (m.prize.x - m.prize.y/m.buttonB.y*m.buttonB.x) / (m.buttonA.x - m.buttonB.x*m.buttonB.x/m.prize.y)
	b := (m.prize.y*m.buttonA.x - m.prize.x*m.buttonA.y) / (m.buttonA.x*m.buttonB.y - m.buttonA.y*m.buttonB.x)

	if m.buttonA.x*a+m.buttonB.x*b == m.prize.x && m.buttonA.y*a+m.buttonB.y*b == m.prize.y {
		return 3*a + b
	}
	return 0
}

func p1(inp []Machine) (sum int) {
	for _, m := range inp {
		sum += m.Play()
	}
	return
}

func p2(inp []Machine) (sum int) {
	for _, m := range inp {
		m.prize = Point{m.prize.x + 10000000000000, m.prize.y + 10000000000000}
		sum += m.Play()
	}
	return
}

func main() {
	input, _ := os.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
