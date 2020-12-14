package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseBuses(buses string) []int {
	var busList []string
	var res []int
	busList = strings.Split(buses, ",")
	for _, k := range busList {
		if k == "x" {
			k = "-1"
		}
		t, _ := strconv.Atoi(k)
		res = append(res, t)
	}
	return res
}

func findEarliestBus(target int, buses *[]int, start int) (int, int) {
	distances := make(map[int]int, len(*buses))
	for _, b := range *buses {
		if b == -1 {
			continue
		}
		for i := start; b*i < target+b; i++ {
			distances[b] = b * i
		}
	}
	minDistance := -1
	winner := 0
	for i, k := range distances {
		if k < minDistance || minDistance == -1 {
			minDistance = distances[i]
			winner = i
		}
	}
	return winner, minDistance - target
}

func part2(buses *[]int, monster int) int {
	result := monster
	mod := 1
	for i, b := range *buses {
		if b == -1 {
			continue
		}
		for (result+i)%b != 0 {
			result += mod
		}
		mod *= b
	}
	return result
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	b.Scan()
	timestamp, _ := strconv.Atoi(b.Text())
	b.Scan()
	busesLine := b.Text()
	buses := parseBuses(busesLine)
	bestBus, arrivalTime := findEarliestBus(timestamp, &buses, 0)
	fmt.Println("Answer (part 1):", bestBus*arrivalTime)

	monster := 100000000000000
	part2Res := part2(&buses, monster)
	fmt.Println("Answer (part 2):", part2Res)
}
