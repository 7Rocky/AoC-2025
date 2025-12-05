package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(arr []int) (sum int) {
	for _, a := range arr {
		sum += a
	}

	return
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	diagram := [][]int{{}}

	for scanner.Scan() {
		row := []int{0}

		for _, c := range scanner.Text() {
			if c == '@' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}

		diagram = append(diagram, append(row, 0))
	}

	diagram = append(diagram, []int{})

	for range len(diagram[1]) {
		diagram[0] = append(diagram[0], 0)
		diagram[len(diagram)-1] = append(diagram[len(diagram)-1], 0)
	}

	accessible := 1
	total := 0

	for accessible != 0 {
		accessible = 0
		toRemove := []complex64{}

		for i := 1; i < len(diagram)-1; i++ {
			for j := 1; j < len(diagram[i])-1; j++ {
				if diagram[i][j] == 1 && sum(diagram[i-1][j-1:j+2])+sum(diagram[i][j-1:j+2])+sum(diagram[i+1][j-1:j+2]) < 5 {
					accessible++
					toRemove = append(toRemove, complex(float32(i), float32(j)))
				}
			}
		}

		if total == 0 {
			fmt.Println("Accessible rolls (1):", accessible)
		}

		total += accessible

		for _, r := range toRemove {
			diagram[int(real(r))][int(imag(r))] = 0
		}
	}

	fmt.Println("Accessible rolls (2):", total)
}
