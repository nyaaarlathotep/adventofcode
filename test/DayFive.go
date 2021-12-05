package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func fivePartOne(input string, isPartOne bool) int {
	//var t1 = [...]int{9, 7}
	//var t2 = [...]int{7, 9}
	//fmt.Println(isPartTwoLine(t1, t2))
	var marker [1000][1000]int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		begin, end := getBeginEnd(line)

		if isPartOne {
			formatBeginEnd(&begin, &end)
			if isPartOneLine(begin, end) {
				mark(&marker, begin, end)
			}
		} else {
			if isPartTwoLine(begin, end) {
				formatMarkTwo(&begin, &end)
				markTwo(&marker, begin, end)
			} else {
				formatBeginEnd(&begin, &end)
				if isPartOneLine(begin, end) {
					mark(&marker, begin, end)
				}
			}
		}
	}

	counter := 0
	for i := range marker[0] {
		for j := range marker[0] {
			if marker[i][j] > 1 {
				counter++
			}

		}
	}

	return counter
}

func getBeginEnd(line string) ([2]int, [2]int) {
	parts := strings.Split(line, " -> ")
	if len(parts) != 2 {
		fmt.Println("!!!!!!!!!!!!")
	}
	be := strings.Split(parts[0], ",")
	b1, _ := strconv.Atoi(be[0])
	b2, _ := strconv.Atoi(be[1])
	var begin = [2]int{b1, b2}
	en := strings.Split(parts[1], ",")
	e1, _ := strconv.Atoi(en[0])
	e2, _ := strconv.Atoi(en[1])
	var end = [2]int{e1, e2}
	return begin, end
}

func mark(marker *[1000][1000]int, begin [2]int, end [2]int) {
	for i := begin[0]; i <= end[0]; i++ {
		for j := begin[1]; j <= end[1]; j++ {
			(*marker)[i][j] = (*marker)[i][j] + 1
		}
	}
}

func markTwo(marker *[1000][1000]int, begin [2]int, end [2]int) {
	if begin[0] < end[0] && begin[1] < end[1] {
		j := begin[1]
		for i := begin[0]; i <= end[0]; i++ {
			(*marker)[i][j] = (*marker)[i][j] + 1
			j++
		}
	} else {
		j := begin[1]
		for i := begin[0]; i <= end[0]; i++ {
			(*marker)[i][j] = (*marker)[i][j] + 1
			j--
		}
	}
}

func formatBeginEnd(begin *[2]int, end *[2]int) {
	if (*begin)[0] > (*end)[0] {
		(*begin)[0], (*end)[0] = (*end)[0], (*begin)[0]
	}
	if (*begin)[1] > (*end)[1] {
		(*begin)[1], (*end)[1] = (*end)[1], (*begin)[1]
	}
}

func formatMarkTwo(begin *[2]int, end *[2]int) {
	if begin[0] > end[0] {
		temp := *begin
		*begin = *end
		*end = temp
	}
}

func isPartOneLine(begin [2]int, end [2]int) bool {
	return begin[0] == end[0] || begin[1] == end[1]
}

func isPartTwoLine(begin [2]int, end [2]int) bool {
	return math.Abs(float64(begin[0]-end[0])) == math.Abs(float64(begin[1]-end[1]))
}
