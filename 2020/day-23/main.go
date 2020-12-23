package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	val  int
	next *Node
}

type Cups struct {
	head    *Node
	current *Node
	nodes   map[int]*Node
	l       int
}

func (c *Cups) move() {
	pickupHead := c.current.next
	pickupTail := pickupHead.next.next
	c.current.next = pickupTail.next

	dest := c.current.val - 1
	if dest < 1 {
		dest += c.l
	}
	destFound := true
	for destFound {
		destFound = false
		for _, v := range [3]int{pickupHead.val, pickupHead.next.val, pickupTail.val} {
			if v == dest {
				destFound = true
				break
			}
		}
		if destFound {
			dest--
			if dest < 1 {
				dest += c.l
			}
		}
	}
	d := c.nodes[dest]
	pickupTail.next = d.next
	d.next = pickupHead

	c.current = c.current.next
}

func (c *Cups) value() int {
	v := 0
	one := c.nodes[1]
	for curr := one.next; curr != one; curr = curr.next {
		v *= 10
		v += curr.val
	}
	return v
}

func parseInput(input string) []int {
	res := make([]int, 0)
	for _, c := range input {
		n, _ := strconv.Atoi(string(c))
		res = append(res, n)
	}
	return res
}

func newNode(val int) *Node {
	n := Node{val: val, next: nil}
	return &n
}

func newCups(input []int) *Cups {
	val := input[0]
	prev := newNode(val)
	c := Cups{prev, prev, make(map[int]*Node), len(input)}
	c.nodes[val] = prev

	for i := 1; i < c.l; i++ {
		val = input[i]
		curr := newNode(val)
		prev.next = curr
		c.nodes[val] = curr
		prev = curr
	}
	c.nodes[input[c.l-1]].next = c.head
	return &c
}

func playPart1(cups []int, iterations int) int {
	c := newCups(cups)
	for i := 0; i < iterations; i++ {
		c.move()
	}
	return c.value()
}

func playPart2(cups []int) int64 {
	dimension := 1000000
	iterations := 10000000
	padded := make([]int, dimension)
	copy(padded, cups)
	for i := len(cups); i < dimension; i++ {
		padded[i] = i + 1
	}
	c := newCups(padded)
	for i := 0; i < iterations; i++ {
		c.move()
	}

	one := c.nodes[1]
	first := one.next.val
	second := one.next.next.val
	return int64(first) * int64(second)
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	b.Scan()
	input := b.Text()
	cups := parseInput(input)

	fmt.Println("Answer (part 1):", playPart1(cups, 100))
	fmt.Println("Answer (part 2):", playPart2(cups))
}
