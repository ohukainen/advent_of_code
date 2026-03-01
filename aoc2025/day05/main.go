package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start uint64
	end   uint64
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ranges, ids := processData(data)

	fmt.Println("Part 1: ", part1(ranges, ids))
	fmt.Println("Part 2: ", part2(ranges))
}

func part1(ranges []Range, ids []uint64) int {
	ans := 0
	for _, id := range ids {

		if isFresh(id, ranges) {
			ans++
		}
	}
	return ans
}

func isFresh(id uint64, ranges []Range) bool {
	i := sort.Search(len(ranges), func(i int) bool { return ranges[i].start > id })

	return i != 0 && id <= ranges[i-1].end
}

func part2(ranges []Range) uint64 {
	var ans uint64
	for _, r := range ranges {
		ans += r.end - r.start + 1
	}
	return ans
}

func processData(data []byte) ([]Range, []uint64) {
	var ranges []Range
	var ids []uint64
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	isRanges := true

	for _, line := range lines {
		if len(line) == 0 {
			isRanges = false
			continue
		}

		if isRanges {
			ranges = append(ranges, parseRange(line))
		} else {
			id, err := strconv.ParseUint(line, 10, 64)
			if err != nil {
				panic(err)
			}
			ids = append(ids, id)
		}
	}

	ranges = sortAndMergeRanges(ranges)

	return ranges, ids
}

func parseRange(rangeStr string) Range {
	splitRange := strings.Split(strings.TrimSpace(rangeStr), "-")

	if len(splitRange) != 2 {
		panic("Range not parsed correctly")
	}

	start, err := strconv.ParseUint(splitRange[0], 10, 64)
	if err != nil {
		panic(err)
	}

	end, err := strconv.ParseUint(splitRange[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return Range{start, end}
}

func sortAndMergeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]
		if r.start <= last.end {
			last.end = max(r.end, last.end)
		} else {
			merged = append(merged, r)
		}
	}
	return merged
}
