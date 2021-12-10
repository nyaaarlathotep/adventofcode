package main

import (
	"fmt"
	"sort"
	"test/DataStruct"
	"test/util"
)

func tenPartOne(input string) int {

	lines := util.GetStringSlice(input, "\n")
	score := 0
	for _, line := range lines {
		stack := DataStruct.NewStack()

		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if char == ")" {
				left := stack.Pop().(string)
				if left != "(" {
					score += 3
					break
				}
			} else if char == "}" {
				left := stack.Pop().(string)
				if left != "{" {
					score += 1197
					break
				}
			} else if char == "]" {
				left := stack.Pop().(string)
				if left != "[" {
					score += 57
					break
				}
			} else if char == ">" {
				left := stack.Pop().(string)
				if left != "<" {
					score += 25137
					break
				}
			} else {
				stack.Push(string(line[i]))
			}
		}
	}

	return score
}

func tenPartTwo(input string) int {

	lines := util.GetStringSlice(input, "\n")
	scores := make([]int, 0)
	for _, line := range lines {
		stack := DataStruct.NewStack()
		corrupted := false

		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if char == ")" {
				left := stack.Pop().(string)
				if left != "(" {
					corrupted = true
					break
				}
			} else if char == "}" {
				left := stack.Pop().(string)
				if left != "{" {
					corrupted = true

					break
				}
			} else if char == "]" {
				left := stack.Pop().(string)
				if left != "[" {
					corrupted = true

					break
				}
			} else if char == ">" {
				left := stack.Pop().(string)
				if left != "<" {
					corrupted = true
					break
				}
			} else {
				stack.Push(string(line[i]))
			}
		}
		tempScore := 0
		if !corrupted {
			for !stack.IsEmpty() {
				char := stack.Pop().(string)
				if char == "(" {
					tempScore = tempScore*5 + 1
				} else if char == "{" {
					tempScore = tempScore*5 + 3

				} else if char == "[" {
					tempScore = tempScore*5 + 2

				} else if char == "<" {
					tempScore = tempScore*5 + 4

				}
			}
			scores = append(scores, tempScore)
		}

	}
	sort.Ints(scores)
	fmt.Println(len(scores))
	fmt.Println(len(scores) / 2)
	return scores[len(scores)/2]
}
