package main

import (
	"test/DataStruct"
	"test/util"
)

func tenPartOne(input string) int {

	lines := util.GetStringSlice(input, "\r\n")
	score := 0
	for _, line := range lines {
		stack := DataStruct.NewStack()

		for i := 0; i < len(line); i++ {
			char:=string(line[i])
			// )
			if char == ")" {
				left := stack.Pop().(string)
				if left != "(" {
					score+=3
					break
				}
			}

			// }
			if char == "}" {
				left := stack.Pop().(string)
				if left != "{" {
					score+=1197
					break
				}
			}

			// ]
			if char == "]" {
				left := stack.Pop().(string)
				if left != "[" {
					score+=57
					break
				}
			}

			// >
			if char == ">" {
				left := stack.Pop().(string)
				if left != "<" {
					score+=25137
					break
				}
			}else {
				stack.Push(string(line[i]))
				//stack.Traverse()
			}
		}
	}

	return score
}
