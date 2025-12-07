package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	var numbers [][]int

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines)-1; i++ {
		numbers = append(numbers, []int{})

		for s := range strings.SplitSeq(lines[i], " ") {
			if s != "" {
				n, _ := strconv.Atoi(s)
				numbers[i] = append(numbers[i], n)
			}
		}
	}

	i := 0
	total := 0

	for _, op := range lines[len(lines)-1] {
		if op == ' ' {
			continue
		}

		var m int

		if op == '+' {
			m = 0

			for j := 0; j < len(lines)-1; j++ {
				m += numbers[j][i]
			}

		} else if op == '*' {
			m = 1

			for j := 0; j < len(lines)-1; j++ {
				m *= numbers[j][i]
			}
		}

		total += m
		i++
	}

	fmt.Println("Grand total (1):", total)

	total = 0
	ns := []int{}

	for j := len(lines[0]) - 1; j >= 0; j-- {
		var b []byte

		for i := 0; i < len(lines)-1; i++ {
			b = append(b, lines[i][j])
		}

		n, _ := strconv.Atoi(strings.TrimSpace(string(b)))
		ns = append(ns, n)

		if lines[len(lines)-1][j] == ' ' {
			continue
		}

		var m int

		if lines[len(lines)-1][j] == '+' {
			m = 0

			for _, n := range ns {
				m += n
			}

		} else if lines[len(lines)-1][j] == '*' {
			m = 1

			for _, n := range ns {
				m *= n
			}
		}

		total += m
		ns = []int{}
		j--
	}

	fmt.Println("Grand total (2):", total)
}
