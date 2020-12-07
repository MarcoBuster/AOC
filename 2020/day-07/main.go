package main

// Alternative solution for the first part with graphs:
// https://gist.github.com/MarcoBuster/652f4ce321638a798fa17f3b2f460be2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	content  string
	quantity int
}

func sanitizeRule(rule string) string {
	res := strings.ReplaceAll(rule, " bags", "")
	res = strings.ReplaceAll(res, " bag", "")
	res = strings.ReplaceAll(res, ".", "")
	res = strings.TrimSpace(res)
	return res
}

func parseRule(line string, root *map[string][]Rule) {
	header := sanitizeRule(strings.Split(line, "contain")[0])
	content := sanitizeRule(strings.Split(line, "contain")[1])
	rules := strings.Split(content, ", ")
	(*root)[header] = make([]Rule, len(rules))

	for _, rule := range rules {
		rule = sanitizeRule(rule)
		quantity, _ := strconv.Atoi(string(rule[0]))
		rule = rule[2:]
		if strings.Contains(rule, "other") {
			continue
		}
		(*root)[header] = append((*root)[header], Rule{rule, quantity})
	}
}

func alreadyInArray(x *[]string, element string) bool {
	for _, c := range *x {
		if element == c {
			return true
		}
	}
	return false
}

func traverseMapPart1(root *map[string][]Rule, query string, seen *[]string) int {
	count := 1
	for header, content := range *root {
		for _, rule := range content {
			if rule.content == query {
				if alreadyInArray(seen, header) {
					continue
				}
				count += traverseMapPart1(root, header, seen)
				*seen = append(*seen, header)
			}
		}
	}
	return count
}

func traverseMapPart2(root *map[string][]Rule, query string) int {
	count := 1
	for _, rule := range (*root)[query] {
		count += traverseMapPart2(root, rule.content) * rule.quantity
	}
	return count
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	i := 0
	rootMap := make(map[string][]Rule)
	for b.Scan() {
		line := b.Text()
		parseRule(line, &rootMap)
		i++
	}
	seen := make([]string, 0, i)
	part1 := traverseMapPart1(&rootMap, "shiny gold", &seen) - 1
	fmt.Println("Answer (part 1):", part1)
	part2 := traverseMapPart2(&rootMap, "shiny gold") - 1
	fmt.Println("Answer (part 2):", part2)
}
