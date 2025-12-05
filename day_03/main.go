package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func findMax(s string) (int, int) {
	index := -1
	max := -1

	for i, c := range s {
		if n := int(c - '0'); max < n {
			max = n
			index = i
		}
	}

	return index, max
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total1 := 0
	total2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		maxPower := 0

		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				maxPower = int(math.Max(float64(10*int(line[i]-'0')+int(line[j]-'0')), float64(maxPower)))
			}
		}

		total1 += maxPower

		start := 0
		maxPower = 0

		for i := range 12 {
			index, n := findMax(line[start : len(line)-11+i])
			start += index + 1
			maxPower = 10*maxPower + n
		}

		total2 += maxPower
	}

	fmt.Println("Total output joltage (1):", total1)
	fmt.Println("Total output joltage (2):", total2)
}
