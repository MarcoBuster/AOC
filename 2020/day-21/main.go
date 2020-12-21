package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func parseInput(line string) ([]string, []string) {
	parts := strings.Split(line, " (contains ")
	ingredients := strings.Split(parts[0], " ")
	allergens := strings.Split(strings.Trim(parts[1], ")"), ", ")
	return ingredients, allergens
}

func intersect(sliceA, sliceB []string) []string {
	exists := make(map[string]bool)
	for _, v := range sliceA {
		exists[v] = true
	}
	var res []string
	for _, v := range sliceB {
		if exists[v] {
			res = append(res, v)
		}
	}
	return res
}

func removeFromSlice(slice []string, x string) []string {
	for i, v := range slice {
		if v == x {
			slice[i] = slice[len(slice)-1]
			slice = slice[:len(slice)-1]
		}
	}
	return slice
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	candidateIngredients := make(map[string][]string)
	ingredientCounts := make(map[string]int)

	for b.Scan() {
		line := b.Text()
		ingredients, allergens := parseInput(line)

		for _, i := range ingredients {
			ingredientCounts[i]++
		}

		for _, a := range allergens {
			if candidateIngredients[a] == nil {
				candidateIngredients[a] = ingredients
			} else {
				candidateIngredients[a] = intersect(candidateIngredients[a], ingredients)
			}
		}
	}

	for {
		done := true
		for allergen, possible := range candidateIngredients {
			if len(possible) != 1 {
				done = false
			} else {
				for allergen2, ingredients2 := range candidateIngredients {
					if allergen2 != allergen {
						candidateIngredients[allergen2] = removeFromSlice(ingredients2, possible[0])
					}
				}
			}
		}
		if done {
			break
		}
	}

	for _, names := range candidateIngredients {
		delete(ingredientCounts, names[0])
	}

	count := 0
	for _, ct := range ingredientCounts {
		count += ct
	}
	fmt.Println("Answer (part 1):", count)

	var names []string
	for k := range candidateIngredients {
		names = append(names, k)
	}
	sort.Strings(names) // alphabetic sort

	var part2Result []string
	for _, n := range names {
		part2Result = append(part2Result, candidateIngredients[n][0])
	}
	fmt.Println("Answer (part 2):", strings.Join(part2Result, ","))
}
