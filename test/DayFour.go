package main

import (
	"strconv"
	"strings"
	"test/util"
)

func fourPartOne(commandLine string, allBoards string) int {
	boards := strings.Split(allBoards, "\n\n")
	commands := strings.Split(commandLine, ",")
	inst := 10000
	var finalScore int
	for _, v := range boards {
		board := util.Get2dString(v, "\n", " ")
		var boardFlag [5][5]string
		instCount := 0
		for _, command := range commands {
			instCount++
			isMatch, i, j := mach2DString(board, command)
			if isMatch {
				boardFlag[i][j] = "1"
				isWin := isWin(boardFlag)
				if isWin {
					score := calScore(board, boardFlag, command)
					if instCount < inst {
						inst = instCount
						finalScore = score
					}
					break
				}
			}
		}
	}

	return finalScore
}

func fourPartTwo(commandLine string, allBoards string) int {
	boards := strings.Split(allBoards, "\n\n")
	commands := strings.Split(commandLine, ",")
	inst := 0
	var finalScore int
	for _, v := range boards {
		board := util.Get2dString(v, "\n", " ")
		var boardFlag [5][5]string
		instCount := 0
		for _, command := range commands {
			instCount++
			isMatch, i, j := mach2DString(board, command)
			if isMatch {
				boardFlag[i][j] = "1"
				isWin := isWin(boardFlag)
				if isWin {
					score := calScore(board, boardFlag, command)
					if instCount > inst {
						inst = instCount
						finalScore = score
					}
					break
				}
			}
		}
	}

	return finalScore
}

func isWin(boardFlag [5][5]string) bool {
	for i := range boardFlag {
		result := true
		for j := range boardFlag[0] {
			if strings.Compare(boardFlag[i][j], "") == 0 {
				result = false
				break
			}
		}
		if result {
			return true
		}
	}

	for i := range boardFlag {
		result := true
		for j := range boardFlag[0] {
			if strings.Compare(boardFlag[j][i], "") == 0 {
				result = false
				break
			}
		}
		if result {
			return true
		}

	}
	return false
}

func mach2DString(board [][]string, matcher string) (bool, int, int) {
	for i := range board {
		for j := range board[0] {
			if strings.Compare(board[i][j], matcher) == 0 {
				return true, i, j

			}
		}
	}
	return false, -1, -1
}

func calScore(board [][]string, boardFlag [5][5]string, command string) int {
	sum := 0
	for i := range boardFlag {
		for j := range boardFlag[0] {
			if strings.Compare(boardFlag[i][j], "") == 0 {
				num, _ := strconv.Atoi(board[i][j])
				sum = sum + num
			}
		}
	}
	commandNum, _ := strconv.Atoi(command)
	if commandNum*sum == 0 {
		return -1
	}
	return commandNum * sum
}
