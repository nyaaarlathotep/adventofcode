package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"test/util"
)

func threePartTwo(input string) int {
	lines := strings.Split(input, "\n")
	gammaSlice, epsilonSlice := getGammaAndSig(lines)
	rangeLines := lines
	for i := 0; i < 12; i++ {
		tempSlice := make([]string, 0)
		for _, line := range rangeLines {
			chars := strings.Split(line, "")
			if strings.Compare(chars[i], gammaSlice[i]) == 0 {
				tempSlice = append(tempSlice, line)
			}
		}
		rangeLines = tempSlice
		gammaSlice, _ = getGammaAndSig(rangeLines)
		if len(tempSlice) == 1 {
			break
		}
	}
	oxygen := util.GetDecimalFromBinaryString(rangeLines)
	fmt.Println("oxygen: ", oxygen)
	rangeLines = lines
	for i := 0; i < 12; i++ {
		tempSlice := make([]string, 0)
		for _, line := range rangeLines {
			chars := strings.Split(line, "")
			if strings.Compare(chars[i], epsilonSlice[i]) == 0 {
				tempSlice = append(tempSlice, line)
			}
		}
		rangeLines = tempSlice
		_, epsilonSlice = getGammaAndSig(rangeLines)
		if len(tempSlice) == 1 {
			break
		}
	}
	CO2 := util.GetDecimalFromBinaryString(rangeLines)
	fmt.Println("CO2: ", CO2)
	return oxygen * CO2
}

func threePartOne(input string) ([]int, []int) {
	line := strings.Split(input, "\n")
	tSlice := make([][]int, 0)
	for _, v := range line {
		temp := make([]int, 0)
		line = strings.Split(v, "")
		for _, char := range line {
			num, _ := strconv.Atoi(char)
			temp = append(temp, num)
		}
		tSlice = append(tSlice, temp)
	}

	gammaSlice := make([]int, 12)
	epsilonSlice := make([]int, 12)
	zeros := make([]int, 12)
	ones := make([]int, 12)
	for i := range tSlice {
		for j := range tSlice[i] {
			if tSlice[i][j] == 0 {
				zeros[j] = zeros[j] + 1
			} else {
				ones[j] = ones[j] + 1
			}
		}
	}

	for i := range ones {
		if ones[i] > zeros[i] {
			gammaSlice[i] = 0
			epsilonSlice[i] = 1
		} else if zeros[i] > ones[i] {
			gammaSlice[i] = 1
			epsilonSlice[i] = 0
		} else {
			fmt.Println("???????")
		}
	}
	fmt.Println("gammaSlice", gammaSlice)
	fmt.Println("epsilonSlice", epsilonSlice)
	var buffer bytes.Buffer
	for _, v := range gammaSlice {
		buffer.WriteString(strconv.Itoa(v))
	}
	gamma, _ := strconv.ParseInt(buffer.String(), 2, 64)
	buffer.Reset()
	for _, v := range epsilonSlice {
		buffer.WriteString(strconv.Itoa(v))
	}
	epsilon, _ := strconv.ParseInt(buffer.String(), 2, 64)
	fmt.Println("power consumption of the submarine: ", gamma*epsilon)
	return gammaSlice, epsilonSlice
}

func getGammaAndSig(line []string) ([]string, []string) {
	tSlice := make([][]int, 0)
	for _, v := range line {
		temp := make([]int, 0)
		line = strings.Split(v, "")
		for _, char := range line {
			num, _ := strconv.Atoi(char)
			temp = append(temp, num)
		}
		tSlice = append(tSlice, temp)
	}

	gammaSlice := make([]string, 12)
	epsilonSlice := make([]string, 12)

	zeros := make([]int, 12)
	ones := make([]int, 12)
	for i := range tSlice {
		for j := range tSlice[i] {
			if tSlice[i][j] == 0 {
				zeros[j] = zeros[j] + 1
			} else {
				ones[j] = ones[j] + 1
			}
		}
	}
	for i := range ones {
		if ones[i] >= zeros[i] {
			gammaSlice[i] = "1"
			epsilonSlice[i] = "0"
		} else if zeros[i] > ones[i] {
			gammaSlice[i] = "0"
			epsilonSlice[i] = "1"
		}
	}
	return gammaSlice, epsilonSlice
}
