package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ParsedLine struct {
	min, max int
	char     rune
	password string
}

func parseLine(line string) ParsedLine {
	minString := strings.Split(line, "-")[0]
	maxString := strings.Split(strings.Split(line, "-")[1], " ")[0]
	charString := strings.Split(strings.Split(line, ":")[0], " ")[1]
	password := strings.Split(line, " ")[2]

	min, _ := strconv.Atoi(minString)
	max, _ := strconv.Atoi(maxString)
	char := rune(charString[0])
	return ParsedLine{min, max, char, password}
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
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
