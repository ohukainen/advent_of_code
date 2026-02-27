package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	grid := toGrid(lines)

	fmt.Println("Part 1: ", part1(grid))
	fmt.Println("Part 2: ", part2(grid))
}

func part1(grid [][]byte) int {
	_, answer := countAndRemoveMovable(grid)
	return answer
}

func part2(grid [][]byte) int {
	answer := 0
	current := grid

	for {
		next, count := countAndRemoveMovable(current)

		current = next
		answer += count
		if count == 0 {
			return answer
		}
	}
}

func countAndRemoveMovable(grid [][]byte) ([][]byte, int) {
	count := 0
	next := copyGrid(grid)

	for y, line := range grid {
		for x, char := range line {
			if char != '@' {
				continue
			}

			if isMovable(x, y, grid) {
				count++
				next[y][x] = '.'
			}
		}
	}

	return next, count
}

func isMovable(x int, y int, grid [][]byte) bool {
	nAdjacent := 0
	xMax := len(grid[0]) - 1
	yMax := len(grid) - 1

	for i := max(x-1, 0); i <= x+1 && i <= xMax; i++ {
		for j := max(y-1, 0); j <= y+1 && j <= yMax; j++ {
			if i == x && j == y {
				continue
			}

			if grid[j][i] == '@' {
				nAdjacent++
			}

			if nAdjacent > 3 {
				return false
			}
		}
	}
	return true
}

func toGrid(lines []string) [][]byte {
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return grid
}

func copyGrid(grid [][]byte) [][]byte {
	next := make([][]byte, len(grid))
	for i, row := range grid {
		next[i] = make([]byte, len(row))
		copy(next[i], row)
	}
	return next
}
