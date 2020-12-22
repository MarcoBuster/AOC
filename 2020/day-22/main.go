package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Deck []int

func combat(player1, player2 *Deck) int {
	if len(*player1) == 0 {
		return 2
	} else if len(*player2) == 0 {
		return 1
	}
	if (*player1)[0] > (*player2)[0] {
		*player1 = append((*player1)[1:], (*player1)[0], (*player2)[0])
		*player2 = (*player2)[1:]
	} else {
		*player2 = append((*player2)[1:], (*player2)[0], (*player1)[0])
		*player1 = (*player1)[1:]
	}
	return combat(player1, player2)
}

func calculateScore(deck Deck) int {
	sum := 0
	for i, m := len(deck)-1, 1; i >= 0; i, m = i-1, m+1 {
		sum += deck[i] * m
	}
	return sum
}

func eqDeck(sl1, sl2 Deck) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i := range sl1 {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

func alreadyPlayed(archival []Deck, deck Deck) bool {
	for _, d := range archival {
		if eqDeck(d, deck) {
			return true
		}
	}
	return false
}

func clearArchive(archival *map[int]map[int][]Deck, depth int) {
	(*archival)[depth] = make(map[int][]Deck)
}

func recursiveCombat(player1, player2 *Deck, archival *map[int]map[int][]Deck, depth int) int {
	if len((*archival)[depth]) == 0 {
		(*archival)[depth] = make(map[int][]Deck)
	}
	if len((*archival)[depth][0]) == 0 {
		(*archival)[depth][0] = make([]Deck, 0)
		(*archival)[depth][1] = make([]Deck, 0)
	}
	// Infinite game prevention. Nobody likes infinite games.
	if alreadyPlayed((*archival)[depth][0], *player1) || alreadyPlayed((*archival)[depth][1], *player2) {
		clearArchive(archival, depth)
		return 1
	}
	(*archival)[depth][0] = append((*archival)[depth][0], *player1)
	(*archival)[depth][1] = append((*archival)[depth][1], *player2)

	if len(*player1) == 0 {
		clearArchive(archival, depth)
		return 2
	} else if len(*player2) == 0 {
		clearArchive(archival, depth)
		return 1
	}

	winner := -1
	if len(*player1)-1 >= (*player1)[0] && len(*player2)-1 >= (*player2)[0] {
		player1Copy := make(Deck, (*player1)[0])
		player2Copy := make(Deck, (*player2)[0])
		copy(player1Copy, (*player1)[1:])
		copy(player2Copy, (*player2)[1:])
		winner = recursiveCombat(&player1Copy, &player2Copy, archival, depth+1)
		if winner == 2 {
			*player2 = append((*player2)[1:], (*player2)[0], (*player1)[0])
			*player1 = (*player1)[1:]
		} else {
			*player1 = append((*player1)[1:], (*player1)[0], (*player2)[0])
			*player2 = (*player2)[1:]
		}
	} else {
		if winner == 1 || (*player1)[0] > (*player2)[0] {
			*player1 = append((*player1)[1:], (*player1)[0], (*player2)[0])
			*player2 = (*player2)[1:]
		} else if winner == 2 || (*player2)[0] > (*player1)[0] {
			*player2 = append((*player2)[1:], (*player2)[0], (*player1)[0])
			*player1 = (*player1)[1:]
		}
	}
	return recursiveCombat(player1, player2, archival, depth)
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	player1 := make(Deck, 0)
	player2 := make(Deck, 0)
	player := 0
	for b.Scan() {
		line := b.Text()
		if line == "Player 1:" {
			player = 1
			continue
		} else if line == "Player 2:" {
			player = 2
			continue
		} else if line == "" {
			continue
		}
		n, _ := strconv.Atoi(line)
		if player == 1 {
			player1 = append(player1, n)
		} else {
			player2 = append(player2, n)
		}
	}
	player1Part1 := make(Deck, len(player1))
	player2Part1 := make(Deck, len(player2))
	copy(player1Part1, player1)
	copy(player2Part1, player2)

	winner := make(Deck, len(player1Part1)+len(player2Part1))
	if combat(&player1Part1, &player2Part1) == 1 {
		copy(winner, player1Part1)
	} else {
		copy(winner, player2Part1)
	}
	score := calculateScore(winner)
	fmt.Println("Answer (part 1):", score)

	player1Part2 := make(Deck, len(player1))
	player2Part2 := make(Deck, len(player2))
	copy(player1Part2, player1)
	copy(player2Part2, player2)

	archival := make(map[int]map[int][]Deck)
	winner = make(Deck, len(player1)+len(player2))
	if recursiveCombat(&player1Part2, &player2Part2, &archival, 0) == 1 {
		copy(winner, player1Part2)
	} else {
		copy(winner, player2Part2)
	}
	score = calculateScore(winner)
	fmt.Println("Answer (part 2):", score)
}
