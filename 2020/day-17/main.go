package main

import (
	"bufio"
	"fmt"
	"os"
)

const ACTIVE = true
const INACTIVE = false

type Cube struct {
	x, y, z int
}
type State = map[Cube]bool
type HyperCube struct {
	x, y, z, w int
}
type HyperState = map[HyperCube]bool

func parseInputPart1(line string, state *State, y int) {
	for x, r := range line {
		cube := Cube{x, y, 0}
		if r == '#' {
			(*state)[cube] = ACTIVE
		} else {
			(*state)[cube] = INACTIVE
		}
	}
}

func parseInputPart2(line string, state *HyperState, y int) {
	for x, r := range line {
		cube := HyperCube{x, y, 0, 0}
		if r == '#' {
			(*state)[cube] = ACTIVE
		} else {
			(*state)[cube] = INACTIVE
		}
	}
}

func getNeighboursPart1(c Cube) []Cube {
	var res []Cube
	for _, z := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			for _, x := range []int{-1, 0, 1} {
				if z == 0 && y == 0 && x == 0 {
					continue
				}
				res = append(res, Cube{c.x + x, c.y + y, c.z + z})
			}
		}
	}
	return res
}

func getNeighboursPart2(c HyperCube) []HyperCube {
	var res []HyperCube
	for _, w := range []int{-1, 0, 1} {
		for _, z := range []int{-1, 0, 1} {
			for _, y := range []int{-1, 0, 1} {
				for _, x := range []int{-1, 0, 1} {
					if w == 0 && z == 0 && y == 0 && x == 0 {
						continue
					}
					res = append(res, HyperCube{c.x + x, c.y + y, c.z + z, c.w + w})
				}
			}
		}
	}
	return res
}

func countActiveNeighboursPart1(state *State, cube Cube) int {
	count := 0
	for _, c := range getNeighboursPart1(cube) {
		if (*state)[c] {
			count++
		}
	}
	return count
}

func countActiveNeighboursPart2(state *HyperState, cube HyperCube) int {
	count := 0
	for _, c := range getNeighboursPart2(cube) {
		if (*state)[c] {
			count++
		}
	}
	return count
}

func addOutsideCubesPart1(state *State) {
	var newCubes []Cube

	for cube := range *state {
		for _, neighbour := range getNeighboursPart1(cube) {
			_, exists := (*state)[neighbour]
			if !exists {
				newCubes = append(newCubes, neighbour)
			}
		}
	}
	for _, c := range newCubes {
		(*state)[c] = false
	}
}

func addOutsideCubesPart2(state *HyperState) {
	var newCubes []HyperCube

	for cube := range *state {
		for _, neighbour := range getNeighboursPart2(cube) {
			_, exists := (*state)[neighbour]
			if !exists {
				newCubes = append(newCubes, neighbour)
			}
		}
	}
	for _, c := range newCubes {
		(*state)[c] = false
	}
}

func cycleStatePart1(state *State) {
	nextState := make(State)
	addOutsideCubesPart1(state)

	for cube := range *state {
		count := countActiveNeighboursPart1(state, cube)
		nextState[cube] = ((*state)[cube] && count == 2) || count == 3
	}
	*state = nextState
}
func cycleStatePart2(state *HyperState) {
	nextState := make(HyperState)
	addOutsideCubesPart2(state)

	for cube := range *state {
		count := countActiveNeighboursPart2(state, cube)
		nextState[cube] = ((*state)[cube] && count == 2) || count == 3
	}
	*state = nextState
}

func countActiveCubesPart1(state State) int {
	count := 0
	for cube := range state {
		if (state)[cube] == ACTIVE {
			count++
		}
	}
	return count
}

func countActiveCubesPart2(state HyperState) int {
	count := 0
	for cube := range state {
		if (state)[cube] == ACTIVE {
			count++
		}
	}
	return count
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	state := make(State)
	hyperstate := make(HyperState)
	for y := 0; b.Scan(); y++ {
		line := b.Text()
		parseInputPart1(line, &state, y)
		parseInputPart2(line, &hyperstate, y)
	}

	for i := 0; i < 6; i++ {
		cycleStatePart1(&state)
		cycleStatePart2(&hyperstate)
	}
	fmt.Println("Answer (part 1):", countActiveCubesPart1(state))
	fmt.Println("Answer (part 2):", countActiveCubesPart2(hyperstate))
}
