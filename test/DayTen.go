package main

import (
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
