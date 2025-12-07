package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	b, _ := io.ReadAll(file)
	lines := strings.Split(string(b), "\n")

	beams := map[int]int{strings.Index(lines[0], "S"): 1}
	split := 0
	sum := 0

	for i := 1; i < len(lines); i++ {
		newBeams := map[int]int{}

		for beam, timelines := range beams {
			if lines[i][beam] == '^' {
				newBeams[beam-1] += timelines
				newBeams[beam+1] += timelines
				split++
			} else if lines[i][beam] == '.' {
				newBeams[beam] += timelines
			}
		}

		beams = newBeams
	}

	for _, timelines := range beams {
		sum += timelines
	}

	fmt.Println("Times the beam will be split (1):", split)
	fmt.Println("Timelines (2):", sum)
}
