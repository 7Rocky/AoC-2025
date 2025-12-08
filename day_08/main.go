package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const NUM_PAIRS = 1000

type box struct {
	x, y, z int
}

type circuit []box

type junction struct {
	distance int
	boxes    [2]box
}

func (b *box) distance(a box) int {
	return int(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}

func (b *box) index(circuits []circuit) int {
	for i, circuit := range circuits {
		if slices.Contains(circuit, *b) {
			return i
		}
	}

	return -1
}

func (j *junction) index(circuits []circuit) [2]int {
	return [2]int{j.boxes[0].index(circuits), j.boxes[1].index(circuits)}
}

func product(circuits []circuit) int {
	result := 1

	for _, circuit := range circuits {
		result *= len(circuit)
	}

	return result
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var boxes []box

	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		z, _ := strconv.Atoi(coordinates[2])
		boxes = append(boxes, box{x, y, z})
	}

	var junctions []junction

	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			junctions = append(junctions, junction{distance: boxes[i].distance(boxes[j]), boxes: [2]box{boxes[i], boxes[j]}})
		}
	}

	slices.SortFunc(junctions, func(a, b junction) int {
		return a.distance - b.distance
	})

	var circuits []circuit

	for _, box := range boxes {
		circuits = append(circuits, circuit{box})
	}

	for i, junction := range junctions {
		if i == NUM_PAIRS {
			slices.SortFunc(circuits, func(a, b circuit) int {
				return len(b) - len(a)
			})

			fmt.Println("Three largest circuit sizes product (1):", product(circuits[:3]))
		}

		if indices := junction.index(circuits); indices == [2]int{-1, -1} {
			circuits = append(circuits, junction.boxes[:])
		} else if indices[0] == -1 {
			circuits[indices[1]] = append(circuits[indices[1]], junction.boxes[0])
		} else if indices[1] == -1 {
			circuits[indices[0]] = append(circuits[indices[0]], junction.boxes[1])
		} else if indices[0] != indices[1] {
			circuits[indices[0]] = append(circuits[indices[0]], circuits[indices[1]]...)
			circuits = slices.Delete(circuits, indices[1], indices[1]+1)
		}

		if len(circuits) == 1 {
			fmt.Println("Last two junction boxes x-coordinate product (2):", junction.boxes[0].x*junction.boxes[1].x)
			break
		}
	}
}
