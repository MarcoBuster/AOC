package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func game(nums *[]int, iterations int) int {
	seen := make(map[int]int, iterations)
	initialLength := len(*nums)

	for i := 0; i < iterations; i++ {
		if i < initialLength {
			seen[(*nums)[i]] = i
			continue
		}
		k, ok := seen[(*nums)[i-1]]
		if ok {
			*nums = append(*nums, (i-1)-k)
		} else {
			*nums = append(*nums, 0)
		}
		seen[(*nums)[len(*nums)-2]] = i - 1
	}

	return (*nums)[len(*nums)-1]
}

func main() {
	f, _ := os.Open("input.txt")
	buf := make([]byte, 16)
	l, _ := f.Read(buf)
	numsStr := string(buf[:l-1])
	var nums []int
	for _, num := range strings.Split(numsStr, ",") {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}
	fmt.Println("Answer (part 1):", game(&nums, 2020))
	fmt.Println("Answer (part 2):", game(&nums, 30000000))
}
