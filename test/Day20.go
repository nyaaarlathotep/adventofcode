package main

import (
	"fmt"
	"strconv"
	"strings"
	"test/util"
)

func day20(input string) int {

	parts := strings.Split(input, "\r\n\r\n")

	rule := parts[0]
	imageMap := util.Get2dString(parts[1], "\r\n", "")
	numMap := convertMap(&imageMap)
	for count := 0; count < 50; count++ {
		nextMap := getNextRes(numMap)
		if count%2 == 0 {
			for i := 0; i < len((*nextMap)[0]); i++ {
				(*nextMap)[0][i] = 1
				(*nextMap)[len(*nextMap)-1][i] = 1
			}

			for i := 0; i < len(*nextMap); i++ {
				(*nextMap)[i][0] = 1
				(*nextMap)[i][len((*nextMap)[0])-1] = 1
			}
		}

		for i := 0; i < len(*numMap); i++ {
			for j := 0; j < len((*numMap)[0]); j++ {

				around := aroundNumStr(*numMap, i, j, count)
				ruleIndex, e := strconv.ParseInt(around, 2, 64)
				if e != nil {
					fmt.Println(e)
				}
				lightMark := string(rule[ruleIndex])
				if lightMark == "#" {
					(*nextMap)[i+1][j+1] = 1
				} else {
					(*nextMap)[i+1][j+1] = 0
				}
			}
		}
		numMap = nextMap
	}
	count := 0
	for _, l := range *numMap {
		for _, num := range l {
			if num == 1 {
				count++
			}
		}

	}
	return count
}

func aroundNumStr(numMap [][]int, i int, j int, count int) string {
	conStr := "0"
	if count%2 == 1 {
		conStr = "1"
	}
	res := ""
	if i == 0 {
		res = conStr + conStr + conStr
	} else if j == 0 {
		res = conStr + strconv.Itoa(numMap[i-1][j]) + strconv.Itoa(numMap[i-1][j+1])
	} else if j == len(numMap[0])-1 {
		res = strconv.Itoa(numMap[i-1][j-1]) + strconv.Itoa(numMap[i-1][j]) + conStr
	} else {
		res = strconv.Itoa(numMap[i-1][j-1]) + strconv.Itoa(numMap[i-1][j]) + strconv.Itoa(numMap[i-1][j+1])
	}

	if j == 0 {
		res = res + conStr + strconv.Itoa(numMap[i][j]) + strconv.Itoa(numMap[i][j+1])
	} else if j == len(numMap[0])-1 {
		res = res + strconv.Itoa(numMap[i][j-1]) + strconv.Itoa(numMap[i][j]) + conStr
	} else {
		res = res + strconv.Itoa(numMap[i][j-1]) + strconv.Itoa(numMap[i][j]) + strconv.Itoa(numMap[i][j+1])
	}

	if i == len(numMap)-1 {
		res = res + conStr+conStr+conStr
	} else if j == 0 {
		res = res + conStr + strconv.Itoa(numMap[i+1][j]) + strconv.Itoa(numMap[i+1][j+1])
	} else if j == len(numMap[0])-1 {
		res = res + strconv.Itoa(numMap[i+1][j-1]) + strconv.Itoa(numMap[i+1][j]) + conStr
	} else {
		res = res + strconv.Itoa(numMap[i+1][j-1]) + strconv.Itoa(numMap[i+1][j]) + strconv.Itoa(numMap[i+1][j+1])
	}

	return res
}

func convertMap(imageMap *[][]string) *[][]int {
	numMap := make([][]int, 0)
	numMap = append(numMap, make([]int, len((*imageMap)[0])+2))
	for i, line := range *imageMap {
		numMap = append(numMap, make([]int, 0))
		numMap[i+1] = append(numMap[i+1], 0)
		for _, light := range line {
			if light == "#" {
				numMap[i+1] = append(numMap[i+1], 1)
			} else if light == "." {
				numMap[i+1] = append(numMap[i+1], 0)

			} else {
				fmt.Println("????")
			}
		}
		numMap[i+1] = append(numMap[i+1], 0)
	}
	numMap = append(numMap, make([]int, len((*imageMap)[0])+2))
	return &numMap
}

func getNextRes(numMap *[][]int) *[][]int {
	nextNumMap := make([][]int, 0)
	nextNumMap = append(nextNumMap, make([]int, len((*numMap)[0])+2))
	for i := range *numMap {
		nextNumMap = append(nextNumMap, make([]int, len((*numMap)[i])+2))
	}
	nextNumMap = append(nextNumMap, make([]int, len((*numMap)[0])+2))
	return &nextNumMap
}

func printImage(numMap *[][]int) {
	for _, l := range *numMap {
		for _, v := range l {
			if v == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
