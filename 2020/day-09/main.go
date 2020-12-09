package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func checkNumber(preamble []int, n int) bool {
	for i := 0; i < len(preamble); i++ {
		for k := 0; k < len(preamble); k++ {
			if preamble[i]+preamble[k] == n {
				return true
			}
		}
	}
	return false
}

func findEncryptionWeakness(input []int, target *int) (int, int) {
	sum := 0
	max := 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
		if sum == *target {
			max = i
			break
		}
		if sum > *target {
			return -1, -1
		}
	}

	input = input[:max]
	sort.Ints(input)
	return input[0], input[len(input)-1]
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	var input []int
	for b.Scan() {
		line := b.Text()
		number, _ := strconv.Atoi(line)
		input = append(input, number)
	}

	var part1 int
	for i := 25; i < len(input); i++ {
		res := checkNumber(input[i-25:i], input[i])
		if !res {
			part1 = input[i]
			break
		}
	}
	fmt.Println("Answer (part 1):", part1)

	for i := 0; i < len(input); i++ {
		min, max := findEncryptionWeakness(input[i:], &part1)
		if min != -1 {
			fmt.Println("Answer (part 2):", min+max)
			break
		}
	}
}
