package main

import (
	"bufio"
	"fmt"
	"os"
)

type Instruction struct {
	dx, dy, dz int
}
type Grid = map[int]map[int]map[int]bool

// https://www.redblobgames.com/grids/hexagons/
func parseInstruction(delimiter string) (int, int, int) {
	var dx, dy, dz int
	switch delimiter {
	case "e":
		dx, dy, dz = 1, -1, 0
	case "w":
		dx, dy, dz = -1, 1, 0
	case "ne":
		dx, dy, dz = 1, 0, -1
	case "nw":
		dx, dy, dz = 0, 1, -1
	case "se":
		dx, dy, dz = 0, -1, 1
	case "sw":
		dx, dy, dz = -1, 0, 1
	}
	return dx, dy, dz
}

func parseInput(line string) []Instruction {
	instructions := make([]Instruction, 0)
	var delimiter string
	for i := 0; i < len(line); i++ {
		if i > 0 && (line[i-1] == 's' || line[i-1] == 'n') {
			continue
		}
		if line[i] == 's' || line[i] == 'n' {
			delimiter = string(line[i]) + string(line[i+1])
		} else {
			delimiter = string(line[i])
		}
		dx, dy, dz := parseInstruction(delimiter)
		instructions = append(instructions, Instruction{dx, dy, dz})
	}
	return instructions
}

func getCoordinates(instructions []Instruction) (int, int, int) {
	var x, y, z int
	for _, instruction := range instructions {
		x += instruction.dx
		y += instruction.dy
		z += instruction.dz
	}
	return x, y, z
}

func run(instructions []Instruction, grid *Grid) {
	x, y, z := getCoordinates(instructions)
	if len((*grid)[x]) == 0 {
		(*grid)[x] = make(map[int]map[int]bool)
	}
	if len((*grid)[x][y]) == 0 {
		(*grid)[x][y] = make(map[int]bool)
	}
	_, ok := (*grid)[x][y][z]
	if ok {
		(*grid)[x][y][z] = !(*grid)[x][y][z]
	} else {
		(*grid)[x][y][z] = false // white = true; black = false

	}
}

func countBlackTiles(grid Grid) int {
	count := 0
	for x := range grid {
		for y := range grid[x] {
			for z := range grid[x][y] {
				if !grid[x][y][z] {
					count++
				}
			}
		}
	}
	return count
}

// Populate the grid with MUCH white tiles
func bootstrapGrid(grid *Grid) {
	newGrid := make(Grid)
	bigNumber := 100 // arbitrary number - it just works™
	for x := -bigNumber; x < bigNumber; x++ {
		if len(newGrid[x]) == 0 {
			newGrid[x] = make(map[int]map[int]bool)
		}
		for y := -bigNumber; y < bigNumber; y++ {
			if len(newGrid[x][y]) == 0 {
				newGrid[x][y] = make(map[int]bool)
			}
			for z := -bigNumber; z < bigNumber; z++ {
				_, ok := (*grid)[x][y][z]
				if ok {
					newGrid[x][y][z] = (*grid)[x][y][z]
				} else {
					newGrid[x][y][z] = true
				}
			}
		}
	}
	*grid = newGrid
}

func getElement(grid Grid, x, y, z int) (bool, bool) { // element, found
	_, ok := grid[x]
	if !ok {
		return false, false
	}
	_, ok = grid[x][y]
	if !ok {
		return false, false
	}
	_, ok = grid[x][y][z]
	if !ok {
		return false, false
	}
	return grid[x][y][z], true
}

func setElement(grid *Grid, x, y, z int, el bool) {
	if len((*grid)[x]) == 0 {
		(*grid)[x] = make(map[int]map[int]bool)
	}
	if len((*grid)[x][y]) == 0 {
		(*grid)[x][y] = make(map[int]bool)
	}
	(*grid)[x][y][z] = el
}

func countAdjacentBlackTiles(grid Grid, x, y, z int) int {
	countBlack := 0
	for _, s := range []string{"e", "w", "ne", "nw", "se", "sw"} {
		dx, dy, dz := parseInstruction(s)
		el, ok := getElement(grid, x+dx, y+dy, z+dz)
		if !ok {
			continue
		}
		if el == false {
			countBlack++
		}
	}
	return countBlack
}

func flipTiles(grid *Grid) {
	newGrid := make(Grid)
	for x := range *grid {
		for y := range (*grid)[x] {
			for z := range (*grid)[x][y] {
				el := (*grid)[x][y][z]
				adjacentCount := countAdjacentBlackTiles(*grid, x, y, z)
				if el == false && (adjacentCount == 0 || adjacentCount > 2) {
					el = true
				} else if el == true && adjacentCount == 2 {
					el = false
				}
				setElement(&newGrid, x, y, z, el)
			}
		}
	}
	*grid = newGrid
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	grid := make(Grid)
	for b.Scan() {
		line := b.Text()
		instructions := parseInput(line)
		run(instructions, &grid)
	}
	fmt.Println("Answer (part 1):", countBlackTiles(grid))

	bootstrapGrid(&grid)
	for i := 0; i < 100; i++ {
		flipTiles(&grid)
	}
	fmt.Println("Answer (part 2):", countBlackTiles(grid))
}
