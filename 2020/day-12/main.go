package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const NORTH rune = 'N'
const SOUTH rune = 'S'
const EAST rune = 'E'
const WEST rune = 'W'
const LEFT rune = 'L'
const RIGHT rune = 'R'
const FORWARD rune = 'F'

var MoveMap = map[rune]Instruction{
	NORTH:   {0, 1, 0, 0},
	SOUTH:   {0, -1, 0, 0},
	EAST:    {1, 0, 0, 0},
	WEST:    {-1, 0, 0, 0},
	LEFT:    {0, 0, 0, -1},
	RIGHT:   {0, 0, 0, 1},
	FORWARD: {0, 0, 1, 0},
}
var FacingNext = []rune{
	NORTH,
	EAST,
	SOUTH,
	WEST,
}

type Ship struct {
	posX, posY int
	facing     rune
}

type Waypoint struct {
	deltaX, deltaY int
}

type Instruction struct {
	deltaX       int
	deltaY       int
	deltaForward int
	rotation     int
}

func easyAbs(a int) int {
	return int(math.Abs(float64(a)))
}

func parseInstruction(instruction string) Instruction {
	direction := rune(instruction[0])
	multiplier, _ := strconv.Atoi(instruction[1:])
	res := MoveMap[direction]
	res.deltaX *= multiplier
	res.deltaY *= multiplier
	res.deltaForward *= multiplier
	res.rotation *= multiplier / 90
	return res
}

func nextFacing(curFacing rune, deltaFacing int) rune {
	curIndex := 0
	for i := 0; i < len(FacingNext); i++ {
		if curFacing == FacingNext[i] {
			curIndex = i
			break
		}
	}
	nextIndex := (curIndex + deltaFacing) % 4
	if nextIndex < 0 {
		nextIndex = 4 - easyAbs(nextIndex)
	}
	return FacingNext[nextIndex]
}

func move(ship *Ship, instruction Instruction) {
	ship.posX += instruction.deltaX
	ship.posY += instruction.deltaY
	if instruction.deltaForward > 0 {
		instruction2 := MoveMap[ship.facing]
		ship.posX += instruction2.deltaX * instruction.deltaForward
		ship.posY += instruction2.deltaY * instruction.deltaForward
	}
	if instruction.rotation != 0 {
		ship.facing = nextFacing(ship.facing, instruction.rotation)
	}
}

func shipToWaypoint(ship *Ship, waypoint *Waypoint, multiplier int) {
	ship.posX += waypoint.deltaX * multiplier
	ship.posY += waypoint.deltaY * multiplier
}

func rotateWaypoint(waypoint *Waypoint, rotation int) {
	if rotation < 0 {
		for i := rotation; i != 0; i++ {
			waypoint.deltaX, waypoint.deltaY = -waypoint.deltaY, waypoint.deltaX
		}
	} else if rotation > 0 {
		for i := rotation; i != 0; i-- {
			waypoint.deltaX, waypoint.deltaY = waypoint.deltaY, -waypoint.deltaX
		}
	}
}

func move2(ship *Ship, waypoint *Waypoint, instruction Instruction) {
	waypoint.deltaX += instruction.deltaX
	waypoint.deltaY += instruction.deltaY
	if instruction.deltaForward > 0 {
		shipToWaypoint(ship, waypoint, instruction.deltaForward)
	}
	if instruction.rotation != 0 {
		rotateWaypoint(waypoint, instruction.rotation)
	}
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	ship := Ship{0, 0, EAST}
	var instructions []string
	for b.Scan() {
		line := b.Text()
		instruction := parseInstruction(line)
		instructions = append(instructions, line)
		move(&ship, instruction)
	}
	distance := easyAbs(ship.posX) + easyAbs(ship.posY)
	fmt.Println("Answer (part 1):", distance)

	ship = Ship{0, 0, EAST}
	waypoint := Waypoint{10, 1}
	for _, instruction := range instructions {
		move2(&ship, &waypoint, parseInstruction(instruction))
	}
	distance = easyAbs(ship.posX) + easyAbs(ship.posY)
	fmt.Println("Answer (part 2):", distance)
}
