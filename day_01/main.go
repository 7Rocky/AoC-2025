package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rotations []int

	for scanner.Scan() {
		line := scanner.Text()

		if n, _ := strconv.Atoi(line[1:]); line[0] == 'L' {
			rotations = append(rotations, -n)
		} else if line[0] == 'R' {
			rotations = append(rotations, n)
		}

	}

	curr := 50
	prev := 0
	password1 := 0
	password2 := 0

	for _, r := range rotations {
		prev = curr

		for curr = (curr + r) % 100; curr < 0; curr += 100 {
		}

		if curr == 0 {
			password1++
			password2++
		} else if 0 < prev && ((curr <= prev && 0 < r) || (prev <= curr && r < 0)) {
			password2++
		}

		password2 += int(math.Abs(float64(r))) / 100
	}

	fmt.Println("Password (1):", password1)
	fmt.Println("Password (2):", password2)
}
