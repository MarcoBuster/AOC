package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bitmask string

type Assignment struct {
	address uint
	value   uint
}
type Memory = map[uint]string

func parseMask(line string) Bitmask {
	return Bitmask(line[7:])
}

func parseAssignment(line string) Assignment {
	expr := regexp.MustCompile("mem\\[(\\d+)] = (\\d+)")
	res := expr.FindStringSubmatch(line)

	address, _ := strconv.Atoi(res[1])
	value, _ := strconv.Atoi(res[2])
	return Assignment{uint(address), uint(value)}
}

func toBin(num uint) []rune {
	return []rune(fmt.Sprintf("%036b", num))
}

func toDec(num string) uint {
	res, _ := strconv.ParseInt(num, 2, 37)
	return uint(res)
}

func execInstructionPart1(mem *Memory, assignment Assignment, mask Bitmask) {
	newValue := toBin(assignment.value)
	for i := 0; i < len(mask); i++ {
		if mask[i] == 'X' {
			continue
		}
		newValue[i] = rune(mask[i])
	}
	(*mem)[assignment.address] = string(newValue)
}

/*
	This function is a clever (but stupid) hack:
	Example. The addresses computed from 0X00X0 are:
	- 000000	x_1 = 0		x_2 = 0
	- 000010	x_1 = 0		x_2 = 1
	- 010000	x_1 = 1		x_2 = 0
	- 010010	x_1 = 1		x_2 = 1
	Isolate the x:
	| iter. # | x_1 | x_2 |
	|    0    |  0  |  0  |
	|    1    |  0  |  1  |
	|    2    |  1  |  0  |
	|    3    |  1  |  1  |
	That's exactly the behavior of the binary to decimal number conversion.
	I emulated that behaviour in this function, by converting a fixed-length
	binary number to a string to access its digits.
*/
func combineAddresses(address []rune) []uint {
	var addresses []uint
	totalXCount := strings.Count(string(address), "X")
	// 2^(totalXCount) iterations needed
	for i := 0; i < int(math.Pow(2, float64(totalXCount))); i++ {
		candidateAddress := make([]rune, len(address))
		copy(candidateAddress, address)
		// Converts i to a fixed-length binary number with
		// length = log_2(total iterations count) = totalXCount
		binCount := fmt.Sprintf("%0"+strconv.Itoa(totalXCount)+"b", uint(i))
		for k, xCount := 0, 0; k < len(address); k++ {
			if address[k] != 'X' {
				candidateAddress[k] = address[k]
			} else {
				candidateAddress[k] = []rune(binCount)[xCount]
				xCount++
			}
		}
		addresses = append(addresses, toDec(string(candidateAddress)))
	}
	return addresses
}

func execInstructionPart2(mem *Memory, assignment Assignment, mask Bitmask) {
	addressStr := toBin(assignment.address)
	for i := 0; i < len(mask); i++ {
		if mask[i] == '0' {
			continue
		} else if mask[i] == '1' || mask[i] == 'X' {
			addressStr[i] = rune(mask[i])
		}
	}
	addresses := combineAddresses(addressStr)
	for _, a := range addresses {
		(*mem)[a] = string(toBin(assignment.value))
	}
}

func sumMemory(mem *Memory) int {
	sum := 0
	for _, v := range *mem {
		sum += int(toDec(v))
	}
	return sum
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	var lines []string
	for b.Scan() {
		line := b.Text()
		lines = append(lines, line)
	}

	var mem Memory
	var mask Bitmask

	// Part 1
	mem = make(Memory)
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = parseMask(line)
		} else {
			assignment := parseAssignment(line)
			execInstructionPart1(&mem, assignment, mask)
		}
	}
	fmt.Println("Answer (part 1):", sumMemory(&mem))

	// Part 2
	mem = make(Memory)
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = parseMask(line)
		} else {
			assignment := parseAssignment(line)
			execInstructionPart2(&mem, assignment, mask)
		}
	}
	fmt.Println("Answer (part 2):", sumMemory(&mem))
}
