package main

import (
	"bufio"
	"fmt"
	"os"
)

const EMPTY rune = 'L'
const OCCUPIED rune = '#'

func countAdjacentSeatsPart1(input [][]rune, x, y int, seatType rune) int {
	count := 0
	if x > 0 && input[y][x-1] == seatType { // left
		count++
	}
	if x < len(input[y])-1 && input[y][x+1] == seatType { // right
		count++
	}
	if y > 0 && input[y-1][x] == seatType { // up
		count++
	}
	if y < len(input)-1 && input[y+1][x] == seatType { // down
		count++
	}
	if x > 0 && y > 0 && input[y-1][x-1] == seatType { // up-left
		count++
	}
	if x > 0 && y < len(input)-1 && input[y+1][x-1] == seatType { // bottom-left
		count++
	}
	if x < len(input[y])-1 && y > 0 && input[y-1][x+1] == seatType { // up-right
		count++
	}
	if x < len(input[y])-1 && y < len(input)-1 && input[y+1][x+1] == seatType { // bottom-right
		count++
	}
	return count
}

func countAdjacentSeatsPart2(input [][]rune, x, y int, seatType rune) int {
	count := 0
	for i, k := y+1, x-1; i < len(input) && k >= 0; i, k = i+1, k-1 { // going bottom-left
		if input[i][k] == EMPTY {
			break
		}
		if input[i][k] == seatType { // bottom-left
			count++
			break
		}
	}
	for i, k := y+1, x+1; i < len(input) && k < len(input[y]); i, k = i+1, k+1 { // going bottom-right
		if input[i][k] == EMPTY {
			break
		}
		if input[i][k] == seatType { // bottom-right
			count++
			break
		}
	}
	for i, k := y-1, x-1; i >= 0 && k >= 0; i, k = i-1, k-1 { // going up-left
		if input[i][k] == EMPTY {
			break
		}
		if input[i][k] == seatType { // up-left
			count++
			break
		}
	}
	for i, k := y-1, x+1; i >= 0 && k < len(input[y]); i, k = i-1, k+1 { // going up-right
		if input[i][k] == EMPTY {
			break
		}
		if input[i][k] == seatType { // up-right
			count++
			break
		}
	}
	for i := y + 1; i < len(input); i++ { // going down
		if input[i][x] == EMPTY {
			break
		}
		if input[i][x] == seatType { // bottom
			count++
			break
		}
	}
	for i := y - 1; i >= 0; i-- { // going up
		if input[i][x] == EMPTY {
			break
		}
		if input[i][x] == seatType { // up
			count++
			break
		}
	}
	for k := x - 1; k >= 0; k-- { // going left
		if input[y][k] == EMPTY {
			break
		}
		if input[y][k] == seatType { // left
			count++
			break
		}
	}
	for k := x + 1; k < len(input[y]); k++ { // going right
		if input[y][k] == EMPTY {
			break
		}
		if input[y][k] == seatType { // right
			count++
			break
		}
	}
	return count
}

func printSeats(seats *[][]rune) {
	for i := 0; i < len(*seats); i++ {
		for k := 0; k < len((*seats)[i]); k++ {
			fmt.Printf("%c", (*seats)[i][k])
		}
		fmt.Println()
	}
	fmt.Println("-------------------------------")
}

func iterateSeatsPart1(seats *[][]rune) {
	var futureSeats [][]rune

	for i := 0; i < len(*seats); i++ {
		futureSeats = append(futureSeats, make([]rune, len((*seats)[i])))
		for k := 0; k < len((*seats)[i]); k++ {
			if (*seats)[i][k] == EMPTY && countAdjacentSeatsPart1(*seats, k, i, OCCUPIED) == 0 {
				futureSeats[i][k] = OCCUPIED
			} else if (*seats)[i][k] == OCCUPIED && countAdjacentSeatsPart1(*seats, k, i, OCCUPIED) >= 4 {
				futureSeats[i][k] = EMPTY
			} else {
				futureSeats[i][k] = (*seats)[i][k]
			}
		}
	}
	for i := 0; i < len(*seats); i++ {
		for k := 0; k < len((*seats)[i]); k++ {
			(*seats)[i][k] = futureSeats[i][k]
		}
	}
}

func iterateSeatsPart2(seats *[][]rune) {
	var futureSeats [][]rune

	for i := 0; i < len(*seats); i++ {
		futureSeats = append(futureSeats, make([]rune, len((*seats)[i])))
		for k := 0; k < len((*seats)[i]); k++ {
			if (*seats)[i][k] == EMPTY && countAdjacentSeatsPart2(*seats, k, i, OCCUPIED) == 0 {
				futureSeats[i][k] = OCCUPIED
			} else if (*seats)[i][k] == OCCUPIED && countAdjacentSeatsPart2(*seats, k, i, OCCUPIED) >= 5 {
				futureSeats[i][k] = EMPTY
			} else {
				futureSeats[i][k] = (*seats)[i][k]
			}
		}
	}
	for i := 0; i < len(*seats); i++ {
		for k := 0; k < len((*seats)[i]); k++ {
			(*seats)[i][k] = futureSeats[i][k]
		}
	}
}

func countOccupiedSeats(seats *[][]rune) int {
	count := 0
	for i := 0; i < len(*seats); i++ {
		for k := 0; k < len((*seats)[i]); k++ {
			if (*seats)[i][k] == OCCUPIED {
				count++
			}
		}
	}
	return count
}

func deepCopy(a *[][]rune) [][]rune {
	var copiedArray [][]rune
	for i := 0; i < len(*a); i++ {
		copiedArray = append(copiedArray, make([]rune, len((*a)[i])))
		for k := 0; k < len((*a)[i]); k++ {
			copiedArray[i][k] = (*a)[i][k]
		}
	}
	return copiedArray
}

func equals(a, b [][]rune) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	for i := 0; i < len(a); i++ {
		for k := 0; k < len((a)[i]); k++ {
			if (a)[i][k] != (b)[i][k] {
				return false
			}
		}
	}
	return true
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	var input [][]rune
	for b.Scan() {
		line := b.Text()
		input = append(input, []rune(line))
	}

	var seats, oldInput [][]rune
	seats = deepCopy(&input)
	for !equals(oldInput, seats) {
		oldInput = deepCopy(&seats)
		iterateSeatsPart1(&seats)
	}
	part1 := countOccupiedSeats(&seats)
	fmt.Println("Answer (part 1):", part1)

	seats = deepCopy(&input)
	for !equals(oldInput, seats) {
		oldInput = deepCopy(&seats)
		iterateSeatsPart2(&seats)
	}
	part2 := countOccupiedSeats(&seats)
	fmt.Println("Answer (part 2):", part2)
}
