package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)	

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	line := (string(data))
	ranges := strings.Split(strings.TrimSpace(line), ",")

	fmt.Println("Part 1: ", part1(ranges))
	fmt.Println("Part 2: ", part2(ranges))
}

func part1(ranges []string) int {
	answer := 0

	for _, rangeStr := range ranges {
		first, last := parseRange(rangeStr)

		answer += addIfRepeats(first, last)
	}
	
	return answer
}

func addIfRepeats(first int, last int) int {
	sum := 0

	for i := first; i <= last; i++ {
		iStr := strconv.Itoa(i)
		if len(iStr) % 2 == 1 {
			continue
		}

		middle := len(iStr)/2
		if iStr[0:middle] == iStr[middle:] {
			sum += i 
		}
	}
	return sum
}

func part2(ranges []string) int {
	answer := 0

	for _, rangeStr := range ranges {
		first, last := parseRange(rangeStr)

		answer += addAnyRepeats(first, last)
	}
	
	return answer
}

func addAnyRepeats(first int, last int) int {
	sum := 0

	for i := first; i <= last; i++ {
		if repeats(i) {
			sum += i
		}
	}
	return sum
}

func repeats(id int) bool {
	idStr := strconv.Itoa(id)
	length := len(idStr)

	for segments := 2; segments <= length; segments++ {
		if length % segments != 0 {
			continue
		}
		segmentLen := length / segments
		
		for segment := range segments {
			start := segment * segmentLen
			middle := start + segmentLen
			end := middle + segmentLen

			if idStr[start:middle] != idStr[middle:end] {
				break 
			}

			if end == length {
				return true
			}
		}
	}
	return false 
}

func parseRange(rangeStr string) (int, int) {
	rangeArray := strings.Split(rangeStr, "-")
	first, err := strconv.Atoi(rangeArray[0])
	if err != nil {
		panic(err)
	}
	last, err := strconv.Atoi(rangeArray[1])
	if err != nil {
		panic(err)
	}

	return first, last
}
