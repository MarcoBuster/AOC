package main

import (
	"bufio"
	"fmt"
	"os"
)

func alreadyInArray(x []rune, element rune) bool {
	for _, c := range x {
		if element == c {
			return true
		}
	}
	return false
}

func countAnswersPart1(answers string) int {
	var seen []rune
	for _, answer := range answers {
		if !alreadyInArray(seen, answer) {
			seen = append(seen, answer)
		}
	}
	return len(seen)
}

func countAnswersPart2(answers string) int {
	charCounts := make(map[rune]int, 0)
	var seen []rune
	nSep := 0
	for _, answer := range answers {
		if answer == '|' {
			nSep++
			seen = []rune{}
			continue
		}
		if alreadyInArray(seen, answer) {
			continue
		}
		seen = append(seen, answer)
		charCounts[answer]++
	}

	res := 0
	for _, c := range charCounts {
		if c == nSep {
			res++
		}
	}
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	var answerGroup string
	answerPart1 := 0
	answerPart2 := 0
	for b.Scan() {
		line := b.Text()
		if line == "" {
			answerPart1 += countAnswersPart1(answerGroup)
			answerPart2 += countAnswersPart2(answerGroup)
			answerGroup = ""
			continue
		}
		answerGroup += "|" + line
	}
	fmt.Println("Answer (part 1):", answerPart1)
	fmt.Println("Answer (part 2):", answerPart2)
}
