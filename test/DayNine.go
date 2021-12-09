package main

import (
	"strconv"
	"test/util"
)

func ninePartOne(input string) int {

	wholeMap := util.Get2dString(input, "\r\n", "")
	lowLand := make([][3]int, 0)
	for i := range wholeMap {
		for j := range wholeMap[0] {
			isLow := true
			if i > 0 && wholeMap[i][j] >= wholeMap[i-1][j] {
				isLow = false
			}
			if isLow && i < len(wholeMap)-1 && wholeMap[i][j] >= wholeMap[i+1][j] {
				isLow = false

			}
			if isLow && j > 0 && wholeMap[i][j] >= wholeMap[i][j-1] {
				isLow = false

			}
			if isLow && j < len(wholeMap[0])-1 && wholeMap[i][j] >= wholeMap[i][j+1] {
				isLow = false

			}
			if isLow {
				height, _ := strconv.Atoi(wholeMap[i][j])
				lowLand = append(lowLand, [3]int{i, j, height})
			}
		}
	}
	total := 0
	for _, v := range lowLand {
		total += v[2]
		total++
	}

	return total
}
