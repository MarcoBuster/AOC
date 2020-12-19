package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type StrictRule struct {
	options []string
}
type Rule struct {
	linkedRules [][]int
}

func parseRule(line string) (int, interface{}) {
	var expr *regexp.Regexp
	var ruleId int
	if strings.ContainsRune(line, '"') {
		expr = regexp.MustCompile("(\\d+): \"(\\w)\"")
		res := expr.FindStringSubmatch(line)
		ruleId, _ = strconv.Atoi(res[1])
		return ruleId, StrictRule{[]string{res[2]}}
	}

	splittedLine := strings.Split(line, " ")
	ruleId, _ = strconv.Atoi(splittedLine[0][:len(splittedLine[0])-1])
	i := 0
	linkedRules := make([][]int, 1, 2)
	for _, rule := range splittedLine[1:] {
		if rule == "|" {
			i++
			linkedRules = append(linkedRules, []int{})
			continue
		}
		n, _ := strconv.Atoi(rule)
		linkedRules[i] = append(linkedRules[i], n)
	}
	return ruleId, Rule{linkedRules}
}

func mergeStrOptions(strOptions [][]string) []string {
	n := len(strOptions)
	indices := make([]int, n)
	var res []string
	for {
		resStr := ""
		for i := 0; i < n; i++ {
			if len(strOptions[i]) == 0 {
				return []string{}
			}
			resStr += strOptions[i][indices[i]]
		}
		res = append(res, resStr)

		next := n - 1
		for next >= 0 && indices[next]+1 >= len(strOptions[next]) {
			next--
		}
		if next < 0 {
			return res
		}
		indices[next]++
		for i := next + 1; i < n; i++ {
			indices[i] = 0
		}
	}
}

func convertRuleToStrictRule(ruleID int, rules map[int]Rule, strictRules *map[int]StrictRule) {
	rule := rules[ruleID]
	var options []string
	for i := 0; i < len(rule.linkedRules); i++ {
		var strOptions [][]string
		for j := 0; j < len(rule.linkedRules[i]); j++ {
			strictRule, ok := (*strictRules)[rule.linkedRules[i][j]]
			if !ok {
				convertRuleToStrictRule(rule.linkedRules[i][j], rules, strictRules)
				strictRule = (*strictRules)[rule.linkedRules[i][j]]
			}
			var strOptions2 []string
			for _, o := range strictRule.options {
				strOptions2 = append(strOptions2, o)
			}
			strOptions = append(strOptions, strOptions2)
		}
		mergedStrOptions := mergeStrOptions(strOptions)
		for _, o := range mergedStrOptions {
			options = append(options, o)
		}
	}
	(*strictRules)[ruleID] = StrictRule{options}
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	defer f.Close()

	strictRules := make(map[int]StrictRule)
	rules := make(map[int]Rule)
	for b.Scan() {
		line := b.Text()
		if line == "" {
			break
		}
		ruleID, rule := parseRule(line)
		_, isStrictRule := rule.(StrictRule)
		if isStrictRule {
			strictRules[ruleID] = rule.(StrictRule)
		} else {
			rules[ruleID] = rule.(Rule)
		}

	}
	for i := range rules {
		convertRuleToStrictRule(i, rules, &strictRules)
	}

	rulesMap := make(map[string]int)
	for k, rule := range strictRules {
		if k != 31 && k != 42 {
			continue
		}
		for _, o := range rule.options {
			rulesMap[o] = k
		}
	}

	fragmentSize := len(strictRules[42].options[0])
	matches1 := 0
	matches2 := 0
ScanLoop:
	for b.Scan() {
		line := b.Text()
		for _, rule := range strictRules {
			for _, o := range rule.options {
				if line == o {
					matches1++
				}
			}
		}
		var seenRules []int
		for i := range line {
			if i > 0 && (i+1)%fragmentSize == 0 {
				sl := line[i+1-fragmentSize : i+1]
				ruleID, ok := rulesMap[sl]
				if !ok {
					fmt.Println("Slice not even found!")
					continue ScanLoop
				}
				seenRules = append(seenRules, ruleID)
			}
		}
		seen42 := 0
		seen31 := 0
		for _, r := range seenRules {
			if seen31 != 0 && r == 42 {
				continue ScanLoop
			} else if r == 31 {
				seen31++
			} else if r == 42 {
				seen42++
			}
		}
		if seen42 < 2 || seen31 == 0 || seen31 >= seen42 {
			continue ScanLoop
		}
		matches2++
	}
	fmt.Println("Answer (part 1):", matches1)
	fmt.Println("Answer (part 2):", matches2)
}
