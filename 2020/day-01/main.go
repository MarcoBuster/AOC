package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	var input []int
	for b.Scan() {
		n, err := strconv.Atoi(b.Text())
		if err != nil {
			fmt.Println("Error!")
		}
		input = append(input, n)
	}

Part1Loop:
	for _, n := range input {
		for _, k := range input {
			if n+k == 2020 {
				fmt.Printf("Answer (part 1): %d\n", n*k)
				break Part1Loop
			}
		}
	}

Part2Loop:
	for _, n := range input {
		for _, k := range input {
			for _, j := range input {
				if n+k+j == 2020 {
					fmt.Printf("Answer (part 2): %d\n", n*k*j)
					break Part2Loop
				}
			}
		}
	}
}
