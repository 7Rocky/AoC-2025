package main

import (
	"fmt"
	"os"
	"strings"
)

var devices = map[string][]string{}

func countWays(start string, end string) (ways int) {
	stack := []string{start}

	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]

		if curr == end {
			ways++
		} else {
			stack = append(stack, devices[curr]...)
		}
	}

	return
}

func countWaysRecursive(start, end string, ways map[string]int) int {
	if count, ok := ways[start]; ok {
		return count
	}

	if start == end {
		return 1
	}

	for _, next := range devices[start] {
		ways[start] += countWaysRecursive(next, end, ways)
	}

	return ways[start]
}

func main() {
	b, _ := os.ReadFile("input.txt")

	for line := range strings.SplitSeq(string(b), "\n") {
		colon := strings.Index(line, ": ")
		devices[line[:colon]] = strings.Split(line[colon+1:], " ")
	}

	ways := countWaysRecursive("svr", "dac", map[string]int{}) *
		countWaysRecursive("dac", "fft", map[string]int{}) *
		countWaysRecursive("fft", "out", map[string]int{})
	ways += countWaysRecursive("svr", "fft", map[string]int{}) *
		countWaysRecursive("fft", "dac", map[string]int{}) *
		countWaysRecursive("dac", "out", map[string]int{})

	fmt.Println("Paths from you to out (1):", countWays("you", "out"))
	fmt.Println("Paths from svr to out passing through fft and dac (2):", ways)
}
