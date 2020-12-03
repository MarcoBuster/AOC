package main

import (
	"bufio"
	"fmt"
	"os"
)

func countTrees(lines *[]string, deltaX, deltaY int) (trees int) {
	var x, y int
	for _, line := range *lines {
		if y%deltaY != 0 {
			y++
			continue
		}
		if line[x%len(line)] == '#' {
			trees++
		}
		x += deltaX
		y++
	}
	return
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	var lines []string
	for b.Scan() {
		lines = append(lines, b.Text())
	}

	part1 := countTrees(&lines, 3, 1)
	fmt.Println("Answer (part 1):", part1)

	part2A := countTrees(&lines, 1, 1)
	part2C := countTrees(&lines, 5, 1)
	part2D := countTrees(&lines, 7, 1)
	part2E := countTrees(&lines, 1, 2)
	fmt.Println("Answer (part 2):", part2A*part1*part2C*part2D*part2E)
}
