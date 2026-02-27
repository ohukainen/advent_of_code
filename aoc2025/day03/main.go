package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}

func part1(lines []string) int {
	answer := 0

	for _, line := range lines {
		first, index := findHighest(line[:len(line)-1])
		second, _ := findHighest(line[index+1:])

		numStr := first + second
		n, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		answer += n
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	for _, line := range lines {
		var sb strings.Builder
		subline := line
		found := 0
		for {
			next, index := findNext(subline, found)

			sb.WriteString(next)
			found += 1
			subline = subline[index+1:]

			if found == 12 {
				break
			}
		}

		n, err := strconv.Atoi(sb.String())
		if err != nil {
			panic(err)
		}
		answer += n
	}
	return answer
}

func findNext(line string, found int) (string, int) {
	invalid := 11 - found
	return findHighest(line[:len(line)-invalid])
}

func findHighest(line string) (string, int) {
	var ret string
	var index int

	for _, joltage := range []string{"9", "8", "7", "6", "5", "4", "3", "2", "1"} {
		if !strings.Contains(line, joltage) {
			continue
		}
		index = strings.Index(line, joltage)
		ret = joltage
		break
	}
	return ret, index
}
