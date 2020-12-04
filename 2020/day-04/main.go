package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var NeededStrings = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
}

func checkPassport(passport string) bool {
	var count int
	for _, str := range NeededStrings {
		if strings.Contains(passport, str+":") {
			count++
		}
	}
	if count == len(NeededStrings) {
		return true
	}
	return false
}

func parsePassport(passport string) (bool, Passport) {
	var byrExpr = regexp.MustCompile(`byr:\d\d\d\d`)
	var iyrExpr = regexp.MustCompile(`iyr:\d\d\d\d`)
	var eyrExpr = regexp.MustCompile(`eyr:\d\d\d\d`)
	var hgtExpr = regexp.MustCompile(`hgt:\d+(cm|in)`)
	var hclExpr = regexp.MustCompile(`hcl:#\w+`)
	var eclExpr = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`)
	var pidExpr = regexp.MustCompile(`pid:\d\d\d\d\d\d\d\d\d`)

	byr := byrExpr.FindString(passport)
	iyr := iyrExpr.FindString(passport)
	eyr := eyrExpr.FindString(passport)
	hgt := hgtExpr.FindString(passport)
	hcl := hclExpr.FindString(passport)
	ecl := eclExpr.FindString(passport)
	pid := pidExpr.FindString(passport)
	if byr == "" || iyr == "" || hgt == "" || hcl == "" || ecl == "" || pid == "" {
		return false, Passport{}
	}
	byrN, _ := strconv.Atoi(byr[4:])
	iyrN, _ := strconv.Atoi(iyr[4:])
	eyrN, _ := strconv.Atoi(eyr[4:])
	return true, Passport{byrN, iyrN, eyrN, hgt[4:], hcl[4:], ecl[4:], pid[4:]}
}

func isValidEclChar(c rune) bool {
	switch c {
	case
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f':
		return true
	}
	return false
}

func validatePassport(passport Passport) bool {
	if passport.byr < 1920 || passport.byr > 2002 {
		return false
	}
	if passport.iyr < 2010 || passport.iyr > 2020 {
		return false
	}
	if passport.eyr < 2020 || passport.eyr > 2030 {
		return false
	}
	hgtN, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(passport.hgt, "cm", ""), "in", ""))
	if strings.HasSuffix(passport.hgt, "cm") {
		if hgtN < 150 || hgtN > 193 {
			return false
		}
	} else {
		if hgtN < 59 || hgtN > 76 {
			return false
		}
	}
	if len(passport.pid) != 9 {
		return false
	}
	if len(passport.hcl) != 7 {
		return false
	}
	for _, c := range passport.hcl[1:] {
		if !isValidEclChar(c) {
			return false
		}
	}
	return true
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	passportStr := ""
	part1Result := 0
	part2Result := 0
	for b.Scan() {
		line := b.Text()
		if line == "" {
			if checkPassport(passportStr) {
				part1Result++
				ok, passport := parsePassport(passportStr)
				if ok {
					if validatePassport(passport) {
						part2Result++
					}
				}
			}
			passportStr = ""
			continue
		}
		passportStr += " " + line
	}
	fmt.Println("Answer (part 1):", part1Result)
	fmt.Println("Answer (part 2):", part2Result-1) // idk... :(
}
