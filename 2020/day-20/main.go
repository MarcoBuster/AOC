package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Tile struct {
	id       int
	contents [][]rune
}
type AdjacentBorder [2]int
type AdjacentBorders []AdjacentBorder

func parseTile(tile string) Tile {
	expr := regexp.MustCompile("(\\d+):")
	tileID, _ := strconv.Atoi(expr.FindStringSubmatch(tile)[1])
	tileString := strings.Split(tile, ":")[1]
	tileContents := make([][]rune, 11)
	y := 0
	for x, c := range tileString {
		tileContents[y] = append(tileContents[y], c)
		if (x+1)%10 == 0 && x != 0 {
			y++
		}
	}
	return Tile{tileID, tileContents[:len(tileContents)-1]}
}

func getBorders(tile Tile) [][]rune {
	var borders [][]rune
	borders = append(borders, tile.contents[0])
	borders = append(borders, tile.contents[len(tile.contents)-1])
	var leftBorder, rightBorder []rune
	for _, row := range tile.contents {
		leftBorder = append(leftBorder, row[0])
		rightBorder = append(rightBorder, row[len(row)-1])
	}
	borders = append(borders, leftBorder)
	borders = append(borders, rightBorder)
	return borders
}

func runeSliceToStr(slice []rune) string {
	res := ""
	for i := range slice {
		res += fmt.Sprintf("%c ", slice[i])
	}
	return res
}

func checkBorderEq(borderA, borderB []rune) bool {
	for i := range borderA {
		if borderA[i] != borderB[i] {
			return false
		}
	}
	return true
}

func compareTiles(tileA, tileB Tile) int {
	bordersA := getBorders(tileA)
	bordersB := getBorders(tileB)
	for i := range bordersA {
		for range bordersB {
			if checkBorderEq(bordersA[i], bordersB[i]) {
				return i
			}
		}
	}
	return -1
}

func (t Tile) Print() {
	for row := range t.contents {
		fmt.Println(runeSliceToStr(t.contents[row]))
	}
	fmt.Println()
}

func (t *Tile) Rotate() { // rotates 90° clockwise
	newContents := make([][]rune, 10)
	for row := range t.contents {
		for col := range t.contents {
			if len(newContents[col]) == 0 {
				newContents[col] = make([]rune, 10)
			}
			newContents[row][col] = t.contents[10-col-1][row]
		}
	}
	(*t).contents = newContents
}

func (t *Tile) VMirror() { // vertical mirror
	newContents := make([][]rune, 10)
	for row := range t.contents {
		for col := range t.contents {
			if len(newContents[col]) == 0 {
				newContents[col] = make([]rune, 10)
			}
			newContents[row][col] = t.contents[row][10-col-1]
		}
	}
	(*t).contents = newContents
}

func (t *Tile) HMirror() { // horizontal mirror
	newContents := make([][]rune, 10)
	for row := range t.contents {
		for col := range t.contents {
			if len(newContents[col]) == 0 {
				newContents[col] = make([]rune, 10)
			}
			newContents[row][col] = t.contents[10-row-1][col]
		}
	}
	(*t).contents = newContents
}

func (t *Tile) DMirror() {
	(*t).Rotate()
	(*t).VMirror()
}

func borderAlreadyFound(t1, t2 Tile, foundBorders AdjacentBorders) bool {
	for _, b := range foundBorders {
		if (t1.id == b[0] && t2.id == b[1]) || (t1.id == b[1] && t2.id == b[0]) {
			return true
		}
	}
	return false
}

func findBorders(tiles *[]Tile) (AdjacentBorders, map[int]int) {
	var borders AdjacentBorders
	countBorders := make(map[int]int)
	for _, t1 := range *tiles {
		var candidates []Tile
		for _, t2 := range *tiles {
			if t1.id == t2.id {
				continue
			}
			for i1 := 0; i1 < 4; i1++ {
				t2.VMirror()
				for i2 := 0; i2 < 4; i2++ {
					t2.HMirror()
					for i3 := 0; i3 < 4; i3++ {
						t2.DMirror()
						res := compareTiles(t1, t2)
						if res == -1 || borderAlreadyFound(t1, t2, borders) {
							continue
						}
						candidates = append(candidates, t2)
						borders = append(borders, AdjacentBorder{t1.id, t2.id})
						countBorders[t1.id]++
						countBorders[t2.id]++
					}
				}
			}
		}
	}
	return borders, countBorders
}

func tileByID(tileID int, tiles *[]Tile) Tile {
	for i := range *tiles {
		if tileID == (*tiles)[i].id {
			return (*tiles)[i]
		}
	}
	return Tile{}
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	var tiles []Tile
	var prevTile string
	for b.Scan() {
		line := b.Text()
		if line == "" {
			tiles = append(tiles, parseTile(prevTile))
			prevTile = ""
			continue
		}
		prevTile += line
	}
	_, countBorders := findBorders(&tiles)
	mul := 1
	var corners []Tile
	for i := range countBorders {
		if countBorders[i] == 2 {
			mul *= i
			corners = append(corners, tileByID(i, &tiles))
		}
	}
	fmt.Println("Answer (part 1):", mul)
}
