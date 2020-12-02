package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ParsedLine struct {
	min, max int
	char     rune
	password string
}

func parseLine(line string) ParsedLine {
	parsing := 0 // 0: min; 1: max; 2: char; 3: password
	var rawMin, rawMax, password string
	var char rune
	for _, c := range line {
		if c == '-' {
			parsing++
			continue
		}
		if c == ' ' {
			parsing++
			continue
		}
		if c == ':' {
			continue
		}

		switch parsing {
		case 0:
			rawMin += string(c)
		case 1:
			rawMax += string(c)
		case 2:
			char = c
		case 3:
			password += string(c)
		}
	}
	min, _ := strconv.Atoi(rawMin)
	max, _ := strconv.Atoi(rawMax)
	return ParsedLine{min, max, char, password}
}

func main() {
	b := bufio.NewScanner(os.Stdin)
	// var input []string
	part1CorrectCount := 0
	part2CorrectCount := 0
	for b.Scan() {
		parsed := parseLine(b.Text())
		count := 0
		for _, c := range parsed.password {
			if c == parsed.char {
				count++
			}
		}
		if parsed.min <= count && count <= parsed.max {
			part1CorrectCount++
		}

		// Part 2
		count = 0
		if parsed.min-1 < len(parsed.password) {
			if rune(parsed.password[parsed.min-1]) == parsed.char {
				count++
			}
		}
		if parsed.max-1 < len(parsed.password) {
			if rune(parsed.password[parsed.max-1]) == parsed.char {
				count++
			}
		}
		if count == 1 {
			part2CorrectCount++
		}
	}
	fmt.Println("Answer (part 1):", part1CorrectCount)
	fmt.Println("Answer (part 2):", part2CorrectCount)
}
