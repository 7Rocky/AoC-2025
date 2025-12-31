package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input.txt")
	index := -1
	pieces := map[int]int{}
	total := 0

	for line := range strings.SplitSeq(string(b), "\n") {
		if strings.Contains(line, "#") {
			pieces[index] += strings.Count(line, "#")
		} else if strings.Contains(line, "x") {
			x := strings.Index(line, "x")
			colon := strings.Index(line, ":")

			height, _ := strconv.Atoi(line[:x])
			width, _ := strconv.Atoi(line[x+1 : colon])

			area := 0

			for i, countString := range strings.Split(line[colon+2:], " ") {
				count, _ := strconv.Atoi(countString)
				area += count * pieces[i]
			}

			if area <= height*width {
				total += 1
			}
		} else if strings.Contains(line, ":") {
			index++
		}
	}

	fmt.Println("Possible regions (1):", total)
}
