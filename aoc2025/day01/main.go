package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
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
	number := 50

	for _, line := range lines {
		dir, n := parseLine(line)

		number = rotate(number, dir, n)

		if number == 0 {
			answer++
		}
	}

	return answer 
}

func rotate(number int, dir byte, n int) int {
	if dir == 'R' {
		return (number + n) % 100
	} else if ((number - n) % 100) < 0 {
		return ((number - n) % 100) + 100
	} else {
		return ((number - n) % 100)
	}
}

func part2(lines []string) int {
	answer := 0
	number := 50

	for _, line := range lines {
		dir, n := parseLine(line)

		for range n {
			var click bool
			number, click = rotateOnce(number, dir)
			if click {
				answer++
			}
		}
	}

	return answer 
}

func rotateOnce(number int, dir byte) (int, bool) {
	if dir == 'R' {
		if number == 99 {
			return 0, true
		}
		return number + 1, false
	} else {
		if number == 0 {
			return 99, false
		}
		if number == 1 {
			return number - 1, true
		}
		return number - 1, false
	}
}

func parseLine(line string) (byte, int) {
	dir, nStr := line[0], line[1:]

	n, err := strconv.Atoi(nStr) 
	if err != nil {
		panic(err)
	}

	return dir, n
}
