package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	tachyonManifold := toTachyonManifold(lines)

	fmt.Println("Part 1: ", part1(tachyonManifold))
	fmt.Println("Part 2: ", part2(tachyonManifold))
}

func part1(tachyonManifold [][]byte) uint64 {
	var ans uint64
	beaming := make([]bool, len(tachyonManifold[0]))
	indexBeam := slices.Index(tachyonManifold[0], 'S')
	beaming[indexBeam] = true

	for _, row := range tachyonManifold[1:] {
		for i, b := range row {
			if b == '^' && beaming[i] {
				ans++
				updateBeaming(row, beaming, i)
			}
		}

	}
	return ans
}

func updateBeaming(row []byte, beaming []bool, index int) {
	beaming[index] = false
	for i := index - 1; i >= 0; i-- {
		if row[i] != '^' {
			beaming[i] = true
			break
		}
	}

	for i := index + 1; i < len(row); i++ {
		if row[i] != '^' {
			beaming[i] = true
			break
		}
	}
}

func part2(tachyonManifold [][]byte) uint64 {
	totalBeams := make([]uint64, len(tachyonManifold[0]))
	indexBeam := slices.Index(tachyonManifold[0], 'S')
	totalBeams[indexBeam] = 1

	for _, row := range tachyonManifold[1:] {
		for i, b := range row {
			if b == '^' && totalBeams[i] > 0 {
				updateTotalBeams(row, totalBeams, i)
			}
		}

	}
	var ans uint64
	for _, n := range totalBeams {
		ans += n
	}
	return ans
}

func updateTotalBeams(row []byte, beaming []uint64, index int) {
	for i := index - 1; i >= 0; i-- {
		if row[i] != '^' {
			beaming[i] += beaming[index]
			break
		}
	}

	for i := index + 1; i < len(row); i++ {
		if row[i] != '^' {
			beaming[i] += beaming[index]
			break
		}
	}

	beaming[index] = 0
}

func toTachyonManifold(lines []string) [][]byte {
	tachyonManifold := make([][]byte, len(lines))

	for i := range tachyonManifold {
		tachyonManifold[i] = []byte(lines[i])
	}

	return tachyonManifold
}
