package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	var lines []string
	var i, x, trees int
	for b.Scan() {
		lines = append(lines, b.Text())
	}
	for _, line := range lines {
		fmt.Println(i, x%len(line))
		if line[x%len(line)] == '#' {
			trees++
		}
		x += 3
		i++
	}
	fmt.Println("Part 1", trees)

	// Part 2 - right 1 down 1
	trees, x = 0, 0
	for _, line := range lines {
		if line[x%len(line)] == '#' {
			trees++
		}
		x += 1
	}
	fmt.Println("Part 2A:", trees)

	// Part 2B - right 5 down 1
	trees, x = 0, 0
	for _, line := range lines {
		if line[x%len(line)] == '#' {
			trees++
		}
		x += 5
	}
	fmt.Println("Part 2B:", trees)

	// Part 2C - right 7 down 1
	trees, x = 0, 0
	for _, line := range lines {
		if line[x%len(line)] == '#' {
			trees++
		}
		x += 7
	}
	fmt.Println("Part 2C:", trees)

	// Part 2D - right 1 down 2
	trees, x, i = 0, 0, -1
	for _, line := range lines {
		i++
		if i%2 == 1 {
			continue
		}
		if line[x%len(line)] == '#' {
			trees++
		}
		x += 1
	}
	fmt.Println("Part 2D:", trees)
}
