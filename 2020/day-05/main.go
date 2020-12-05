package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func navigateTicket(r int, ticket string) int {
	var seats []int
	for i := 0; i <= r; i++ {
		seats = append(seats, i)
	}
	for _, c := range ticket {
		if c == 'F' || c == 'L' { // lower half
			seats = seats[:int(float64(len(seats))/float64(2)+0.5)]
		} else if c == 'B' || c == 'R' { // upper half
			seats = seats[len(seats)/2:]
		}
	}
	return seats[0]
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	high := 0
	var ticketIDs []int
	for b.Scan() {
		ticket := b.Text()
		row := navigateTicket(127, ticket[:7])
		column := navigateTicket(7, ticket[7:10])
		ticketID := row*8 + column
		if ticketID > high {
			high = ticketID
		}
		ticketIDs = append(ticketIDs, ticketID)
	}
	fmt.Println("Answer (part 1)", high)

	sort.Ints(ticketIDs)
	for i, t := range ticketIDs {
		if i+1 >= len(ticketIDs) {
			break
		}
		if t+1 != ticketIDs[i+1] && t+2 == ticketIDs[i+1] {
			fmt.Println("Answer (part 2):", t+1)
			break
		}
	}
}
