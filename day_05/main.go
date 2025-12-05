package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ranges := [][2]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		split := strings.Split(line, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		ranges = append(ranges, [2]int{start, end})
	}

	fresh := 0

	for scanner.Scan() {
		id, _ := strconv.Atoi(scanner.Text())

		for _, r := range ranges {
			if r[0] <= id && id <= r[1] {
				fresh++
				break
			}
		}
	}

	fmt.Println("Fresh ingredient IDs (1):", fresh)

	lower := slices.MinFunc(ranges, func(r1 [2]int, r2 [2]int) int { return r1[0] - r2[0] })[0]
	higher := slices.MaxFunc(ranges, func(r1 [2]int, r2 [2]int) int { return r1[1] - r2[1] })[1]

	fresh = higher - lower + 1

	for curr := lower + 1; curr < higher; {
		index := slices.IndexFunc(ranges, func(r [2]int) bool {
			return r[0] <= curr && curr <= r[1]
		})

		if index != -1 {
			curr = ranges[index][1] + 1
		} else {
			min := higher

			for _, r := range ranges {
				if 0 < r[0]-curr && r[0]-curr < min {
					min = r[0] - curr
				}
			}

			fresh -= min
			curr += min
		}
	}

	fmt.Println("Fresh ingredient IDs (2):", fresh)
}
