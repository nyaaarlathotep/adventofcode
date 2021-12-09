package main

import (
	"strconv"
	"test/DataStruct"
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

func ninePartTwo(input string) int {

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

	areas := make([]DataStruct.Set, 0)

	for _, v := range lowLand {
		area := DataStruct.NewSet()
		spread(area, &wholeMap, v[0], v[1])
		areas = append(areas, *area)
	}

	max := 0
	maxTwo := 0
	maxThree := 0

	for _, area := range areas {
		if area.Size() > max {
			maxThree = maxTwo
			maxTwo = max
			max = area.Size()
		} else if area.Size() > maxTwo {
			maxThree = maxTwo
			maxTwo = area.Size()
		} else if area.Size() > maxThree {
			maxThree = area.Size()
		}
	}

	return max * maxTwo * maxThree
}

func spread(area *DataStruct.Set, wholeMap *[][]string, i int, j int) {
	area.Add([2]int{i, j})
	(*wholeMap)[i][j] = "9"

	if i > 0 {
		height, _ := strconv.Atoi((*wholeMap)[i-1][j])
		if height < 9 && !area.Contains([2]int{i - 1, j}) {
			spread(area, wholeMap, i-1, j)
		}
	}
	if i < len(*wholeMap)-1 {
		height, _ := strconv.Atoi((*wholeMap)[i+1][j])
		if height < 9 && !area.Contains([2]int{i + 1, j}) {
			spread(area, wholeMap, i+1, j)
		}
	}
	if j > 0 {
		height, _ := strconv.Atoi((*wholeMap)[i][j-1])
		if height < 9 && !area.Contains([2]int{i, j - 1}) {
			spread(area, wholeMap, i, j-1)
		}
	}
	if j < len((*wholeMap)[0])-1 {
		height, _ := strconv.Atoi((*wholeMap)[i][j+1])
		if height < 9 && !area.Contains([2]int{i, j + 1}) {
			spread(area, wholeMap, i, j+1)
		}
	}
}

func spreadTwo(area *DataStruct.Set, wholeMap *[][]string, i int, j int) {
	area.Add([2]int{i, j})

	if i > 0 {
		height, _ := strconv.Atoi((*wholeMap)[i-1][j])
		if height < 9 && !area.Contains([2]int{i - 1, j}) {
			spread(area, wholeMap, i-1, j)
		}
	}
	if i < len(*wholeMap)-1 {
		height, _ := strconv.Atoi((*wholeMap)[i+1][j])
		if height < 9 && !area.Contains([2]int{i + 1, j}) {
			spread(area, wholeMap, i+1, j)
		}
	}
	if j > 0 {
		height, _ := strconv.Atoi((*wholeMap)[i][j-1])
		if height < 9 && !area.Contains([2]int{i, j - 1}) {
			spread(area, wholeMap, i, j-1)
		}
	}
	if j < len((*wholeMap)[0])-1 {
		height, _ := strconv.Atoi((*wholeMap)[i][j+1])
		if height < 9 && !area.Contains([2]int{i, j + 1}) {
			spread(area, wholeMap, i, j+1)
		}
	}
}
