package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func countDifferences(input *[]int) (delta1, delta3 int) {
	for i := 1; i < len(*input); i++ {
		if (*input)[i]-((*input)[i-1]) == 1 {
			delta1++
		} else if (*input)[i]-((*input)[i-1]) == 3 {
			delta3++
		}
	}
	return
}

func getDeletableDifferences(input *[]int) []int {
	deltaD := 0
	var valid []int
	for i := 1; i < len(*input); i++ {
		if i+1 >= len(*input) {
			continue
		}
		diff := (*input)[i+1] - ((*input)[i])
		if diff == 1 {
			if ((*input)[i] - (*input)[i-1]) > 2 {
				continue
			}
			deltaD++
			valid = append(valid, (*input)[i])
		}
	}
	return valid
}

func countArrangements(valid *[]int) float64 {
	subsequentCounter := 0
	subsequentThrees := 0
	for i := 1; i < len(*valid); i++ {
		if (*valid)[i]-(*valid)[i-1] == 1 {
			subsequentCounter++
		} else {
			if subsequentCounter > 1 {
				subsequentThrees++
			}
			subsequentCounter = 0
		}
	}
	if subsequentCounter > 1 {
		subsequentThrees++
	}
	notSubsequentThrees := len(*valid) - subsequentThrees*3
	return math.Pow(7, float64(subsequentThrees)) * math.Pow(2, float64(notSubsequentThrees))
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	var input []int
	for b.Scan() {
		line := b.Text()
		n, _ := strconv.Atoi(line)
		input = append(input, n)
	}
	sort.Ints(input)

	input = append([]int{0}, input...)
	input = append(input, input[len(input)-1]+3)
	delta1, delta3 := countDifferences(&input)
	fmt.Println("Answer (part 1):", delta1*delta3)

	valid := getDeletableDifferences(&input)
	arrangements := countArrangements(&valid)
	fmt.Printf("Answer (part 2): %.f\n", arrangements)
}
