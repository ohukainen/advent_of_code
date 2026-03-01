package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Problem struct {
	numbers  [][]byte
	operator string
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}

func part1(lines []string) int {
	ans := 0
	problems := extractProblemsPart1(lines)

	for _, problem := range problems {
		ans += solveProblem(problem)
	}
	return ans
}

func extractProblemsPart1(lines []string) []Problem {
	operators := strings.Fields(lines[len(lines)-1])
	problems := make([]Problem, len(operators))

	for i, operator := range operators {
		problems[i].operator = operator
	}

	for _, line := range lines[:len(lines)-1] {
		fields := strings.Fields(line)

		for j, field := range fields {
			problems[j].numbers = append(problems[j].numbers, []byte(field))
		}
	}

	return problems
}

func part2(lines []string) int {
	ans := 0

	problems := extractProblemsPart2(lines)

	for _, problem := range problems {
		ans += solveProblem(problem)
	}
	return ans
}

func extractProblemsPart2(lines []string) []Problem {
	var problems []Problem
	numberLines := lines[:len(lines)-1]
	totalNumbers := len(numberLines)
	operators := lines[len(lines)-1]

	maxLineLen := 0
	for _, line := range numberLines {
		if maxLineLen < len(line) {
			maxLineLen = len(line)
		}
	}
	maxLineLen = max(maxLineLen, len(operators))

	for i := range maxLineLen {
		if i < len(operators) && (operators[i] == '+' || operators[i] == '*') {
			var problem Problem
			problem.operator = string(operators[i])
			problem.numbers = make([][]byte, totalNumbers)

			problems = append(problems, problem)
		}
		for k := range totalNumbers {
			lineStr := numberLines[k]
			line := []byte(lineStr)

			problems[len(problems)-1].numbers[k] = append(problems[len(problems)-1].numbers[k], line[i])
		}
	}

	problems = convertToCephalopodMath(problems)

	return problems
}

func convertToCephalopodMath(problems []Problem) []Problem {
	for i, problem := range problems {
		maxLen := 0
		for _, nStr := range problem.numbers {
			if len(nStr) > maxLen {
				maxLen = len(nStr)
			}
		}

		nums := make([][]byte, maxLen)

		for _, nStr := range problem.numbers {
			reversed := slices.Clone(nStr)
			slices.Reverse(reversed)
			for j, r := range reversed {
				nums[j] = append(nums[j], r)
			}
		}

		problems[i].numbers = make([][]byte, len(nums))
		copy(problems[i].numbers, nums)
	}
	return problems
}

func solveProblem(problem Problem) int {
	ans := 0
	if problem.operator == "*" {
		ans = 1
	}

	for _, nStr := range problem.numbers {
		trimmed := strings.TrimSpace(string(nStr))
		if trimmed == "" {
			continue
		}
		n, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(err)
		}

		switch problem.operator {
		case "+":
			ans += n
		case "*":
			ans *= n
		default:
			panic("unknown operator")
		}
	}
	return ans
}
