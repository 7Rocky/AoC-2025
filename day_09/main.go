package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type tile struct {
	x, y int
}

type rectangle struct {
	a, b tile
	area int
}

func newRectangle(a, b tile) rectangle {
	return rectangle{a, b, (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func isInside(r rectangle, edges []tile) bool {
	for _, e := range edges {
		if min(r.a.x, r.b.x) < e.x && e.x < max(r.a.x, r.b.x) && min(r.a.y, r.b.y) < e.y && e.y < max(r.a.y, r.b.y) {
			return false
		}
	}

	return true
}

func main() {
	b, _ := os.ReadFile("input.txt")

	var tiles []tile
	var rectangles []rectangle
	var edges []tile

	for line := range strings.SplitSeq(string(b), "\n") {
		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		tiles = append(tiles, tile{x, y})
	}

	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			rectangles = append(rectangles, newRectangle(tiles[i], tiles[j]))
		}
	}

	slices.SortFunc(rectangles, func(a, b rectangle) int {
		return b.area - a.area
	})

	fmt.Println("Largest area (1):", rectangles[0].area)

	for i := 1; i <= len(tiles); i++ {
		if tiles[i-1].x == tiles[i%len(tiles)].x {
			for y := min(tiles[i-1].y, tiles[i%len(tiles)].y); y < max(tiles[i-1].y, tiles[i%len(tiles)].y); y++ {
				edges = append(edges, tile{tiles[i-1].x, y})
			}
		} else if tiles[i-1].y == tiles[i%len(tiles)].y {
			for x := min(tiles[i-1].x, tiles[i%len(tiles)].x); x < max(tiles[i-1].x, tiles[i%len(tiles)].x); x++ {
				edges = append(edges, tile{x, tiles[i-1].y})
			}
		}
	}

	for _, r := range rectangles {
		if isInside(r, edges) {
			fmt.Println("Largest area with red and green tiles (2):", r.area)
			break
		}
	}
}
