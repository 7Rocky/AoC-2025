package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type machine struct {
	target       []int
	buttons      [][]int
	requirements []int
}

func sum(arr []int) (s int) {
	for _, a := range arr {
		s += a
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func parseMachine(line string) (m machine) {
	split := strings.Split(line, " ")
	m.target = stringToTarget(split[0])

	for i := 1; i < len(split)-1; i++ {
		s := split[i][1 : len(split[i])-1]
		button := make([]int, strings.Count(s, ",")+1)
		m.buttons = append(m.buttons, button)

		for i, d := range strings.Split(s, ",") {
			fmt.Sscanf(d, "%d", &button[i])
		}
	}

	s := split[len(split)-1][1 : len(split[len(split)-1])-1]
	m.requirements = make([]int, strings.Count(s, ",")+1)

	for i, d := range strings.Split(s, ",") {
		fmt.Sscanf(d, "%d", &m.requirements[i])
	}

	return
}

func stringToTarget(s string) (target []int) {
	for _, c := range s {
		if c == '.' {
			target = append(target, 0)
		} else if c == '#' {
			target = append(target, 1)
		}
	}

	return
}

func targetToString(target []int) string {
	var b []byte

	for _, t := range target {
		if t%2 == 1 {
			b = append(b, '#')
		} else {
			b = append(b, '.')
		}
	}

	return string(b)
}

func decimalToBinary(n int, length int) (bits []int) {
	for n != 0 {
		bits = append(bits, n%2)
		n /= 2
	}

	for len(bits) < length {
		bits = append(bits, 0)
	}

	return
}

func (m *machine) power() int {
	visited := map[string]struct{}{}
	queue := []string{targetToString(decimalToBinary(0, len(m.target)))}
	combinations := map[string]*[][]int{queue[0]: {}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := visited[curr]; ok {
			continue
		}

		visited[curr] = struct{}{}

		for _, button := range m.buttons {
			state := stringToTarget(curr)

			for _, s := range button {
				state[s] = 1 - state[s]
			}

			next := targetToString(state)
			queue = append(queue, next)

			if p, ok := combinations[next]; ok {
				if len(*p) > len(*combinations[curr])+1 {
					*combinations[next] = append(*combinations[curr], button)
				}
			} else {
				combinations[next] = &[][]int{}
				*combinations[next] = append(*combinations[curr], button)
			}
		}
	}

	return len(*combinations[targetToString(m.target)])
}

var cache = map[string][][]int{}

func (m *machine) waysToTarget(target []int) [][]int {
	if res, ok := cache[targetToString(target)]; ok {
		return res
	}

	var combinations [][]int

	for n := range 1 << len(m.buttons) {
		state := make([]int, len(target))
		binary := decimalToBinary(n, len(m.buttons))

		for i, press := range binary {
			if press == 1 {
				for _, button := range m.buttons[i] {
					state[button] = 1 - state[button]
				}
			}
		}

		if slices.EqualFunc(state, target, func(s, t int) bool { return s == t%2 }) {
			combinations = append(combinations, binary)
		}
	}

	cache[targetToString(target)] = combinations
	return combinations
}

func (m *machine) powerWithRequirements(presses []int, level int, total *int, path *[]int) {
	if sum(presses) == 0 {
		res := 0

		for i, p := range (*path)[:level] {
			res += p * (1 << i)
		}

		*total = min(*total, res)
		return
	}

	if len(*path) < level+1 {
		*path = append(*path, math.MaxInt)
	}

	newPresses := make([]int, len(presses))

	for _, way := range m.waysToTarget(presses) {
		copy(newPresses, presses)
		(*path)[level] = sum(way)

		for i, b := range way {
			for _, button := range m.buttons[i] {
				newPresses[button] -= b
			}
		}

		for i := range newPresses {
			newPresses[i] /= 2
		}

		if !slices.ContainsFunc(newPresses, func(p int) bool { return p < 0 }) {
			m.powerWithRequirements(newPresses, level+1, total, path)
		}
	}
}

func main() {
	b, _ := os.ReadFile("input.txt")
	s := [2]int{0, 0}

	for line := range strings.SplitSeq(string(b), "\n") {
		m := parseMachine(line)
		s[0] += m.power()
		t := math.MaxInt
		cache = map[string][][]int{}
		m.powerWithRequirements(m.requirements, 0, &t, &[]int{})
		s[1] += t
	}

	fmt.Println("Fewest button presses (1):", s[0])
	fmt.Println("Fewest button presses (2):", s[1])
}
