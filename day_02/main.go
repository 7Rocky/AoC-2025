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
	scanner.Scan()

	sum1 := 0
	sum2 := 0

	for r := range strings.SplitSeq(scanner.Text(), ",") {
		limits := strings.Split(r, "-")
		start, _ := strconv.Atoi(limits[0])
		end, _ := strconv.Atoi(limits[1])

		for id := start; id <= end; id++ {
			s := strconv.Itoa(id)

			if s[:len(s)/2] == s[len(s)/2:] {
				sum1 += id
			}

			for i := 1; i <= len(s)/2; i++ {
				if s == strings.Repeat(s[:i], len(s)/i) {
					sum2 += id
					break
				}
			}
		}
	}

	fmt.Println("Sum of invalid IDs (1):", sum1)
	fmt.Println("Sum of invalid IDs (2):", sum2)
}
