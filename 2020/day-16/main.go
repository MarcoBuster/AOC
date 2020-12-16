package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}
type Field struct {
	name   string
	values [2]Range
}
type Ticket = []int

func unsafeToInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func parseFields(line string) Field {
	expr := regexp.MustCompile("(.*): (\\d+)-(\\d+) or (\\d+)-(\\d+)")
	res := expr.FindStringSubmatch(line)
	return Field{
		res[1],
		[2]Range{
			{unsafeToInt(res[2]), unsafeToInt(res[3])},
			{unsafeToInt(res[4]), unsafeToInt(res[5])},
		},
	}
}

func parseTicket(line string) Ticket {
	numsSlice := strings.Split(line, ",")
	var ticket Ticket
	for _, n := range numsSlice {
		ticket = append(ticket, unsafeToInt(n))
	}
	return ticket
}

func checkTicketFields(v int, fields []Field) bool {
	isValid := len(fields)
	for _, f := range fields {
		if (v < f.values[0].start || v > f.values[0].end) && (v < f.values[1].start || v > f.values[1].end) {
			isValid--
		}
	}
	if isValid == 0 {
		return false
	}
	return true
}

func isValidTicket(t Ticket, fields []Field) int {
	for _, v := range t {
		if !checkTicketFields(v, fields) {
			return v
		}
	}
	return -1
}

func bruteforceField(tickets []Ticket, f Field, candidateIndex int) bool {
	for _, t := range tickets {
		v := t[candidateIndex]
		if (v < f.values[0].start || v > f.values[0].end) && (v < f.values[1].start || v > f.values[1].end) {
			return false
		}
	}
	return true
}

func removeIndex(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func removeDuplicates(indexFields map[int][]string, stringToRemove string) {
	for k, v := range indexFields {
		if len(v) == 1 {
			continue
		}
		for i := 0; i < len(v); i++ {
			if v[i] == stringToRemove {
				indexFields[k] = removeIndex(v, i)
			}
		}
	}
}

func generateMappings(indexFields map[int][]string) map[int]string {
	mappings := make(map[int]string, len(indexFields))
	for i := 0; i < len(indexFields); i++ {
		for k, v := range indexFields {
			if len(v) == 1 {
				mappings[k] = v[0]
				removeDuplicates(indexFields, v[0])
			}
		}
	}
	return mappings
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	cur := 0
	var fields []Field
	var ticket Ticket
	var nearbyTickets []Ticket
	for b.Scan() {
		line := b.Text()
		if line == "" {
			continue
		}
		if line == "your ticket:" || line == "nearby tickets:" {
			cur++
			continue
		}

		if cur == 0 { // Parse fields
			field := parseFields(line)
			fields = append(fields, field)
		} else if cur == 1 { // Parse ticket
			ticket = parseTicket(line)
		} else if cur == 2 { // Parse nearby tickets
			nearbyTickets = append(nearbyTickets, parseTicket(line))
		}
	}

	var validTickets []Ticket
	var invalidValues []int
	for _, t := range nearbyTickets {
		v := isValidTicket(t, fields)
		if v == -1 {
			validTickets = append(validTickets, t)
		} else {
			invalidValues = append(invalidValues, v)
		}
	}
	invalidValuesSum := 0
	for _, v := range invalidValues {
		invalidValuesSum += v
	}
	fmt.Println("Answer (part 1):", invalidValuesSum)

	indexField := make(map[int][]string)
	for i := 0; i < len(fields); i++ {
		for _, f := range fields {
			if bruteforceField(validTickets, f, i) {
				indexField[i] = append(indexField[i], f.name)
			}
		}
	}
	mappings := generateMappings(indexField)
	mul := 1
	for k, v := range mappings {
		if !strings.HasPrefix(v, "departure") {
			continue
		}
		//goland:noinspection ALL
		mul *= ticket[k]
	}
	fmt.Println("Answer (part 2):", mul)
}
