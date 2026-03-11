package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) rectArea(other Coordinate) int {
	abs := func(n int) int {
		if n >= 0 {
			return n
		}
		return -n
	}
	return (abs(c.x-other.x) + 1) * (abs(c.y-other.y) + 1)
}

type Line struct {
	y    int
	xMin int
	xMax int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	coords := getCoordinates(lines)

	fmt.Println("Part 1: ", part1(coords))
	fmt.Println("Part 2: ", part2(coords))
}

func part1(coords []Coordinate) int {
	maxArea := 0
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			area := c1.rectArea(c2)
			if maxArea < area {
				maxArea = area
			}

		}
	}
	return maxArea
}

func part2(coords []Coordinate) int {
	lines := buildSortedLines(coords)
	insideIntervalsMap := getInsideIntervalsMap(lines)

	maxArea := 0
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			area := c1.rectArea(c2)
			if maxArea < area && isValid(c1, c2, insideIntervalsMap) {
				maxArea = area
			}

		}
	}
	return maxArea
}

func buildSortedLines(coords []Coordinate) []Line {
	crds := slices.Clone(coords)
	slices.SortFunc(crds, func(c1 Coordinate, c2 Coordinate) int {
		return c1.y - c2.y
	})

	var lines []Line

	for i := 0; i < len(crds)-1; i = i + 2 {
		c1 := crds[i]
		c2 := crds[i+1]

		if c1.y != c2.y {
			panic(fmt.Sprintf("Coordinates not on the same row, got: c1.y: %d and c2.y %d", c1.y, c2.y))
		}

		line := Line{c1.y, min(c1.x, c2.x), max(c1.x, c2.x)}
		lines = append(lines, line)
	}

	return lines
}

func getInsideIntervalsMap(lines []Line) map[int][]Line {
	currentY := lines[0].y
	insideIntervalsMap := map[int][]Line{}
	insideIntervals := []Line{lines[0]}
	insideIntervalsMap[currentY] = insideIntervals

	for _, line := range lines[1:] {
		for y := currentY + 1; y < line.y; y++ {
			insideIntervalsMap[y] = insideIntervals
		}

		newInsideIntervals, extends := getNewIntervals(line, insideIntervals)

		if extends {
			insideIntervalsMap[line.y] = newInsideIntervals
		} else {
			insideIntervalsMap[line.y] = insideIntervals
		}
		insideIntervals = newInsideIntervals
		currentY = line.y
	}
	return insideIntervalsMap
}

func getNewIntervals(newLine Line, intervals []Line) ([]Line, bool) {
	slices.SortFunc(intervals, func(l1 Line, l2 Line) int {
		return l1.xMin - l2.xMin
	})

	var newIntervals []Line
	extends := false
	matched := false

	for _, l := range intervals {
		if newLine.xMin == l.xMin && newLine.xMax == l.xMax {
			// Closes an interval
			matched = true
		} else if l.xMin == newLine.xMin {
			newIntervals = append(newIntervals, Line{l.y, newLine.xMax, l.xMax})
			matched = true
		} else if l.xMax == newLine.xMax {
			newIntervals = append(newIntervals, Line{l.y, l.xMin, newLine.xMin})
			matched = true
		} else if l.xMin == newLine.xMax {
			newIntervals = append(newIntervals, Line{l.y, newLine.xMin, l.xMax})
			extends = true
			matched = true
		} else if l.xMax == newLine.xMin {
			newIntervals = append(newIntervals, Line{l.y, l.xMin, newLine.xMax})
			extends = true
			matched = true
		} else {
			newIntervals = append(newIntervals, l)
		}
	}

	if !matched {
		newIntervals = append(newIntervals, newLine)
		extends = true
	}

	newIntervals = mergeAdjacent(newIntervals)
	return newIntervals, extends
}

func mergeAdjacent(intervals []Line) []Line {
	if len(intervals) == 0 {
		return intervals
	}
	slices.SortFunc(intervals, func(l1 Line, l2 Line) int {
		return l1.xMin - l2.xMin
	})
	newIntervals := []Line{intervals[0]}
	for _, l := range intervals[1:] {
		last := &newIntervals[len(newIntervals)-1]
		if l.xMin <= last.xMax {
			last.xMax = max(last.xMax, l.xMax)
		} else {
			newIntervals = append(newIntervals, l)
		}
	}
	return newIntervals
}

func isValid(c1 Coordinate, c2 Coordinate, insideIntervalsMap map[int][]Line) bool {
	xMin := min(c1.x, c2.x)
	xMax := max(c1.x, c2.x)
	yMin := min(c1.y, c2.y)
	yMax := max(c1.y, c2.y)

	for y := yMin; y <= yMax; y++ {
		if !isInside(xMin, xMax, insideIntervalsMap[y]) {
			return false
		}
	}
	return true
}

func isInside(xMin int, xMax int, lines []Line) bool {
	for _, line := range lines {
		if line.xMin <= xMin && xMax <= line.xMax {
			return true
		}
	}
	return false
}

func getCoordinates(lines []string) []Coordinate {
	coords := []Coordinate{}
	for _, line := range lines {
		coordStr := strings.Split(line, ",")
		if len(coordStr) != 2 {
			panic(fmt.Sprintf("Expected 2 coordinates, got: %d", len(coordStr)))
		}

		x, err := strconv.Atoi(coordStr[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(coordStr[1])
		if err != nil {
			panic(err)
		}
		coords = append(coords, Coordinate{x, y})
	}
	return coords
}
