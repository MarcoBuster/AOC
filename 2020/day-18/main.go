package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SUM = '+'
const MUL = '*'

func evaluateExpr(expr string) (int, int) {
	res := 0
	mode := SUM
	for i := 0; i < len(expr); i++ {
		c := rune(expr[i])
		if c == ')' {
			return res, i
		} else if c == '(' {
			recursiveRes, newIndex := evaluateExpr(expr[i+1:])
			if mode == SUM {
				res += recursiveRes
			} else if mode == MUL {
				res *= recursiveRes
			}
			i += newIndex + 1
			continue
		} else if c == SUM || c == MUL {
			mode = c
		} else {
			n, _ := strconv.Atoi(string(c))
			if mode == SUM {
				res += n
			} else if mode == MUL {
				res *= n
			}
		}
	}
	return res, len(expr) - 1
}

func evaluateExpr2(expr string) (int, int) {
	res := 0
	mode := SUM
	for i := 0; i < len(expr); i++ {
		c := rune(expr[i])
		if c == ')' {
			return res, i
		} else if c == '(' {
			recursiveRes, newIndex := evaluateExpr2(expr[i+1:])
			if mode == SUM {
				res += recursiveRes
			} else if mode == MUL {
				res *= recursiveRes
			}
			i += newIndex + 1
			continue
		} else if c == SUM || c == MUL {
			mode = c
		} else {
			n, _ := strconv.Atoi(string(c))
			if mode == SUM {
				res += n
			} else if mode == MUL {
				if i+1 < len(expr)-1 && expr[i+1] == SUM {
					n2, newIndex := evaluateExpr2(expr[i:])
					res *= n2
					i += newIndex
					if i+1 < len(expr)-1 && expr[i+1] == ')' {
						i--
					}
				} else {
					res *= n
				}
			}
		}
	}
	return res, len(expr) - 1
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	sum := 0
	sum2 := 0
	for b.Scan() {
		line := b.Text()
		line = strings.ReplaceAll(line, " ", "")
		res, _ := evaluateExpr(line)
		sum += res

		res2, _ := evaluateExpr2(line)
		sum2 += res2
	}
	fmt.Println("Answer (part 1):", sum)
	fmt.Println("Answer (part 2):", sum2)
}
