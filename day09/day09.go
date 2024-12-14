package main

import (
	"fmt"
	"os"
)

func AsciiToInt(ascii byte) int {
	return int(ascii - '0')
}

func p1(inp string) int {
	result := make([]int, 0)

	reverseIndex := len(inp) - 1
	remaining := AsciiToInt(inp[reverseIndex])

	for i := 0; i < reverseIndex; i += 2 {
		val := AsciiToInt(inp[i])
		for j := 0; j < val; j++ {
			result = append(result, i/2)
		}
		if i >= reverseIndex {
			break
		}
		free := AsciiToInt(inp[i+1])
		for j := free; j > 0; {
			k := remaining
			if j < remaining {
				k = j
			}
			for ; k > 0; k-- {
				result = append(result, reverseIndex/2)
				j--
				remaining--
			}
			if j > 0 {
				reverseIndex -= 2
				remaining = AsciiToInt(inp[reverseIndex])
			}
		}
	}

	for i := 0; i < remaining; i++ {
		result = append(result, reverseIndex/2)
	}

	sum := 0
	for i, v := range result {
		sum += i * v
	}

	return sum
}

type Segment struct {
	index, size, free int
	files             []File
}

func (s *Segment) AddFile(f File) {
	s.files = append(s.files, f)
	s.free -= f.size
}

func (s *Segment) Erase() {
	s.files = make([]File, 0)
	s.free = s.size
}

func (s *Segment) Val() int {
	val := 0
	i := s.index
	for _, f := range s.files {
		val += f.Val(i)
		i += f.size
	}
	return val
}

type File struct {
	size int
	id   int
}

func (f *File) Val(start int) int {
	val := 0
	for i := 0; i < f.size; i++ {
		val += (start + i) * f.id
	}
	return val
}

func p2(inp string) int {
	disc := make([]Segment, 0)

	index := 0
	last := len(inp) - 1
	for i, index := 0, 0; i < last; i += 2 {
		fileSize, free := AsciiToInt(inp[i]), AsciiToInt(inp[i+1])
		disc = append(disc, Segment{index: index, size: fileSize, free: 0, files: []File{{id: i / 2, size: fileSize}}})
		disc = append(disc, Segment{index: index + fileSize, size: free, free: free, files: make([]File, 0)})
		index += fileSize + free
	}
	disc = append(disc, Segment{index: index, size: AsciiToInt(inp[last]), free: 0, files: []File{{id: last / 2, size: AsciiToInt(inp[last])}}})

	for i := len(disc) - 1; i > 0; i -= 2 {
		for j := 0; j < i; j++ {
			if disc[i].size <= disc[j].free {
				disc[j].AddFile(disc[i].files[0])
				disc[i].Erase()
				break
			}
		}
	}

	sum := 0
	for _, segment := range disc {
		sum += segment.Val()
	}

	return sum
}
func main() {
	input, _ := os.ReadFile("input.txt")
	inp := string(input)
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
