package main

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int64
}

type Connection struct {
	distance int64
	p1, p2   Point
}

type UnionFind struct {
	parent map[Point]Point
	rank   map[Point]int
}

func NewUnionFind(points []Point) *UnionFind {
	uf := &UnionFind{
		parent: make(map[Point]Point),
		rank:   make(map[Point]int),
	}
	for _, point := range points {
		uf.parent[point] = point
	}
	return uf
}

func (uf *UnionFind) Find(p Point) Point {
	if uf.parent[p] != p {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UnionFind) Union(connection Connection) bool {
	rootP1 := uf.Find(connection.p1)
	rootP2 := uf.Find(connection.p2)
	if rootP1 == rootP2 {
		return false
	}

	if uf.rank[rootP1] > uf.rank[rootP2] {
		uf.parent[rootP2] = rootP1
	} else if uf.rank[rootP2] > uf.rank[rootP1] {
		uf.parent[rootP1] = rootP2
	} else {
		uf.parent[rootP2] = rootP1
		uf.rank[rootP1]++
	}
	return true
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	points := parsePoints(lines)
	connections := getSortedConnections(points)

	fmt.Println("Part 1: ", part1(points, connections))
	fmt.Println("Part 2: ", part2(points, connections))
}

func part1(points []Point, connections []Connection) int {
	uf := NewUnionFind(points)

	for i := range 1000 {
		uf.Union(connections[i])
	}

	sizes := getSortedSizes(uf)
	return sizes[0] * sizes[1] * sizes[2]
}

func getSortedSizes(uf *UnionFind) []int {
	sizeMap := make(map[Point]int)
	for point := range maps.Keys(uf.parent) {
		root := uf.Find(point)
		sizeMap[root]++
	}
	sizes := slices.Collect(maps.Values(sizeMap))
	slices.SortFunc(sizes, func(lhs int, rhs int) int { return rhs - lhs })
	return sizes
}

func part2(points []Point, connections []Connection) int64 {
	var lastConnection Connection
	uf := NewUnionFind(points)

	for _, connection := range connections {
		if uf.Union(connection) {
			lastConnection = connection
		}
	}

	return lastConnection.p1.x * lastConnection.p2.x
}

func cartesianDistanceSq(p1 Point, p2 Point) int64 {
	deltaXSquare := (p1.x - p2.x) * (p1.x - p2.x)
	deltaYSquare := (p1.y - p2.y) * (p1.y - p2.y)
	deltaZSquare := (p1.z - p2.z) * (p1.z - p2.z)
	return deltaXSquare + deltaYSquare + deltaZSquare
}

func parsePoints(lines []string) []Point {
	var points []Point
	for _, line := range lines {
		pointsStr := strings.Split(line, ",")
		if len(pointsStr) != 3 {
			panic(fmt.Sprintf("Invalid input: expected 3 coordinates, got: %d", len(pointsStr)))
		}

		x, err := strconv.ParseInt(pointsStr[0], 10, 64)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseInt(pointsStr[1], 10, 64)
		if err != nil {
			panic(err)
		}
		z, err := strconv.ParseInt(pointsStr[2], 10, 64)
		if err != nil {
			panic(err)
		}
		points = append(points, Point{x, y, z})
	}
	return points
}

func getSortedConnections(points []Point) []Connection {
	connections := []Connection{}
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			deltaSq := cartesianDistanceSq(p1, p2)
			connection := Connection{deltaSq, p1, p2}
			connections = append(connections, connection)
		}
	}

	slices.SortFunc(connections, func(c1 Connection, c2 Connection) int {
		return cmp.Compare(c1.distance, c2.distance)
	})
	return connections
}
