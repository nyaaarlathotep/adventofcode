package main

import (
	"strings"
)

func eightPartOne(input string) int64 {
	numLengthMap := make(map[int][]int)
	numLengthMap[6] = []int{0}
	numLengthMap[2] = []int{1}
	numLengthMap[5] = []int{2, 3, 5}
	numLengthMap[4] = []int{4}
	numLengthMap[6] = []int{6, 9}
	numLengthMap[3] = []int{7}
	numLengthMap[7] = []int{8}

	lines := strings.Split(input, "\r\n")
	var counter int64
	for _, line := range lines {
		temp := strings.Split(line, " | ")
		output := temp[1]
		numStrs := strings.Split(output, " ")
		for _, numStr := range numStrs {
			length := len(numStr)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				counter++
			}
		}
	}

	return counter
}

func eightPartTwo(input string) int64 {
	numLengthMap := make(map[int][]int)
	numLengthMap[2] = []int{1}
	numLengthMap[3] = []int{7}
	numLengthMap[7] = []int{8}
	numLengthMap[4] = []int{4}

	numLengthMap[5] = []int{2, 3, 5}
	numLengthMap[6] = []int{0, 6, 9}

	lines := strings.Split(input, "\r\n")
	var counter int64
	for _, line := range lines {
		temp := strings.Split(line, " | ")
		output := temp[1]
		tenNums := temp[0]
		numMap := make(map[int]string)

		tenNum := strings.Split(tenNums, " ")
		for _, numStr := range tenNum {
			length := len(numStr)
			if length == 2 {
				numMap[1] = numStr
			} else if length == 3 {
				numMap[7] = numStr

			} else if length == 4 {
				numMap[4] = numStr

			} else if length == 7 {
				numMap[8] = numStr
			}
		}

		for _, numStr := range tenNum {
			if len(numStr) == 5 {
				if strContain(numMap[1], numStr, 0) {
					numMap[3] = numStr
				} else if strContain(numMap[4], numStr, 1) {
					numMap[5] = numStr
				} else {
					numMap[2] = numStr
				}
			} else if len(numStr) == 6 {
				if strContain(numMap[4], numStr, 0) {
					numMap[9] = numStr
				} else if strContain(numMap[1], numStr, 0) {
					numMap[0] = numStr
				} else {
					numMap[6] = numStr
				}
			}
		}

		var totalNow int
		outStrs := strings.Split(output, " ")
		for _, outStr := range outStrs {
			for i, value := range numMap {
				if len(outStr) == len(value) && strContain(outStr, value, 0) {
					totalNow = totalNow*10 + i
					break
				}
			}
		}
		counter = counter + int64(totalNow)
		totalNow = 0

	}

	return counter
}


func strContain(small string, big string, misCount int) bool {
	sSlice := strings.Split(small, "")
	bSlice := strings.Split(big, "")
	for _, s := range sSlice {
		get := false
		for _, ss := range bSlice {
			if s == ss {
				get = true
			}
		}
		if !get {
			misCount--
		}
		if misCount < 0 {
			return false
		}
	}
	return true
}
